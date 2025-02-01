# indico

## Installation

To run this project need:

1. [Taskfile](https://taskfile.dev/installation/)
2. [Go](https://go.dev/doc/install) 1.23
3. [Postman](https://www.postman.com/downloads/)
4. [Docker](https://docs.docker.com/engine/install/)

run this command on terminal it will download dependencies

```bash
task setup
```

```bash
go mod tidy
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`GOOSE_DRIVER`

`GOOSE_DBSTRING`

`GOOSE_MIGRATION_DIR`

`PORT`

`JWT_SECRET`

## Run Locally

Run with docker

```bash
docker-compose up -d
```

Run without docker

```bash
task run
```

## Documentation

import using postman this json

```json
{
  "info": {
    "_postman_id": "402a59b6-1972-4b7d-8eee-055bf62a9d52",
    "name": "Indico",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    "_exporter_id": "14623263"
  },
  "item": [
    {
      "name": "User",
      "item": [
        {
          "name": "Register User",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"email\":\"staff@mail.com\",\r\n    \"password\":\"test1234\",\r\n    \"full_name\":\"staff\",\r\n    \"role_id\":\"9af10c35-dc8b-4bd9-bf05-af5b7dfe083d\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/register",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "register"]
            }
          },
          "response": []
        },
        {
          "name": "Login User",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"email\":\"dadang@mail.com\",\r\n    \"password\":\"test1234\"\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/login",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "login"]
            }
          },
          "response": []
        },
        {
          "name": "Get Users Info",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODc5NjksInN1YiI6IjAxOTRjMGNlLWIxZDQtNzc0Yi1hMTdkLWFlZGM0MmMzZmQxYiJ9.gOvsrdX5JeyuMobNeAapdiHcKYju7dTGj_L7ntfnA_I",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/users/me",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "users", "me"]
            }
          },
          "response": []
        },
        {
          "name": "Get All Admin",
          "request": {
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/users",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "users"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "Role",
      "item": [
        {
          "name": "Add Roles",
          "request": {
            "method": "POST",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/users/register",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "users", "register"]
            }
          },
          "response": []
        },
        {
          "name": "List Roles",
          "request": {
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/roles",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "roles"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "Product",
      "item": [
        {
          "name": "Add Product",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"name\":\"Pencil\",\r\n    \"quantity\":200\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/products",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "products"]
            }
          },
          "response": []
        },
        {
          "name": "Get List Product",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODgwMDYsInN1YiI6IjAxOTRjMGM3LWQ3YjEtN2U1My04MjVmLTYyZDdlNWZmNjU1ZiJ9.bnyteFhWSBbjA7trz7whSSFq4PjBRZ_ZY99UtfjLa1s",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/products",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "products"]
            }
          },
          "response": []
        },
        {
          "name": "Get Product",
          "request": {
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/products/:id",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "products", ":id"],
              "variable": [
                {
                  "key": "id",
                  "value": ""
                }
              ]
            }
          },
          "response": []
        },
        {
          "name": "Update Product",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODgwMDYsInN1YiI6IjAxOTRjMGM3LWQ3YjEtN2U1My04MjVmLTYyZDdlNWZmNjU1ZiJ9.bnyteFhWSBbjA7trz7whSSFq4PjBRZ_ZY99UtfjLa1s",
                  "type": "string"
                }
              ]
            },
            "method": "PUT",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"name\": \"Pulpen\",\r\n    \"sku\": \"\",\r\n    \"quantity\": 200\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/products/:id",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "products", ":id"],
              "variable": [
                {
                  "key": "id",
                  "value": "0194c11c-8461-743b-852e-d170687b77a9"
                }
              ]
            }
          },
          "response": []
        },
        {
          "name": "Delete Product",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODgwMDYsInN1YiI6IjAxOTRjMGM3LWQ3YjEtN2U1My04MjVmLTYyZDdlNWZmNjU1ZiJ9.bnyteFhWSBbjA7trz7whSSFq4PjBRZ_ZY99UtfjLa1s",
                  "type": "string"
                }
              ]
            },
            "method": "DELETE",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/products/:id",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "products", ":id"],
              "variable": [
                {
                  "key": "id",
                  "value": "0194c0d2-2454-7530-bb81-af8b898e0e98"
                }
              ]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "Warehouse",
      "item": [
        {
          "name": "Create Warehouse",
          "request": {
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"name\":\"gudang 1\",\r\n    \"capacity\":5000\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/locations",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "locations"]
            }
          },
          "response": []
        },
        {
          "name": "List Warehouse",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODgwMDYsInN1YiI6IjAxOTRjMGM3LWQ3YjEtN2U1My04MjVmLTYyZDdlNWZmNjU1ZiJ9.bnyteFhWSBbjA7trz7whSSFq4PjBRZ_ZY99UtfjLa1s",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/locations",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "locations"]
            }
          },
          "response": []
        }
      ]
    },
    {
      "name": "Order",
      "item": [
        {
          "name": "Create Order Receive",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg0ODgwMDYsInN1YiI6IjAxOTRjMGM3LWQ3YjEtN2U1My04MjVmLTYyZDdlNWZmNjU1ZiJ9.bnyteFhWSBbjA7trz7whSSFq4PjBRZ_ZY99UtfjLa1s",
                  "type": "string"
                }
              ]
            },
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"product_id\": \"0194c11c-8461-743b-852e-d170687b77a9\",\r\n    \"order_status\": 1,\r\n    \"quantity\": 50\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/orders/receive",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "orders", "receive"]
            }
          },
          "response": []
        },
        {
          "name": "Create Order Ship",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg1MDY5NzAsInN1YiI6IjAxOTRjMGNlLWIxZDQtNzc0Yi1hMTdkLWFlZGM0MmMzZmQxYiJ9.UtOnVvRrMQk0O_qlIuSiMpNGXZ-OJKYy99Vhp_SL-CA",
                  "type": "string"
                }
              ]
            },
            "method": "POST",
            "header": [],
            "body": {
              "mode": "raw",
              "raw": "{\r\n    \"product_id\": \"0194c11c-8461-743b-852e-d170687b77a9\",\r\n    \"order_status\": 1,\r\n    \"quantity\": 20\r\n}",
              "options": {
                "raw": {
                  "language": "json"
                }
              }
            },
            "url": {
              "raw": "http://localhost:7575/api/v1/orders/ship",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "orders", "ship"]
            }
          },
          "response": []
        },
        {
          "name": "List Order",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg1MDY5NzAsInN1YiI6IjAxOTRjMGNlLWIxZDQtNzc0Yi1hMTdkLWFlZGM0MmMzZmQxYiJ9.UtOnVvRrMQk0O_qlIuSiMpNGXZ-OJKYy99Vhp_SL-CA",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/orders",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "orders"]
            }
          },
          "response": []
        },
        {
          "name": "Get Order",
          "request": {
            "auth": {
              "type": "bearer",
              "bearer": [
                {
                  "key": "token",
                  "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg1MDY5NzAsInN1YiI6IjAxOTRjMGNlLWIxZDQtNzc0Yi1hMTdkLWFlZGM0MmMzZmQxYiJ9.UtOnVvRrMQk0O_qlIuSiMpNGXZ-OJKYy99Vhp_SL-CA",
                  "type": "string"
                }
              ]
            },
            "method": "GET",
            "header": [],
            "url": {
              "raw": "http://localhost:7575/api/v1/orders/:id",
              "protocol": "http",
              "host": ["localhost"],
              "port": "7575",
              "path": ["api", "v1", "orders", ":id"],
              "variable": [
                {
                  "key": "id",
                  "value": "0194c1f0-e8c9-70a4-a5ef-9bf46c14b7b3"
                }
              ]
            }
          },
          "response": []
        }
      ]
    }
  ]
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@sutantodadang](https://www.github.com/sutantodadang)
