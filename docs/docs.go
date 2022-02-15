// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "saberazizi66@yahoo.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hello"
                ],
                "summary": "hello",
                "parameters": [
                    {
                        "type": "string",
                        "description": "firstName param",
                        "name": "firstName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "lastName param",
                        "name": "lastName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.HelloDto"
                        }
                    }
                }
            }
        },
        "/os": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "os"
                ],
                "summary": "os",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/person/add": {
            "post": {
                "description": "post the status of server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person api"
                ],
                "summary": "add person",
                "parameters": [
                    {
                        "description": "person body",
                        "name": "personDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Person"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AddPersonsResponseDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDto"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDto"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDto"
                        }
                    },
                    "504": {
                        "description": "Gateway Timeout",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponseDto"
                        }
                    }
                }
            }
        },
        "/person/findAll": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person api"
                ],
                "summary": "find All person",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.FindAllPersonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddPersonsResponseDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorResponseDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "dto.FindAllPersonResponse": {
            "type": "object",
            "properties": {
                "persons": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Person"
                    }
                }
            }
        },
        "dto.HelloDto": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.Person": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "nationalCode": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0.0-1400/11/26",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "saber golang gin",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
