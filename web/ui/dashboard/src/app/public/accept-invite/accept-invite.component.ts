import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { GeneralService } from 'src/app/services/general/general.service';
import { AcceptInviteService } from './accept-invite.service';

@Component({
	selector: 'app-accept-invite',
	templateUrl: './accept-invite.component.html',
	styleUrls: ['./accept-invite.component.scss']
})
export class AcceptInviteComponent implements OnInit {
	showPassword = false;
	showConfirmPassword = false;
	loading = false;
	token!: string;
	fetchingDetails = false;
	acceptTerms = false;
	userDetailsAvailable = false;
	acceptInviteForm: FormGroup = this.formBuilder.group({
		first_name: ['', Validators.required],
		last_name: ['', Validators.required],
		email: ['', Validators.required],
		role: this.formBuilder.group({
			type: ['super_user']
		}),
		password: ['', Validators.compose([Validators.minLength(8), Validators.required])],
		password_confirmation: ['', Validators.required]
	});

	constructor(private formBuilder: FormBuilder, private acceptInviteService: AcceptInviteService, private route: ActivatedRoute, private router: Router, private generalService: GeneralService) {}

	ngOnInit() {
		this.token = this.route.snapshot.queryParams.token;
		this.getUserDetails(this.token);
	}

	async getUserDetails(token: string) {
		this.fetchingDetails = true;
		try {
			const response = await this.acceptInviteService.getUserDetails(token);
			response.data.user ? (this.userDetailsAvailable = true) : (this.userDetailsAvailable = false);
			const inviteeDetails = response.data.token;
			const userDetails = response.data.user;
			this.acceptInviteForm.patchValue({
				first_name: userDetails?.first_name ? userDetails.first_name : '',
				last_name: userDetails?.last_name ? userDetails.last_name : '',
				email: inviteeDetails.invitee_email,
				role: { type: inviteeDetails.role.type }
			});
			this.fetchingDetails = false;
		} catch {
			this.fetchingDetails = false;
		}
	}
	async acceptInvite() {
		if (this.acceptInviteForm.invalid) {
			(<any>Object).values(this.acceptInviteForm.controls).forEach((control: FormControl) => {
				control?.markAsTouched();
			});
			return;
		}
		this.loading = true;
		try {
			const response = await this.acceptInviteService.acceptInvite({ token: this.token, body: this.acceptInviteForm.value });
			this.loading = false;
			if (response.status === true) {
				this.generalService.showNotification({ style: 'success', message: response.message });
				this.router.navigateByUrl('login');
			}
		} catch (error: any) {
			this.loading = false;
			this.generalService.showNotification({ style: 'error', message: error.error.message });
		}
	}

	checkPassword(): boolean {
		const newPassword = this.acceptInviteForm.value.password;
		const confirmPassword = this.acceptInviteForm.value.password_confirmation;
		if (newPassword === confirmPassword) {
			return true;
		} else {
			return false;
		}
	}

	checkForNumber(password: string): boolean {
		const regex = /\d/;
		return regex.test(password);
	}

	checkForSpecialCharacter(password: string): boolean {
		const regex = /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+/;
		return regex.test(password);
	}
}
