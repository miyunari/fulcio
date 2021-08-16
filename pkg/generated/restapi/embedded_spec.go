// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Fulcio",
    "version": "1.0.0"
  },
  "host": "fulcio.sigstore.dev",
  "basePath": "/api/v1",
  "paths": {
    "/signingCert": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "create a cert, return content with a location header (with URL to CTL entry)",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/pem-certificate-chain"
        ],
        "operationId": "signingCert",
        "parameters": [
          {
            "description": "Request for signing certificate",
            "name": "CertificateRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CertificateRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Generated Certificate Chain",
            "schema": {
              "type": "string"
            },
            "headers": {
              "SCT": {
                "type": "string",
                "format": "byte",
                "description": "Signed Certificate Timestamp from Entry in CT Log"
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "default": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "CertificateRequest": {
      "type": "object",
      "required": [
        "publicKey",
        "signedEmailAddress"
      ],
      "properties": {
        "publicKey": {
          "type": "object",
          "required": [
            "content"
          ],
          "properties": {
            "algorithm": {
              "type": "string",
              "enum": [
                "ecdsa"
              ]
            },
            "content": {
              "type": "string",
              "format": "byte"
            }
          }
        },
        "signedEmailAddress": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "The content supplied to the server was invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        }
      }
    },
    "InternalServerError": {
      "description": "There was an internal error in the server while processing the request",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        }
      }
    },
    "Unauthorized": {
      "description": "The request could not be authorized",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        },
        "WWW-Authenticate": {
          "type": "string",
          "description": "Information about required authentication to access server"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Fulcio",
    "version": "1.0.0"
  },
  "host": "fulcio.sigstore.dev",
  "basePath": "/api/v1",
  "paths": {
    "/signingCert": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "create a cert, return content with a location header (with URL to CTL entry)",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/pem-certificate-chain"
        ],
        "operationId": "signingCert",
        "parameters": [
          {
            "description": "Request for signing certificate",
            "name": "CertificateRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CertificateRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Generated Certificate Chain",
            "schema": {
              "type": "string"
            },
            "headers": {
              "SCT": {
                "type": "string",
                "format": "byte",
                "description": "Signed Certificate Timestamp from Entry in CT Log"
              }
            }
          },
          "400": {
            "description": "The content supplied to the server was invalid",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "Content-Type": {
                "type": "string"
              }
            }
          },
          "401": {
            "description": "The request could not be authorized",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "Content-Type": {
                "type": "string"
              },
              "WWW-Authenticate": {
                "type": "string",
                "description": "Information about required authentication to access server"
              }
            }
          },
          "default": {
            "description": "There was an internal error in the server while processing the request",
            "schema": {
              "$ref": "#/definitions/Error"
            },
            "headers": {
              "Content-Type": {
                "type": "string"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CertificateRequest": {
      "type": "object",
      "required": [
        "publicKey",
        "signedEmailAddress"
      ],
      "properties": {
        "publicKey": {
          "type": "object",
          "required": [
            "content"
          ],
          "properties": {
            "algorithm": {
              "type": "string",
              "enum": [
                "ecdsa"
              ]
            },
            "content": {
              "type": "string",
              "format": "byte"
            }
          }
        },
        "signedEmailAddress": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "CertificateRequestPublicKey": {
      "type": "object",
      "required": [
        "content"
      ],
      "properties": {
        "algorithm": {
          "type": "string",
          "enum": [
            "ecdsa"
          ]
        },
        "content": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "The content supplied to the server was invalid",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        }
      }
    },
    "InternalServerError": {
      "description": "There was an internal error in the server while processing the request",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        }
      }
    },
    "Unauthorized": {
      "description": "The request could not be authorized",
      "schema": {
        "$ref": "#/definitions/Error"
      },
      "headers": {
        "Content-Type": {
          "type": "string"
        },
        "WWW-Authenticate": {
          "type": "string",
          "description": "Information about required authentication to access server"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
