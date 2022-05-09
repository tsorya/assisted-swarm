// Code generated by go-swagger; DO NOT EDIT.

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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Assisted swarm",
    "title": "AssistedSwarm",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "api.assisted-swarm.com",
  "basePath": "/api/assisted-swarm",
  "paths": {
    "/agents": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "List all running agents.",
        "operationId": "ListAgents",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/agent_list"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "swarm"
        ],
        "summary": "Create new agent.",
        "operationId": "CreateNewAgent",
        "parameters": [
          {
            "description": "Create new agent for swarm.",
            "name": "new-agent-params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/new-agent-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success.",
            "schema": {
              "$ref": "#/definitions/agent"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/agents/{agent_id}": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Get specific agent.",
        "operationId": "GetAgent",
        "parameters": [
          {
            "type": "integer",
            "name": "agent_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/agent"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "swarm"
        ],
        "summary": "Delete agent.",
        "operationId": "DeleteAgent",
        "parameters": [
          {
            "type": "integer",
            "name": "agent_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Success"
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/exit": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Exit the process.",
        "operationId": "Exit",
        "responses": {
          "204": {
            "description": "Success."
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Health check.",
        "operationId": "Health",
        "responses": {
          "204": {
            "description": "Success."
          }
        }
      }
    }
  },
  "definitions": {
    "agent": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer"
        },
        "status": {
          "$ref": "#/definitions/agent_status"
        },
        "terminated_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "agent_list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/agent"
      }
    },
    "agent_status": {
      "type": "string",
      "enum": [
        "Running",
        "Terminated"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "kind",
        "id",
        "href",
        "code",
        "reason"
      ],
      "properties": {
        "code": {
          "description": "Globally unique code of the error, composed of the unique identifier of the API and the numeric identifier of the error. For example, if the numeric identifier of the error is 93 and the identifier of the API is assisted_install then the code will be ASSISTED-INSTALL-93.",
          "type": "string"
        },
        "href": {
          "description": "Self link.",
          "type": "string"
        },
        "id": {
          "description": "Numeric identifier of the error.",
          "type": "integer",
          "format": "int32",
          "maximum": 504,
          "minimum": 400
        },
        "kind": {
          "description": "Indicates the type of this object. Will always be 'Error'.",
          "type": "string",
          "enum": [
            "Error"
          ]
        },
        "reason": {
          "description": "Human-readable description of the error.",
          "type": "string"
        }
      }
    },
    "new-agent-params": {
      "type": "object",
      "properties": {
        "agent_version": {
          "type": "string"
        },
        "cacert": {
          "type": "string"
        },
        "containers_conf": {
          "type": "string"
        },
        "containers_storage_conf": {
          "type": "string"
        },
        "dry_cluster_hosts_path": {
          "type": "string"
        },
        "dry_fake_reboot_marker_path": {
          "type": "string"
        },
        "dry_forced_host_id": {
          "type": "string",
          "format": "uuid"
        },
        "dry_forced_host_ipv4": {
          "type": "string",
          "pattern": "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)[\\/]([1-9]|[1-2][0-9]|3[0-2]?)$"
        },
        "dry_forced_hostname": {
          "type": "string"
        },
        "dry_forced_mac_address": {
          "type": "string",
          "format": "mac"
        },
        "infra_env_id": {
          "type": "string",
          "format": "uuid"
        },
        "pull_secret": {
          "type": "string"
        },
        "service_url": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Assisted swarm installation",
      "name": "swarm"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Assisted swarm",
    "title": "AssistedSwarm",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "api.assisted-swarm.com",
  "basePath": "/api/assisted-swarm",
  "paths": {
    "/agents": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "List all running agents.",
        "operationId": "ListAgents",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/agent_list"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "swarm"
        ],
        "summary": "Create new agent.",
        "operationId": "CreateNewAgent",
        "parameters": [
          {
            "description": "Create new agent for swarm.",
            "name": "new-agent-params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/new-agent-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success.",
            "schema": {
              "$ref": "#/definitions/agent"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/agents/{agent_id}": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Get specific agent.",
        "operationId": "GetAgent",
        "parameters": [
          {
            "type": "integer",
            "name": "agent_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/agent"
            }
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "swarm"
        ],
        "summary": "Delete agent.",
        "operationId": "DeleteAgent",
        "parameters": [
          {
            "type": "integer",
            "name": "agent_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Success"
          },
          "400": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Forbidden.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Not Found.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Error.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/exit": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Exit the process.",
        "operationId": "Exit",
        "responses": {
          "204": {
            "description": "Success."
          }
        }
      }
    },
    "/health": {
      "get": {
        "tags": [
          "swarm"
        ],
        "summary": "Health check.",
        "operationId": "Health",
        "responses": {
          "204": {
            "description": "Success."
          }
        }
      }
    }
  },
  "definitions": {
    "agent": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer"
        },
        "status": {
          "$ref": "#/definitions/agent_status"
        },
        "terminated_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "agent_list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/agent"
      }
    },
    "agent_status": {
      "type": "string",
      "enum": [
        "Running",
        "Terminated"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "kind",
        "id",
        "href",
        "code",
        "reason"
      ],
      "properties": {
        "code": {
          "description": "Globally unique code of the error, composed of the unique identifier of the API and the numeric identifier of the error. For example, if the numeric identifier of the error is 93 and the identifier of the API is assisted_install then the code will be ASSISTED-INSTALL-93.",
          "type": "string"
        },
        "href": {
          "description": "Self link.",
          "type": "string"
        },
        "id": {
          "description": "Numeric identifier of the error.",
          "type": "integer",
          "format": "int32",
          "maximum": 504,
          "minimum": 400
        },
        "kind": {
          "description": "Indicates the type of this object. Will always be 'Error'.",
          "type": "string",
          "enum": [
            "Error"
          ]
        },
        "reason": {
          "description": "Human-readable description of the error.",
          "type": "string"
        }
      }
    },
    "new-agent-params": {
      "type": "object",
      "properties": {
        "agent_version": {
          "type": "string"
        },
        "cacert": {
          "type": "string"
        },
        "containers_conf": {
          "type": "string"
        },
        "containers_storage_conf": {
          "type": "string"
        },
        "dry_cluster_hosts_path": {
          "type": "string"
        },
        "dry_fake_reboot_marker_path": {
          "type": "string"
        },
        "dry_forced_host_id": {
          "type": "string",
          "format": "uuid"
        },
        "dry_forced_host_ipv4": {
          "type": "string",
          "pattern": "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)[\\/]([1-9]|[1-2][0-9]|3[0-2]?)$"
        },
        "dry_forced_hostname": {
          "type": "string"
        },
        "dry_forced_mac_address": {
          "type": "string",
          "format": "mac"
        },
        "infra_env_id": {
          "type": "string",
          "format": "uuid"
        },
        "pull_secret": {
          "type": "string"
        },
        "service_url": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Assisted swarm installation",
      "name": "swarm"
    }
  ]
}`))
}