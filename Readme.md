# XM Broker Service

This project is a microservice built with Golang that handles company operations such as creating, updating, retrieving, and deleting companies. The service uses Kafka for event streaming and JWT for authentication.

## Requirements

- Golang (version 1.16 or higher)
- Docker
- Docker Compose
- PostgreSQL
- Kafka (optional for bonus)

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/xm-broker-service.git
cd xm-broker-service
```

# Setup environment

Ensure you have Docker and Docker Compose installed on your machine. The service requires a config.yaml file to configure database, JWT, and other settings. You can either:

- Provide the path to config.yaml via the CONFIG_PATH environment variable, or
- Pass it as a command-line argument to the service. [./xm server -c <config.yaml>]

### Example Config file

```yaml
db_config:
  host: "db"
  password: "postgres"
  user: "postgres"
  db_name: "xm"
  port: "5432"
  slow_query_threshold: 10

jwt_secret_key: "a3dcb4d229de6fde0db5686dee47145d2d4f3f4b8e1a7f5c5b22d6a9f6142f37"
```

## Run Docker Compose

```bash
sudo docker-compose up -d
```

This command will:

- Start PostgreSQL.
  Run the Golang service (xm) which handles the company-related operations. Server is exposed in port 8090

## API Endpoints

### 1. Login and Obtain JWT Token

**Request**:

```bash
curl --location 'localhost:8090/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "shudip@gmail.com"
}'
```

- Method: POST
- Endpoint: /v1/login

Payload:

```json
{
  "email": "user@example.com"
}
```

**Response**:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoiMjAyNC0xMC0yMFQwMzo0NjozMS44NTE5NzkzNzZaIiwiZW1haWwiOiJzaHVkaXBAZ21haWwuY29tIiwicGVybWlzc2lvbnMiOnsiY29tbW9uIjp0cnVlLCJjcmVhdGVDb21wYW55Ijp0cnVlLCJkZWxldGVDb21wYW55Ijp0cnVlLCJ1cGRhdGVDb21wYW55Ijp0cnVlfX0.kkv5PhdC33Ow9Z0pZkcql_0C0zm9rys-Ypgck0tcd2o"
}
```

The response will contain a JWT token, which should be used for authorized requests.

```
Note: Jwt Token contains a map with permissions which is handled by the middleware for checking permission
```

Example payload of the parsed jwt token

Payload data

```json
{
  "created_at": "2024-10-20T03:46:31.851979376Z",
  "email": "shudip@gmail.com",
  "permissions": {
    "common": true,
    "createCompany": true,
    "deleteCompany": true,
    "updateCompany": true
  }
}
```

### 2. Get Company by ID

```bash
curl --location 'localhost:8090/v1/companies/599c4b1f-8901-4d3c-bf87-795bc82b6f64' \
--header 'Authorization: <JWT Token>'
```

- Method: GET
- Endpoint: /v1/companies/{id}

Returns a company by id

### 3. Create a Company

```bash
curl --location 'localhost:8090/v1/companies' \
--header 'Authorization: Bearer <JWT Token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "599c4b1f-8901-4d3c-bf87-795bc82b6f64",
    "name": "TechCorp",
    "description": "A leading tech company specializing in software development and cloud solutions.",
    "num_employees": 4,
    "registered": true,
    "type": "Corporations"
}'
```

- Method: POST
- Endpoint: /v1/companies

Creates a company. It also validates user request valid company type name length num_employees etc.

### 4. Update a Company

```bash
curl --location --request PUT 'localhost:8090/v1/companies/eff3cf3d-9959-40b8-9659-f56bc84d60d5' \
--header 'Authorization: <JWT Token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "599c4b1f-8901-4d3c-bf87-795bc82b6f64",
    "name": "TechCorp",
    "description": "A leading tech company",
    "num_employees": 4,
    "registered": true,
    "type": "Corporations"
}'
```

- Method: PUT
- Endpoint: /v1/companies/{id}

Updates the information of an company

### 5. Delete a Company

```bash
curl --location --request DELETE 'localhost:8090/v1/companies/e7b42c20-cb00-43f4-8192-2ce7413e4d65' \
--header 'Authorization: <JWT Token>'
```

- Method: DELETE
- Endpoint: /v1/companies/{id}

Deletes a company by id
