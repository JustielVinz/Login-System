// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "The user will log in the data here",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log in"
                ],
                "summary": "Login the user and fetch the data",
                "operationId": "log in UI",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/struct_test.AdminAcc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\"\t\"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "The user will create an Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Creating New Account"
                ],
                "summary": "Set up the Account",
                "operationId": "createAccount",
                "parameters": [
                    {
                        "description": "Create  Account",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/struct_test.AdminAcc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/struct_test.AdminAcc"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/struct_test.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "post": {
                "description": "Gather the specific data that user wants",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameter"
                ],
                "summary": "Fecth the data",
                "operationId": "Parameter ID",
                "parameters": [
                    {
                        "description": "Fetch Success",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/struct_test.AdminAcc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\"\t\"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "struct_test.AdminAcc": {
            "description": "struct for getting the credentials",
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "struct_test.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.16.2",
	Host:             "",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Log In System",
	Description:      "Log in Code for Go session",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
