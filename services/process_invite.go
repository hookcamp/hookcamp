package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/frain-dev/convoy/api/models"
	"github.com/oklog/ulid/v2"

	"github.com/frain-dev/convoy/datastore"
	"github.com/frain-dev/convoy/pkg/log"
	"github.com/frain-dev/convoy/queue"
	"github.com/frain-dev/convoy/util"
)

type ProcessInviteService struct {
	Queue         queue.Queuer
	InviteRepo    datastore.OrganisationInviteRepository
	UserRepo      datastore.UserRepository
	OrgRepo       datastore.OrganisationRepository
	OrgMemberRepo datastore.OrganisationMemberRepository

	Token    string
	Accepted bool
	NewUser  *models.User
}

func (pis *ProcessInviteService) Run(ctx context.Context) error {
	iv, err := pis.InviteRepo.FetchOrganisationInviteByToken(ctx, pis.Token)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to fetch organisation member invite by token and email")
		return util.NewServiceError(http.StatusBadRequest, errors.New("failed to fetch organisation member invite"))
	}

	if iv.Status != datastore.InviteStatusPending {
		return util.NewServiceError(http.StatusBadRequest, fmt.Errorf("organisation member invite already %s", iv.Status.String()))
	}

	if time.Now().After(iv.ExpiresAt) { // if current date has surpassed expiry date
		return util.NewServiceError(http.StatusBadRequest, errors.New("organisation member invite already expired"))
	}

	if !pis.Accepted {
		iv.Status = datastore.InviteStatusDeclined
		err = pis.InviteRepo.UpdateOrganisationInvite(ctx, iv)
		if err != nil {
			log.FromContext(ctx).WithError(err).Error("failed to update declined organisation invite")
			return util.NewServiceError(http.StatusBadRequest, errors.New("failed to update declined organisation invite"))
		}
		return nil
	}

	user, err := pis.UserRepo.FindUserByEmail(ctx, iv.InviteeEmail)
	if err != nil {
		if errors.Is(err, datastore.ErrUserNotFound) {
			user, err = pis.createNewUser(ctx, pis.NewUser, iv.InviteeEmail)
			if err != nil {
				return err
			}
		} else {
			log.FromContext(ctx).WithError(err).Error("failed to find user by email")
			return util.NewServiceError(http.StatusBadRequest, errors.New("failed to find user by email"))
		}
	}

	org, err := pis.OrgRepo.FetchOrganisationByID(ctx, iv.OrganisationID)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to fetch organisation by id")
		return util.NewServiceError(http.StatusBadRequest, errors.New("failed to fetch organisation by id"))
	}

	_, err = NewOrganisationMemberService(pis.OrgMemberRepo).CreateOrganisationMember(ctx, org, user, &iv.Role)
	if err != nil {
		return err
	}

	iv.Status = datastore.InviteStatusAccepted
	err = pis.InviteRepo.UpdateOrganisationInvite(ctx, iv)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to update accepted organisation invite")
		return util.NewServiceError(http.StatusBadRequest, errors.New("failed to update accepted organisation invite"))
	}

	return nil
}

func (pis *ProcessInviteService) createNewUser(ctx context.Context, newUser *models.User, email string) (*datastore.User, error) {
	if newUser == nil {
		return nil, util.NewServiceError(http.StatusBadRequest, errors.New("new user is nil"))
	}

	err := util.Validate(newUser)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to validate new user information")
		return nil, util.NewServiceError(http.StatusBadRequest, err)
	}

	p := datastore.Password{Plaintext: newUser.Password}
	err = p.GenerateHash()
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to generate user password hash")
		return nil, util.NewServiceError(http.StatusBadRequest, errors.New("failed to create organisation member invite"))
	}

	user := &datastore.User{
		UID:       ulid.Make().String(),
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     email,
		Password:  string(p.Hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = pis.UserRepo.CreateUser(ctx, user)
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to create user")
		return nil, util.NewServiceError(http.StatusBadRequest, errors.New("failed to create user"))
	}

	return user, nil
}
