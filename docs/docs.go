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
        "/api/admin/adminrights/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "remove admin rights",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "give admin rights",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/api/admin/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/api/profile/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "delete profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/api/profile/changepasswd": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "description": "password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/api/profile/changeusername": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Change Username",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangeUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/auth/adminlogin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "LogIn for admin",
                "parameters": [
                    {
                        "description": "email password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "LogIn",
                "parameters": [
                    {
                        "description": "email password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/auth/resetpasswd": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ResetPasswd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponce"
                        }
                    }
                }
            }
        },
        "/auth/steam": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "go to steam",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/auth/steam/callback": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "get from steam",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.errorResponce": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ChangePasswd": {
            "type": "object",
            "required": [
                "new_password",
                "new_password2",
                "password"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "new_password2": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.ChangeUserInfo": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "device_token": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.ResetPasswd": {
            "type": "object",
            "required": [
                "email",
                "password",
                "password2"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password2": {
                    "type": "string"
                }
            }
        },
        "models.SignUp": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "device_token": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
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
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "54.83.180.49:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Ordens App",
	Description:      "API server for Ordens application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
