package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://getconvoy.io/terms",
        "contact": {
            "name": "Convoy Support",
            "url": "https://getconvoy.io/docs",
            "email": "support@getconvoy.io"
        },
        "license": {
            "name": "Mozilla Public License 2.0",
            "url": "https://www.mozilla.org/en-US/MPL/2.0/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/projects/{projectID}/endpoints": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches an endpoints",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "List all endpoints",
                "operationId": "GetEndpoints",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "The owner ID of the endpoint",
                        "name": "ownerId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "endpoint-1",
                        "description": "The name of the endpoint",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.EndpointResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates an endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Create an endpoint",
                "operationId": "CreateEndpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Endpoint Details",
                        "name": "endpoint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateEndpoint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EndpointResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/endpoints/{endpointID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches an endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Retrieve endpoint",
                "operationId": "GetEndpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EndpointResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint updates an endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Update an endpoint",
                "operationId": "UpdateEndpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Endpoint Details",
                        "name": "endpoint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateEndpoint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EndpointResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint deletes an endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Delete endpoint",
                "operationId": "DeleteEndpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/endpoints/{endpointID}/expire_secret": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint expires and re-generates the endpoint secret.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Roll endpoint secret",
                "operationId": "ExpireSecret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Expire Secret Body Parameters",
                        "name": "endpoint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ExpireSecret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EndpointResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/endpoints/{endpointID}/pause": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint toggles an endpoint status between the active and paused states",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Pause endpoint",
                "operationId": "PauseEndpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Endpoint ID",
                        "name": "endpointID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EndpointResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves all event deliveries paginated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Deliveries"
                ],
                "summary": "List all event deliveries",
                "operationId": "GetEventDeliveriesPaged",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2008-05-02T15:04:05",
                        "description": "The end date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of endpoint IDs to filter by",
                        "name": "endpointId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Event ID to filter by",
                        "name": "eventId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "EventType to filter by",
                        "name": "event_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IdempotencyKey to filter by",
                        "name": "idempotencyKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05",
                        "description": "The start date",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of event delivery statuses to filter by",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SubscriptionID to filter by",
                        "name": "subscriptionId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.EventDeliveryResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/batchretry": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint batch retries multiple event deliveries at once.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Deliveries"
                ],
                "summary": "Batch retry event delivery",
                "operationId": "BatchRetryEventDelivery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2008-05-02T15:04:05",
                        "description": "The end date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of endpoint IDs to filter by",
                        "name": "endpointId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Event ID to filter by",
                        "name": "eventId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "EventType to filter by",
                        "name": "event_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IdempotencyKey to filter by",
                        "name": "idempotencyKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05",
                        "description": "The start date",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of event delivery statuses to filter by",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SubscriptionID to filter by",
                        "name": "subscriptionId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/forceresend": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint enables you retry a previously successful event delivery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Deliveries"
                ],
                "summary": "Force retry event delivery",
                "operationId": "ForceResendEventDeliveries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "event delivery ids",
                        "name": "deliveryIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.IDs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/{eventDeliveryID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches an event delivery.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Deliveries"
                ],
                "summary": "Retrieve an event delivery",
                "operationId": "GetEventDelivery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event delivery id",
                        "name": "eventDeliveryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EventDeliveryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/{eventDeliveryID}/deliveryattempts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches an app message's delivery attempts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delivery Attempts"
                ],
                "summary": "List delivery attempts",
                "operationId": "GetDeliveryAttempts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event delivery id",
                        "name": "eventDeliveryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/datastore.DeliveryAttempt"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/{eventDeliveryID}/deliveryattempts/{deliveryAttemptID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches an app event delivery attempt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delivery Attempts"
                ],
                "summary": "Retrieve a delivery attempt",
                "operationId": "GetDeliveryAttempt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event delivery id",
                        "name": "eventDeliveryID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "delivery attempt id",
                        "name": "deliveryAttemptID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/datastore.DeliveryAttempt"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/eventdeliveries/{eventDeliveryID}/resend": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retries an event delivery.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event Deliveries"
                ],
                "summary": "Retry event delivery",
                "operationId": "ResendEventDelivery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event delivery id",
                        "name": "eventDeliveryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EventDeliveryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches app events with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "List all events",
                "operationId": "GetEventsPaged",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2008-05-02T15:04:05",
                        "description": "The end date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of endpoint ids to filter by",
                        "name": "endpointId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IdempotencyKey to filter by",
                        "name": "idempotencyKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Any arbitrary value to filter the events payload",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of Source IDs to filter the events by.",
                        "name": "sourceId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05",
                        "description": "The start date",
                        "name": "startDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.EventResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates an endpoint event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Create an event",
                "operationId": "CreateEndpointEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event Details",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/batchreplay": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint replays multiple events at once.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Batch replay events",
                "operationId": "BatchReplayEvents",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2008-05-02T15:04:05",
                        "description": "The end date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of endpoint ids to filter by",
                        "name": "endpointId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "IdempotencyKey to filter by",
                        "name": "idempotencyKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Any arbitrary value to filter the events payload",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of Source IDs to filter the events by.",
                        "name": "sourceId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05",
                        "description": "The start date",
                        "name": "startDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/datastore.Event"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "$ref": "#/definitions/handlers.Stub"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/broadcast": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates a event that is broadcast to every endpoint whose subscription matches the given event type.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Create a broadcast event",
                "operationId": "CreateBroadcastEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Broadcast Event Details",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BroadcastEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EventResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/dynamic": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint does not require creating endpoint and subscriptions ahead of time. Instead, you supply the endpoint and the payload, and Convoy delivers the events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Dynamic Events",
                "operationId": "CreateDynamicEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event Details",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DynamicEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Stub"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/fanout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint uses the owner_id to fan out an event to multiple endpoints.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Fan out an event",
                "operationId": "CreateEndpointFanoutEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event Details",
                        "name": "event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FanoutEvent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/{eventID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Retrieve an event",
                "operationId": "GetEndpointEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EventResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/events/{eventID}/replay": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint replays an event afresh assuming it is a new event.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Replay event",
                "operationId": "ReplayEndpointEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.EventResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/meta-events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches meta events with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meta Events"
                ],
                "summary": "List all meta events",
                "operationId": "GetMetaEventsPaged",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2008-05-02T15:04:05",
                        "description": "The end date",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2006-01-02T15:04:05",
                        "description": "The start date",
                        "name": "startDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.MetaEventResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/meta-events/{metaEventID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves a meta event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meta Events"
                ],
                "summary": "Retrieve a meta event",
                "operationId": "GetMetaEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "meta event id",
                        "name": "metaEventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.MetaEventResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/meta-events/{metaEventID}/resend": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retries a meta event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meta Events"
                ],
                "summary": "Retry meta event",
                "operationId": "ResendMetaEvent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "meta event id",
                        "name": "metaEventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.MetaEventResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/portal-links": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches multiple portal links",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal Links"
                ],
                "summary": "List all portal links",
                "operationId": "LoadPortalLinksPaged",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "The owner ID of the endpoint",
                        "name": "ownerId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "endpoint-1",
                        "description": "The name of the endpoint",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.PortalLinkResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates a portal link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal Links"
                ],
                "summary": "Create a portal link",
                "operationId": "CreatePortalLink",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Portal Link Details",
                        "name": "portallink",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PortalLink"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.PortalLinkResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/portal-links/{portalLinkID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves a portal link by its id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal Links"
                ],
                "summary": "Retrieve a portal link",
                "operationId": "GetPortalLink",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "portal link id",
                        "name": "portalLinkID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.PortalLinkResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint updates a portal link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal Links"
                ],
                "summary": "Update a portal link",
                "operationId": "UpdatePortalLink",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "portal link id",
                        "name": "portalLinkID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Portal Link Details",
                        "name": "portallink",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PortalLink"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.PortalLinkResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/portal-links/{portalLinkID}/revoke": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint revokes a portal link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal Links"
                ],
                "summary": "Revoke a portal link",
                "operationId": "RevokePortalLink",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "portal link id",
                        "name": "portalLinkID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/sources": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches multiple sources",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "List all sources",
                "operationId": "LoadSourcesPaged",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "twitter",
                        "description": "The custom source provider e.g. twitter, shopify",
                        "name": "provider",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "http",
                        "description": "The source type e.g. http, pub_sub",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.SourceResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates a source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Create a source",
                "operationId": "CreateSource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Source Details",
                        "name": "source",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateSource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SourceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/sources/{sourceID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves a source by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Retrieve a source",
                "operationId": "GetSource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Source ID",
                        "name": "sourceID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SourceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint updates a source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Update a source",
                "operationId": "UpdateSource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "source id",
                        "name": "sourceID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Source Details",
                        "name": "source",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateSource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SourceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint deletes a source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Delete a source",
                "operationId": "DeleteSource",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "source id",
                        "name": "sourceID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/subscriptions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint fetches all the subscriptions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "List all subscriptions",
                "operationId": "GetSubscriptions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "next",
                            "prev"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Next",
                            "Prev"
                        ],
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "A list of endpointIDs to filter by",
                        "name": "endpointId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Subscription name to filter by",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JA5MEES38RRK3HTEJC647K",
                        "description": "A pagination cursor to fetch the next page of a list",
                        "name": "next_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 20,
                        "description": "The number of items to return per page",
                        "name": "perPage",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "01H0JATTVCXZK8FRDX1M1JN3QY",
                        "description": "A pagination cursor to fetch the previous page of a list",
                        "name": "prev_page_cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "ASC | DESC",
                        "description": "Sort order, values are ` + "`" + `ASC` + "`" + ` or ` + "`" + `DESC` + "`" + `, defaults to ` + "`" + `DESC` + "`" + `",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.PagedResponse"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "content": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.SubscriptionResponse"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint creates a subscriptions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Create a subscription",
                "operationId": "CreateSubscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Subscription details",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateSubscription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SubscriptionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/subscriptions/test_filter": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint validates that a filter will match a certain payload structure.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Validate subscription filter",
                "operationId": "TestSubscriptionFilter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Filter Details",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TestFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/subscriptions/test_function": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint test runs a transform function against a payload.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Test a subscription function",
                "operationId": "TestSubscriptionFunction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Function Details",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FunctionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.FunctionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/projects/{projectID}/subscriptions/{subscriptionID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint retrieves a single subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Retrieve a subscription",
                "operationId": "GetSubscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "subscription id",
                        "name": "subscriptionID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SubscriptionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint updates a subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Update a subscription",
                "operationId": "UpdateSubscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "subscription id",
                        "name": "subscriptionID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Subscription Details",
                        "name": "subscription",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateSubscription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.SubscriptionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint deletes a subscription",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Delete subscription",
                "operationId": "DeleteSubscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "subscription id",
                        "name": "subscriptionID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handlers.Stub"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "datastore.AlertConfiguration": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "threshold": {
                    "type": "string"
                }
            }
        },
        "datastore.AmqpCredentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "datastore.AmqpPubSubConfig": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/datastore.AmqpCredentials"
                },
                "bindedExchange": {
                    "type": "string"
                },
                "deadLetterExchange": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "queue": {
                    "type": "string"
                },
                "routingKey": {
                    "type": "string"
                },
                "schema": {
                    "type": "string"
                },
                "vhost": {
                    "type": "string"
                }
            }
        },
        "datastore.ApiKey": {
            "type": "object",
            "properties": {
                "header_name": {
                    "type": "string"
                },
                "header_value": {
                    "type": "string"
                }
            }
        },
        "datastore.BasicAuth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "datastore.CLIMetadata": {
            "type": "object",
            "properties": {
                "event_type": {
                    "type": "string"
                },
                "source_id": {
                    "type": "string"
                }
            }
        },
        "datastore.CustomResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "content_type": {
                    "type": "string"
                }
            }
        },
        "datastore.DeliveryAttempt": {
            "type": "object",
            "properties": {
                "api_version": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "endpoint_id": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "http_status": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "msg_id": {
                    "type": "string"
                },
                "request_http_header": {
                    "$ref": "#/definitions/datastore.HttpHeader"
                },
                "response_data": {
                    "type": "string"
                },
                "response_http_header": {
                    "$ref": "#/definitions/datastore.HttpHeader"
                },
                "status": {
                    "type": "boolean"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "datastore.Device": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "endpoint_id": {
                    "type": "string"
                },
                "host_name": {
                    "type": "string"
                },
                "last_seen_at": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/datastore.DeviceStatus"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "datastore.DeviceStatus": {
            "type": "string",
            "enum": [
                "offline",
                "online",
                "disabled"
            ],
            "x-enum-varnames": [
                "DeviceStatusOffline",
                "DeviceStatusOnline",
                "DeviceStatusDisabled"
            ]
        },
        "datastore.EncodingType": {
            "type": "string",
            "enum": [
                "base64",
                "hex"
            ],
            "x-enum-varnames": [
                "Base64Encoding",
                "HexEncoding"
            ]
        },
        "datastore.Endpoint": {
            "type": "object",
            "properties": {
                "advanced_signatures": {
                    "type": "boolean"
                },
                "authentication": {
                    "$ref": "#/definitions/datastore.EndpointAuthentication"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "events": {
                    "type": "integer"
                },
                "http_timeout": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "rate_limit": {
                    "type": "integer"
                },
                "rate_limit_duration": {
                    "type": "integer"
                },
                "secrets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/datastore.Secret"
                    }
                },
                "slack_webhook_url": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/datastore.EndpointStatus"
                },
                "support_email": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "datastore.EndpointAuthentication": {
            "type": "object",
            "properties": {
                "api_key": {
                    "$ref": "#/definitions/datastore.ApiKey"
                },
                "type": {
                    "$ref": "#/definitions/datastore.EndpointAuthenticationType"
                }
            }
        },
        "datastore.EndpointAuthenticationType": {
            "type": "string",
            "enum": [
                "api_key"
            ],
            "x-enum-varnames": [
                "APIKeyAuthentication"
            ]
        },
        "datastore.EndpointStatus": {
            "type": "string",
            "enum": [
                "active",
                "inactive",
                "pending",
                "paused"
            ],
            "x-enum-varnames": [
                "ActiveEndpointStatus",
                "InactiveEndpointStatus",
                "PendingEndpointStatus",
                "PausedEndpointStatus"
            ]
        },
        "datastore.Event": {
            "type": "object",
            "properties": {
                "app_id": {
                    "description": "Deprecated",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "deleted_at": {
                    "type": "string"
                },
                "endpoint_metadata": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/datastore.Endpoint"
                    }
                },
                "endpoints": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "event_type": {
                    "type": "string"
                },
                "headers": {
                    "$ref": "#/definitions/httpheader.HTTPHeader"
                },
                "idempotency_key": {
                    "type": "string"
                },
                "is_duplicate_event": {
                    "type": "boolean"
                },
                "project_id": {
                    "type": "string"
                },
                "raw": {
                    "type": "string"
                },
                "source_id": {
                    "type": "string"
                },
                "source_metadata": {
                    "$ref": "#/definitions/datastore.Source"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url_query_params": {
                    "type": "string"
                }
            }
        },
        "datastore.EventDeliveryStatus": {
            "type": "string",
            "enum": [
                "Scheduled",
                "Processing",
                "Discarded",
                "Failure",
                "Success",
                "Retry"
            ],
            "x-enum-varnames": [
                "ScheduledEventStatus",
                "ProcessingEventStatus",
                "DiscardedEventStatus",
                "FailureEventStatus",
                "SuccessEventStatus",
                "RetryEventStatus"
            ]
        },
        "datastore.FilterConfiguration": {
            "type": "object",
            "properties": {
                "event_types": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "filter": {
                    "$ref": "#/definitions/datastore.FilterSchema"
                }
            }
        },
        "datastore.FilterSchema": {
            "type": "object",
            "properties": {
                "body": {
                    "$ref": "#/definitions/datastore.M"
                },
                "headers": {
                    "$ref": "#/definitions/datastore.M"
                }
            }
        },
        "datastore.GooglePubSubConfig": {
            "type": "object",
            "properties": {
                "project_id": {
                    "type": "string"
                },
                "service_account": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "subscription_id": {
                    "type": "string"
                }
            }
        },
        "datastore.HMac": {
            "type": "object",
            "properties": {
                "encoding": {
                    "$ref": "#/definitions/datastore.EncodingType"
                },
                "hash": {
                    "type": "string"
                },
                "header": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "datastore.HttpHeader": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "datastore.KafkaAuth": {
            "type": "object",
            "properties": {
                "hash": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tls": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "datastore.KafkaPubSubConfig": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/datastore.KafkaAuth"
                },
                "brokers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "consumer_group_id": {
                    "type": "string"
                },
                "topic_name": {
                    "type": "string"
                }
            }
        },
        "datastore.M": {
            "type": "object",
            "additionalProperties": true
        },
        "datastore.MetaEventAttempt": {
            "type": "object",
            "properties": {
                "request_http_header": {
                    "$ref": "#/definitions/datastore.HttpHeader"
                },
                "response_data": {
                    "type": "string"
                },
                "response_http_header": {
                    "$ref": "#/definitions/datastore.HttpHeader"
                }
            }
        },
        "datastore.Metadata": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data to be sent to endpoint.",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "interval_seconds": {
                    "type": "integer"
                },
                "next_send_time": {
                    "type": "string"
                },
                "num_trials": {
                    "description": "NumTrials: number of times we have tried to deliver this Event to\nan application",
                    "type": "integer"
                },
                "raw": {
                    "type": "string"
                },
                "retry_limit": {
                    "type": "integer"
                },
                "strategy": {
                    "$ref": "#/definitions/datastore.StrategyProvider"
                }
            }
        },
        "datastore.PageDirection": {
            "type": "string",
            "enum": [
                "next",
                "prev"
            ],
            "x-enum-varnames": [
                "Next",
                "Prev"
            ]
        },
        "datastore.PaginationData": {
            "type": "object",
            "properties": {
                "has_next_page": {
                    "type": "boolean"
                },
                "has_prev_page": {
                    "type": "boolean"
                },
                "next_page_cursor": {
                    "type": "string"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page_cursor": {
                    "type": "string"
                }
            }
        },
        "datastore.ProviderConfig": {
            "type": "object",
            "properties": {
                "twitter": {
                    "$ref": "#/definitions/datastore.TwitterProviderConfig"
                }
            }
        },
        "datastore.PubSubConfig": {
            "type": "object",
            "properties": {
                "amqp": {
                    "$ref": "#/definitions/datastore.AmqpPubSubConfig"
                },
                "google": {
                    "$ref": "#/definitions/datastore.GooglePubSubConfig"
                },
                "kafka": {
                    "$ref": "#/definitions/datastore.KafkaPubSubConfig"
                },
                "sqs": {
                    "$ref": "#/definitions/datastore.SQSPubSubConfig"
                },
                "type": {
                    "$ref": "#/definitions/datastore.PubSubType"
                },
                "workers": {
                    "type": "integer"
                }
            }
        },
        "datastore.PubSubType": {
            "type": "string",
            "enum": [
                "sqs",
                "google",
                "kafka",
                "amqp"
            ],
            "x-enum-varnames": [
                "SqsPubSub",
                "GooglePubSub",
                "KafkaPubSub",
                "AmqpPubSub"
            ]
        },
        "datastore.RateLimitConfiguration": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "duration": {
                    "type": "integer"
                }
            }
        },
        "datastore.RetryConfiguration": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "integer"
                },
                "retry_count": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/datastore.StrategyProvider"
                }
            }
        },
        "datastore.SQSPubSubConfig": {
            "type": "object",
            "properties": {
                "access_key_id": {
                    "type": "string"
                },
                "default_region": {
                    "type": "string"
                },
                "queue_name": {
                    "type": "string"
                },
                "secret_key": {
                    "type": "string"
                }
            }
        },
        "datastore.Secret": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "datastore.Source": {
            "type": "object",
            "properties": {
                "body_function": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "custom_response": {
                    "$ref": "#/definitions/datastore.CustomResponse"
                },
                "deleted_at": {
                    "type": "string"
                },
                "forward_headers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "header_function": {
                    "type": "string"
                },
                "idempotency_keys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_disabled": {
                    "type": "boolean"
                },
                "mask_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "provider": {
                    "$ref": "#/definitions/datastore.SourceProvider"
                },
                "provider_config": {
                    "$ref": "#/definitions/datastore.ProviderConfig"
                },
                "pub_sub": {
                    "$ref": "#/definitions/datastore.PubSubConfig"
                },
                "type": {
                    "$ref": "#/definitions/datastore.SourceType"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "verifier": {
                    "$ref": "#/definitions/datastore.VerifierConfig"
                }
            }
        },
        "datastore.SourceProvider": {
            "type": "string",
            "enum": [
                "github",
                "twitter",
                "shopify"
            ],
            "x-enum-varnames": [
                "GithubSourceProvider",
                "TwitterSourceProvider",
                "ShopifySourceProvider"
            ]
        },
        "datastore.SourceType": {
            "type": "string",
            "enum": [
                "http",
                "rest_api",
                "pub_sub",
                "db_change_stream"
            ],
            "x-enum-varnames": [
                "HTTPSource",
                "RestApiSource",
                "PubSubSource",
                "DBChangeStream"
            ]
        },
        "datastore.StrategyProvider": {
            "type": "string",
            "enum": [
                "linear",
                "exponential"
            ],
            "x-enum-varnames": [
                "LinearStrategyProvider",
                "ExponentialStrategyProvider"
            ]
        },
        "datastore.SubscriptionType": {
            "type": "string",
            "enum": [
                "cli",
                "api"
            ],
            "x-enum-varnames": [
                "SubscriptionTypeCLI",
                "SubscriptionTypeAPI"
            ]
        },
        "datastore.TwitterProviderConfig": {
            "type": "object",
            "properties": {
                "crc_verified_at": {
                    "type": "string"
                }
            }
        },
        "datastore.VerifierConfig": {
            "type": "object",
            "properties": {
                "api_key": {
                    "$ref": "#/definitions/datastore.ApiKey"
                },
                "basic_auth": {
                    "$ref": "#/definitions/datastore.BasicAuth"
                },
                "hmac": {
                    "$ref": "#/definitions/datastore.HMac"
                },
                "type": {
                    "$ref": "#/definitions/datastore.VerifierType"
                }
            }
        },
        "datastore.VerifierType": {
            "type": "string",
            "enum": [
                "noop",
                "hmac",
                "basic_auth",
                "api_key"
            ],
            "x-enum-varnames": [
                "NoopVerifier",
                "HMacVerifier",
                "BasicAuthVerifier",
                "APIKeyVerifier"
            ]
        },
        "handlers.Stub": {
            "type": "object"
        },
        "httpheader.HTTPHeader": {
            "type": "object",
            "additionalProperties": {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        },
        "models.AlertConfiguration": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "Count",
                    "type": "integer"
                },
                "threshold": {
                    "description": "Threshold",
                    "type": "string"
                }
            }
        },
        "models.AmqpAuth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "models.AmqpExchange": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "routingKey": {
                    "type": "string"
                }
            }
        },
        "models.AmqpPubSubconfig": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/models.AmqpAuth"
                },
                "bindExchange": {
                    "$ref": "#/definitions/models.AmqpExchange"
                },
                "deadLetterExchange": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "queue": {
                    "type": "string"
                },
                "schema": {
                    "type": "string"
                },
                "vhost": {
                    "type": "string"
                }
            }
        },
        "models.ApiKey": {
            "type": "object",
            "properties": {
                "header_name": {
                    "type": "string"
                },
                "header_value": {
                    "type": "string"
                }
            }
        },
        "models.BasicAuth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.BroadcastEvent": {
            "type": "object",
            "properties": {
                "custom_headers": {
                    "description": "Specifies custom headers you want convoy to add when the event is dispatched to your endpoint",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "event_type": {
                    "description": "Event Type is used for filtering and debugging e.g invoice.paid",
                    "type": "string"
                },
                "idempotency_key": {
                    "description": "Specify a key for event deduplication",
                    "type": "string"
                }
            }
        },
        "models.CreateEndpoint": {
            "type": "object",
            "properties": {
                "advanced_signatures": {
                    "description": "Convoy supports two [signature formats](https://getconvoy.io/docs/manual/signatures)\n-- simple or advanced. If left unspecified, we default to false.",
                    "type": "boolean"
                },
                "appID": {
                    "description": "Deprecated but necessary for backward compatibility",
                    "type": "string"
                },
                "authentication": {
                    "description": "This is used to define any custom authentication required by the endpoint. This\nshouldn't be needed often because webhook endpoints usually should be exposed to\nthe internet.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.EndpointAuthentication"
                        }
                    ]
                },
                "description": {
                    "description": "Human-readable description of the endpoint. Think of this as metadata describing\nthe endpoint",
                    "type": "string"
                },
                "http_timeout": {
                    "description": "Define endpoint http timeout in seconds.",
                    "type": "integer"
                },
                "is_disabled": {
                    "description": "This is used to manually enable/disable the endpoint.",
                    "type": "boolean"
                },
                "name": {
                    "description": "Endpoint name.",
                    "type": "string"
                },
                "owner_id": {
                    "description": "The OwnerID is used to group more than one endpoint together to achieve\n[fanout](https://getconvoy.io/docs/manual/endpoints#Endpoint%20Owner%20ID)",
                    "type": "string"
                },
                "rate_limit": {
                    "description": "Rate limit is the total number of requests to be sent to an endpoint in\nthe time duration specified in RateLimitDuration",
                    "type": "integer"
                },
                "rate_limit_duration": {
                    "description": "Rate limit duration specifies the time range for the rate limit.",
                    "type": "integer"
                },
                "secret": {
                    "description": "Endpoint's webhook secret. If not provided, Convoy autogenerates one for the endpoint.",
                    "type": "string"
                },
                "slack_webhook_url": {
                    "description": "Slack webhook URL is an alternative method to support email where endpoint developers\ncan receive failure notifications on a slack channel.",
                    "type": "string"
                },
                "support_email": {
                    "description": "Endpoint developers support email. This is used for communicating endpoint state\nchanges. You should always turn this on when disabling endpoints are enabled.",
                    "type": "string"
                },
                "url": {
                    "description": "URL is the endpoint's URL prefixed with https. non-https urls are currently\nnot supported.",
                    "type": "string"
                }
            }
        },
        "models.CreateEvent": {
            "type": "object",
            "properties": {
                "app_id": {
                    "description": "Deprecated but necessary for backward compatibility.",
                    "type": "string"
                },
                "custom_headers": {
                    "description": "Specifies custom headers you want convoy to add when the event is dispatched to your endpoint",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "object"
                },
                "endpoint_id": {
                    "description": "Specifies the endpoint to send this event to.",
                    "type": "string"
                },
                "event_type": {
                    "description": "Event Type is used for filtering and debugging e.g invoice.paid",
                    "type": "string"
                },
                "idempotency_key": {
                    "description": "Specify a key for event deduplication",
                    "type": "string"
                }
            }
        },
        "models.CreateSource": {
            "type": "object",
            "properties": {
                "body_function": {
                    "description": "Function is a javascript function used to mutate the payload\nimmediately after ingesting an event",
                    "type": "string"
                },
                "custom_response": {
                    "description": "Custom response is used to define a custom response for incoming\nwebhooks project sources only.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CustomResponse"
                        }
                    ]
                },
                "header_function": {
                    "description": "Function is a javascript function used to mutate the headers\nimmediately after ingesting an event",
                    "type": "string"
                },
                "idempotency_keys": {
                    "description": "IdempotencyKeys are used to specify parts of a webhook request to uniquely\nidentify the event in an incoming webhooks project.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Source name.",
                    "type": "string"
                },
                "provider": {
                    "description": "Use this to specify one of our predefined source types.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datastore.SourceProvider"
                        }
                    ]
                },
                "pub_sub": {
                    "description": "PubSub are used to specify message broker sources for outgoing\nwebhooks projects.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.PubSubConfig"
                        }
                    ]
                },
                "type": {
                    "description": "Source Type.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datastore.SourceType"
                        }
                    ]
                },
                "verifier": {
                    "description": "Verifiers are used to verify webhook events ingested in incoming\nwebhooks projects.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.VerifierConfig"
                        }
                    ]
                }
            }
        },
        "models.CreateSubscription": {
            "type": "object",
            "properties": {
                "alert_config": {
                    "description": "Alert configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.AlertConfiguration"
                        }
                    ]
                },
                "app_id": {
                    "description": "Deprecated but necessary for backward compatibility",
                    "type": "string"
                },
                "endpoint_id": {
                    "description": "Destination endpoint ID",
                    "type": "string"
                },
                "filter_config": {
                    "description": "Filter configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.FilterConfiguration"
                        }
                    ]
                },
                "function": {
                    "description": "Convoy supports mutating your request payload using a js function. Use this field\nto specify a ` + "`" + `transform` + "`" + ` function for this purpose. See this[https://docs.getconvoy.io/product-manual/subscriptions#functions] for more",
                    "type": "string"
                },
                "name": {
                    "description": "Subscription Nme",
                    "type": "string"
                },
                "rate_limit_config": {
                    "description": "Rate limit configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RateLimitConfiguration"
                        }
                    ]
                },
                "retry_config": {
                    "description": "Retry configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RetryConfiguration"
                        }
                    ]
                },
                "source_id": {
                    "description": "Source Id",
                    "type": "string"
                }
            }
        },
        "models.CustomResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "content_type": {
                    "type": "string"
                }
            }
        },
        "models.DynamicEvent": {
            "type": "object",
            "properties": {
                "custom_headers": {
                    "description": "Specifies custom headers you want convoy to add when the event is dispatched to your endpoint",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "event_type": {
                    "description": "Event Type is used for filtering and debugging e.g invoice.paid",
                    "type": "string"
                },
                "event_types": {
                    "description": "A list of event types for the subscription filter config",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "idempotency_key": {
                    "description": "Specify a key for event deduplication",
                    "type": "string"
                },
                "secret": {
                    "description": "Endpoint's webhook secret. If not provided, Convoy autogenerates one for the endpoint.",
                    "type": "string"
                },
                "url": {
                    "description": "URL is the endpoint's URL prefixed with https. non-https urls are currently\nnot supported.",
                    "type": "string"
                }
            }
        },
        "models.EndpointAuthentication": {
            "type": "object",
            "properties": {
                "api_key": {
                    "$ref": "#/definitions/models.ApiKey"
                },
                "type": {
                    "$ref": "#/definitions/datastore.EndpointAuthenticationType"
                }
            }
        },
        "models.EndpointResponse": {
            "type": "object",
            "properties": {
                "advanced_signatures": {
                    "type": "boolean"
                },
                "authentication": {
                    "$ref": "#/definitions/datastore.EndpointAuthentication"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "events": {
                    "type": "integer"
                },
                "http_timeout": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "rate_limit": {
                    "type": "integer"
                },
                "rate_limit_duration": {
                    "type": "integer"
                },
                "secrets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/datastore.Secret"
                    }
                },
                "slack_webhook_url": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/datastore.EndpointStatus"
                },
                "support_email": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.EventDeliveryResponse": {
            "type": "object",
            "properties": {
                "cli_metadata": {
                    "$ref": "#/definitions/datastore.CLIMetadata"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                },
                "device_metadata": {
                    "$ref": "#/definitions/datastore.Device"
                },
                "endpoint_id": {
                    "type": "string"
                },
                "endpoint_metadata": {
                    "$ref": "#/definitions/datastore.Endpoint"
                },
                "event_id": {
                    "type": "string"
                },
                "event_metadata": {
                    "$ref": "#/definitions/datastore.Event"
                },
                "event_type": {
                    "type": "string"
                },
                "headers": {
                    "$ref": "#/definitions/httpheader.HTTPHeader"
                },
                "idempotency_key": {
                    "type": "string"
                },
                "latency": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/datastore.Metadata"
                },
                "project_id": {
                    "type": "string"
                },
                "source_metadata": {
                    "$ref": "#/definitions/datastore.Source"
                },
                "status": {
                    "$ref": "#/definitions/datastore.EventDeliveryStatus"
                },
                "subscription_id": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url_query_params": {
                    "type": "string"
                }
            }
        },
        "models.EventResponse": {
            "type": "object",
            "properties": {
                "app_id": {
                    "description": "Deprecated",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "deleted_at": {
                    "type": "string"
                },
                "endpoint_metadata": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/datastore.Endpoint"
                    }
                },
                "endpoints": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "event_type": {
                    "type": "string"
                },
                "headers": {
                    "$ref": "#/definitions/httpheader.HTTPHeader"
                },
                "idempotency_key": {
                    "type": "string"
                },
                "is_duplicate_event": {
                    "type": "boolean"
                },
                "project_id": {
                    "type": "string"
                },
                "raw": {
                    "type": "string"
                },
                "source_id": {
                    "type": "string"
                },
                "source_metadata": {
                    "$ref": "#/definitions/datastore.Source"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url_query_params": {
                    "type": "string"
                }
            }
        },
        "models.ExpireSecret": {
            "type": "object",
            "properties": {
                "expiration": {
                    "description": "Amount of time to wait before expiring the old endpoint secret.\nIf AdvancedSignatures is turned on for the project, signatures for both secrets will be generated up until\nthe old signature is expired.",
                    "type": "integer"
                },
                "secret": {
                    "description": "New Endpoint secret value.",
                    "type": "string"
                }
            }
        },
        "models.FS": {
            "type": "object",
            "properties": {
                "body": {
                    "$ref": "#/definitions/datastore.M"
                },
                "headers": {
                    "$ref": "#/definitions/datastore.M"
                }
            }
        },
        "models.FanoutEvent": {
            "type": "object",
            "properties": {
                "custom_headers": {
                    "description": "Specifies custom headers you want convoy to add when the event is dispatched to your endpoint",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "Data is an arbitrary JSON value that gets sent as the body of the\nwebhook to the endpoints",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "event_type": {
                    "description": "Event Type is used for filtering and debugging e.g invoice.paid",
                    "type": "string"
                },
                "idempotency_key": {
                    "description": "Specify a key for event deduplication",
                    "type": "string"
                },
                "owner_id": {
                    "description": "Used for fanout, sends this event to all endpoints with this OwnerID.",
                    "type": "string"
                }
            }
        },
        "models.FilterConfiguration": {
            "type": "object",
            "properties": {
                "event_types": {
                    "description": "List of event types that the subscription should match",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "filter": {
                    "description": "Body \u0026 Header filters",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.FS"
                        }
                    ]
                }
            }
        },
        "models.FilterSchema": {
            "type": "object",
            "properties": {
                "body": {},
                "header": {}
            }
        },
        "models.FunctionRequest": {
            "type": "object",
            "properties": {
                "function": {
                    "type": "string"
                },
                "payload": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.FunctionResponse": {
            "type": "object",
            "properties": {
                "log": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "payload": {}
            }
        },
        "models.GooglePubSubConfig": {
            "type": "object",
            "properties": {
                "project_id": {
                    "type": "string"
                },
                "service_account": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "subscription_id": {
                    "type": "string"
                }
            }
        },
        "models.HMac": {
            "type": "object",
            "properties": {
                "encoding": {
                    "$ref": "#/definitions/datastore.EncodingType"
                },
                "hash": {
                    "type": "string"
                },
                "header": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "models.IDs": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "A list of event delivery IDs to forcefully resend.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.KafkaAuth": {
            "type": "object",
            "properties": {
                "hash": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tls": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.KafkaPubSubConfig": {
            "type": "object",
            "properties": {
                "auth": {
                    "$ref": "#/definitions/models.KafkaAuth"
                },
                "brokers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "consumer_group_id": {
                    "type": "string"
                },
                "topic_name": {
                    "type": "string"
                }
            }
        },
        "models.MetaEventResponse": {
            "type": "object",
            "properties": {
                "attempt": {
                    "$ref": "#/definitions/datastore.MetaEventAttempt"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "event_type": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/datastore.Metadata"
                },
                "project_id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/datastore.EventDeliveryStatus"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.PagedResponse": {
            "type": "object",
            "properties": {
                "content": {},
                "pagination": {
                    "$ref": "#/definitions/datastore.PaginationData"
                }
            }
        },
        "models.PortalLink": {
            "type": "object",
            "properties": {
                "can_manage_endpoint": {
                    "description": "Specify whether endpoint management can be done through the Portal Link UI",
                    "type": "boolean"
                },
                "endpoints": {
                    "description": "IDs of endpoints in this portal link",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Portal Link Name",
                    "type": "string"
                },
                "owner_id": {
                    "description": "Alternatively specify OwnerID, the portal link will inherit all the endpoints with this owner ID",
                    "type": "string"
                }
            }
        },
        "models.PortalLinkResponse": {
            "type": "object",
            "properties": {
                "can_manage_endpoint": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "endpoint_count": {
                    "type": "integer"
                },
                "endpoints": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "endpoints_metadata": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/datastore.Endpoint"
                    }
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.PubSubConfig": {
            "type": "object",
            "properties": {
                "amqp": {
                    "$ref": "#/definitions/models.AmqpPubSubconfig"
                },
                "google": {
                    "$ref": "#/definitions/models.GooglePubSubConfig"
                },
                "kafka": {
                    "$ref": "#/definitions/models.KafkaPubSubConfig"
                },
                "sqs": {
                    "$ref": "#/definitions/models.SQSPubSubConfig"
                },
                "type": {
                    "$ref": "#/definitions/datastore.PubSubType"
                },
                "workers": {
                    "type": "integer"
                }
            }
        },
        "models.RateLimitConfiguration": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "duration": {
                    "type": "integer"
                }
            }
        },
        "models.RetryConfiguration": {
            "type": "object",
            "properties": {
                "duration": {
                    "description": "Used to specify a valid Go time duration e.g 10s, 1h3m for how long to wait between event delivery retries",
                    "type": "string"
                },
                "interval_seconds": {
                    "description": "Used to specify a time in seconds for how long to wait between event delivery retries,",
                    "type": "integer"
                },
                "retry_count": {
                    "description": "Used to specify the max number of retries",
                    "type": "integer"
                },
                "type": {
                    "description": "Retry Strategy type",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datastore.StrategyProvider"
                        }
                    ]
                }
            }
        },
        "models.SQSPubSubConfig": {
            "type": "object",
            "properties": {
                "access_key_id": {
                    "type": "string"
                },
                "default_region": {
                    "type": "string"
                },
                "queue_name": {
                    "type": "string"
                },
                "secret_key": {
                    "type": "string"
                }
            }
        },
        "models.SourceResponse": {
            "type": "object",
            "properties": {
                "body_function": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "custom_response": {
                    "$ref": "#/definitions/datastore.CustomResponse"
                },
                "deleted_at": {
                    "type": "string"
                },
                "forward_headers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "header_function": {
                    "type": "string"
                },
                "idempotency_keys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_disabled": {
                    "type": "boolean"
                },
                "mask_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "provider": {
                    "$ref": "#/definitions/datastore.SourceProvider"
                },
                "provider_config": {
                    "$ref": "#/definitions/datastore.ProviderConfig"
                },
                "pub_sub": {
                    "$ref": "#/definitions/datastore.PubSubConfig"
                },
                "type": {
                    "$ref": "#/definitions/datastore.SourceType"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "verifier": {
                    "$ref": "#/definitions/datastore.VerifierConfig"
                }
            }
        },
        "models.SubscriptionResponse": {
            "type": "object",
            "properties": {
                "alert_config": {
                    "description": "subscription config",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datastore.AlertConfiguration"
                        }
                    ]
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "device_metadata": {
                    "$ref": "#/definitions/datastore.Device"
                },
                "endpoint_metadata": {
                    "$ref": "#/definitions/datastore.Endpoint"
                },
                "filter_config": {
                    "$ref": "#/definitions/datastore.FilterConfiguration"
                },
                "function": {
                    "$ref": "#/definitions/null.String"
                },
                "name": {
                    "type": "string"
                },
                "rate_limit_config": {
                    "$ref": "#/definitions/datastore.RateLimitConfiguration"
                },
                "retry_config": {
                    "$ref": "#/definitions/datastore.RetryConfiguration"
                },
                "source_metadata": {
                    "$ref": "#/definitions/datastore.Source"
                },
                "type": {
                    "$ref": "#/definitions/datastore.SubscriptionType"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.TestFilter": {
            "type": "object",
            "properties": {
                "request": {
                    "description": "Same Request \u0026 Headers",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.FilterSchema"
                        }
                    ]
                },
                "schema": {
                    "description": "Sample test schema",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.FilterSchema"
                        }
                    ]
                }
            }
        },
        "models.UpdateCustomResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "content_type": {
                    "type": "string"
                }
            }
        },
        "models.UpdateEndpoint": {
            "type": "object",
            "properties": {
                "advanced_signatures": {
                    "description": "Convoy supports two [signature formats](https://getconvoy.io/docs/manual/signatures)\n-- simple or advanced. If left unspecified, we default to false.",
                    "type": "boolean"
                },
                "authentication": {
                    "description": "This is used to define any custom authentication required by the endpoint. This\nshouldn't be needed often because webhook endpoints usually should be exposed to\nthe internet.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.EndpointAuthentication"
                        }
                    ]
                },
                "description": {
                    "description": "Human-readable description of the endpoint. Think of this as metadata describing\nthe endpoint",
                    "type": "string"
                },
                "http_timeout": {
                    "description": "Define endpoint http timeout in seconds.",
                    "type": "integer"
                },
                "is_disabled": {
                    "description": "This is used to manually enable/disable the endpoint.",
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "description": "The OwnerID is used to group more than one endpoint together to achieve\n[fanout](https://getconvoy.io/docs/manual/endpoints#Endpoint%20Owner%20ID)",
                    "type": "string"
                },
                "rate_limit": {
                    "description": "Rate limit is the total number of requests to be sent to an endpoint in\nthe time duration specified in RateLimitDuration",
                    "type": "integer"
                },
                "rate_limit_duration": {
                    "description": "Rate limit duration specifies the time range for the rate limit.",
                    "type": "integer"
                },
                "secret": {
                    "description": "Endpoint's webhook secret. If not provided, Convoy autogenerates one for the endpoint.",
                    "type": "string"
                },
                "slack_webhook_url": {
                    "description": "Slack webhook URL is an alternative method to support email where endpoint developers\ncan receive failure notifications on a slack channel.",
                    "type": "string"
                },
                "support_email": {
                    "description": "Endpoint developers support email. This is used for communicating endpoint state\nchanges. You should always turn this on when disabling endpoints are enabled.",
                    "type": "string"
                },
                "url": {
                    "description": "URL is the endpoint's URL prefixed with https. non-https urls are currently\nnot supported.",
                    "type": "string"
                }
            }
        },
        "models.UpdateSource": {
            "type": "object",
            "properties": {
                "body_function": {
                    "description": "Function is a javascript function used to mutate the payload\nimmediately after ingesting an event",
                    "type": "string"
                },
                "custom_response": {
                    "description": "Custom response is used to define a custom response for incoming\nwebhooks project sources only.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.UpdateCustomResponse"
                        }
                    ]
                },
                "forward_headers": {
                    "description": "Soecfy header you want convoy to save from the ingest request and forward to your endpoints when the event is dispatched.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "header_function": {
                    "description": "Function is a javascript function used to mutate the headers\nimmediately after ingesting an event",
                    "type": "string"
                },
                "idempotency_keys": {
                    "description": "IdempotencyKeys are used to specify parts of a webhook request to uniquely\nidentify the event in an incoming webhooks project.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_disabled": {
                    "description": "This is used to manually enable/disable the source.",
                    "type": "boolean"
                },
                "name": {
                    "description": "Source name.",
                    "type": "string"
                },
                "pub_sub": {
                    "description": "PubSub are used to specify message broker sources for outgoing\nwebhooks projects, you only need to specify this when the source type is ` + "`" + `pub_sub` + "`" + `.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.PubSubConfig"
                        }
                    ]
                },
                "type": {
                    "description": "Source Type.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datastore.SourceType"
                        }
                    ]
                },
                "verifier": {
                    "description": "Verifiers are used to verify webhook events ingested in incoming\nwebhooks projects.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.VerifierConfig"
                        }
                    ]
                }
            }
        },
        "models.UpdateSubscription": {
            "type": "object",
            "properties": {
                "alert_config": {
                    "description": "Alert configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.AlertConfiguration"
                        }
                    ]
                },
                "app_id": {
                    "description": "Deprecated but necessary for backward compatibility",
                    "type": "string"
                },
                "endpoint_id": {
                    "description": "Destination endpoint ID",
                    "type": "string"
                },
                "filter_config": {
                    "description": "Filter configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.FilterConfiguration"
                        }
                    ]
                },
                "function": {
                    "description": "Convoy supports mutating your request payload using a js function. Use this field\nto specify a ` + "`" + `transform` + "`" + ` function for this purpose. See this[https://docs.getconvoy.io/product-manual/subscriptions#functions] for more",
                    "type": "string"
                },
                "name": {
                    "description": "Subscription Nme",
                    "type": "string"
                },
                "rate_limit_config": {
                    "description": "Rate limit configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RateLimitConfiguration"
                        }
                    ]
                },
                "retry_config": {
                    "description": "Retry configuration",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RetryConfiguration"
                        }
                    ]
                },
                "source_id": {
                    "description": "Source Id",
                    "type": "string"
                }
            }
        },
        "models.VerifierConfig": {
            "type": "object",
            "properties": {
                "api_key": {
                    "$ref": "#/definitions/models.ApiKey"
                },
                "basic_auth": {
                    "$ref": "#/definitions/models.BasicAuth"
                },
                "hmac": {
                    "$ref": "#/definitions/models.HMac"
                },
                "type": {
                    "$ref": "#/definitions/datastore.VerifierType"
                }
            }
        },
        "null.String": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
        },
        "util.ServerResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "Organisation related APIs",
            "name": "Organisations"
        },
        {
            "description": "Subscription related APIs",
            "name": "Subscriptions"
        },
        {
            "description": "Endpoint related APIs",
            "name": "Endpoints"
        },
        {
            "description": "Event related APIs",
            "name": "Events"
        },
        {
            "description": "Source related APIs",
            "name": "Sources"
        },
        {
            "description": "EventDelivery related APIs",
            "name": "Event Deliveries"
        },
        {
            "description": "Delivery Attempt related APIs",
            "name": "Delivery Attempts"
        },
        {
            "description": "Project related APIs",
            "name": "Projects"
        },
        {
            "description": "Portal Links related APIs",
            "name": "Portal Links"
        },
        {
            "description": "Meta Events related APIs",
            "name": "Meta Events"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "24.1.4",
	Host:             "dashboard.getconvoy.io",
	BasePath:         "/api",
	Schemes:          []string{"https"},
	Title:            "Convoy API Reference",
	Description:      "Convoy is a fast and secure webhooks proxy. This document contains datastore.s API specification.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
