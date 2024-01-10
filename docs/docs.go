// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ly Tu",
            "url": "https://github.com/timly278/pomodoro",
            "email": "timly278@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Loggin user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Loggin user",
                "parameters": [
                    {
                        "description": "user login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "403": {
                        "description": "password does not match",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "406": {
                        "description": "email has not verified",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "description": "Refresh access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh access token",
                "parameters": [
                    {
                        "description": "refresh access token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Refresh token is unauthorized",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create new user and send verification code to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "New user registers",
                "parameters": [
                    {
                        "description": "Create new user",
                        "name": "NewUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "406": {
                        "description": "email spam, verification code has created and sent",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "email existed",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/auth/send-emailverification": {
            "post": {
                "description": "Send verification code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Send email for verification code",
                "parameters": [
                    {
                        "description": "send code",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.SendCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "406": {
                        "description": "email spam, verification code has created and sent",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/auth/verify-code": {
            "post": {
                "description": "Verify code that sent over email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify email verification code",
                "parameters": [
                    {
                        "description": "verify code",
                        "name": "Email\u0026Code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.VerificationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "email has been verified successfully",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/jobs/accessed-days": {
            "get": {
                "description": "Get total days user has accessed on pomodoro from date to date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get Days Accessed",
                "parameters": [
                    {
                        "description": "Get days",
                        "name": "DateToDate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.GetStatisticRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get days accessed",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/jobs/focused-minutes": {
            "get": {
                "description": "Get total minutes user has spent on pomodoro from date to date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get Minutes Focused",
                "parameters": [
                    {
                        "description": "Get minutes",
                        "name": "DateToDate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.GetStatisticRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get minutes focused",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/jobs/pomodoros": {
            "get": {
                "description": "Get pomodoros from date to date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "List Pomodoros",
                "parameters": [
                    {
                        "description": "Get pomodoros",
                        "name": "GetPomos",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.GetPomodorosRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Pomodoro"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new pomodoro",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Create New Pomodoro",
                "parameters": [
                    {
                        "description": "New pomodoro",
                        "name": "NewPomo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.CreatePomodoroRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Pomodoro"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not found user_id or type_id",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/jobs/types": {
            "get": {
                "description": "Get all pomodoros type of this user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get Pomodoro Types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "Update pomodoros type of this user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Update Pomodoro Types",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Type ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update pomo type",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.CreateNewTypeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update type successfully",
                        "schema": {
                            "$ref": "#/definitions/db.Type"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not found user_id or type_id",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new pomodoro types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Create New Pomodoro Types",
                "parameters": [
                    {
                        "description": "New pomodoro type",
                        "name": "NewPomoType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.CreateNewTypeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/jobs/update-user-setting": {
            "put": {
                "description": "Update user setting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Update User Setting",
                "parameters": [
                    {
                        "description": "Update user setting",
                        "name": "userSetting",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.UpdateUserSettingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update user setting successfully",
                        "schema": {
                            "$ref": "#/definitions/response.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal serever error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Pomodoro": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "focus_degree": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "task_id": {
                    "$ref": "#/definitions/sql.NullInt64"
                },
                "type_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "db.Type": {
            "type": "object",
            "properties": {
                "autostart_break": {
                    "type": "boolean"
                },
                "color": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "goalperday": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "longbreak": {
                    "type": "integer"
                },
                "longbreakinterval": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "shortbreak": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "delivery.CreateNewTypeRequest": {
            "type": "object",
            "required": [
                "autostart_break",
                "color",
                "duration",
                "goal_per_day",
                "longbreak",
                "longbreakinterval",
                "name",
                "shortbreak"
            ],
            "properties": {
                "autostart_break": {
                    "type": "boolean"
                },
                "color": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer",
                    "minimum": 1
                },
                "goal_per_day": {
                    "type": "integer",
                    "minimum": 1
                },
                "longbreak": {
                    "type": "integer",
                    "minimum": 1
                },
                "longbreakinterval": {
                    "type": "integer",
                    "minimum": 1
                },
                "name": {
                    "type": "string"
                },
                "shortbreak": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "delivery.CreatePomodoroRequest": {
            "type": "object",
            "required": [
                "focus_degree",
                "type_id"
            ],
            "properties": {
                "focus_degree": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "task_id": {
                    "type": "integer"
                },
                "type_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "delivery.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "delivery.GetPomodorosRequest": {
            "type": "object",
            "required": [
                "from_date",
                "limit",
                "page",
                "to_date"
            ],
            "properties": {
                "from_date": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer",
                    "minimum": 1
                },
                "page": {
                    "type": "integer",
                    "minimum": 1
                },
                "to_date": {
                    "type": "string"
                }
            }
        },
        "delivery.GetStatisticRequest": {
            "type": "object",
            "required": [
                "from_date",
                "to_date"
            ],
            "properties": {
                "from_date": {
                    "type": "string"
                },
                "to_date": {
                    "type": "string"
                }
            }
        },
        "delivery.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 6
                }
            }
        },
        "delivery.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "delivery.SendCodeRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "delivery.UpdateUserSettingRequest": {
            "type": "object",
            "required": [
                "alarm_sound",
                "repeat_alarm",
                "username"
            ],
            "properties": {
                "alarm_sound": {
                    "type": "string"
                },
                "repeat_alarm": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 1
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "delivery.VerificationRequest": {
            "type": "object",
            "required": [
                "code",
                "email"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "response.response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "sql.NullInt64": {
            "type": "object",
            "properties": {
                "int64": {
                    "type": "integer"
                },
                "valid": {
                    "description": "Valid is true if Int64 is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "18.140.71.34",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Pomodoro API",
	Description:      "Pomodoro Application Api Server. This app helps people study and work at better productivity.\nBackend Language: Golang\nDatabase: PostgreSQL, Redis\nFramework: Gin, sqlc, uber-fx (for dependency injection), uber-zap (for logging files)\nDeployment: AWS EC2, Nginx",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
