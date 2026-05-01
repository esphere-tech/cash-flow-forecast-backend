# Cash Flow Forecast Backend

A robust and scalable REST API built with Go and Gin for managing personal cash flow forecasting. This backend provides user authentication, cash entry management, and financial data analysis capabilities.

## Features

- **User Authentication** — Secure signup, login, and JWT-based session management
- **Cash Entry Management** — Create, read, update, and delete cash entries (inflows/outflows)
- **Bulk Operations** — Import multiple cash entries at once
- **UUID-based IDs** — Secure, collision-resistant identifiers
- **Neon PostgreSQL** — Persistent cloud database storage
- **CORS Support** — Cross-origin request handling
- **Docker Ready** — Containerized deployment with multi-stage builds
- **Structured Logging** — Request/response tracking and error diagnostics

## Tech Stack

- **Language:** Go 1.21+
- **Framework:** Gin Web Framework
- **Database:** PostgreSQL (Neon) with GORM ORM
- **Authentication:** JWT (github.com/golang-jwt/jwt)
- **ID Generation:** UUID (github.com/google/uuid)
- **Environment:** godotenv

## Prerequisites

- Go 1.21 or higher
- A Neon PostgreSQL database URL
- Docker & Docker Compose (optional)

## Installation

### Clone the Repository

```bash
git clone github.com/waltertaya
cd cash-flow-forecast-backend
```

### Install Dependencies

```bash
make deps
```

Or manually:

```bash
go mod download
go mod tidy
```

### Environment Setup

Create a `.env` file in the project root:

```env
JWT_SECRET=your-secret-key-here
PORT=8080
DATABASE_URL=postgresql://<user>:<password>@<host>/<database>?sslmode=require
```

## Running the Project

### Local Development

```bash
make run
```

Or:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

### Using Docker

Build and run with Docker:

```bash
make docker-run
```

Or manually:

```bash
docker build -t cash-flow-forecast:latest .
docker run -p 8080:8080 --env-file .env cash-flow-forecast:latest
```

## API Routes

### Base URL

```
http://localhost:8080/api/v1
```

### Authentication Routes

#### Sign Up

- **Endpoint:** `POST /auth/signup`
- **Description:** Register a new user
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "secure_password"
  }
  ```
- **Response:** `201 Created`
  ```json
  {
    "message": "User created successfully"
  }
  ```
- **Validation:**
  - Email: required, valid email format
  - Password: required, minimum 6 characters

#### Login

- **Endpoint:** `POST /auth/login`
- **Description:** Authenticate and receive auth token
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "secure_password"
  }
  ```
- **Response:** `200 OK`
  ```json
  {
    "message": "Logged in successfully"
  }
  ```
- **Sets:** HTTP-only cookie `auth_token` (valid 24 hours)

#### Logout

- **Endpoint:** `POST /auth/logout`
- **Description:** Clear authentication session
- **Auth:** Required
- **Response:** `200 OK`
  ```json
  {
    "message": "Logged out successfully"
  }
  ```

#### Get Current User

- **Endpoint:** `GET /auth/me`
- **Description:** Retrieve authenticated user info
- **Auth:** Required (JWT token)
- **Response:** `200 OK`
  ```json
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com"
  }
  ```

### Cash Entry Routes

All entry routes require authentication (JWT token in `auth_token` cookie).

#### Get All Entries

- **Endpoint:** `GET /entries`
- **Description:** Retrieve all cash entries for the authenticated user
- **Auth:** Required
- **Response:** `200 OK`
  ```json
  [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "user_id": "660e8400-e29b-41d4-a716-446655440000",
      "type": "inflow",
      "amount": 1500.00,
      "category": "salary",
      "description": "Monthly salary",
      "date": "2024-05-01",
      "created_at": 1704067200
    }
  ]
  ```

#### Create Cash Entry

- **Endpoint:** `POST /entries`
- **Description:** Add a single cash entry
- **Auth:** Required
- **Request Body:**
  ```json
  {
    "type": "inflow",
    "amount": 1500.00,
    "category": "salary",
    "description": "Monthly salary",
    "date": "2024-05-01"
  }
  ```
- **Response:** `201 Created`
- **Validation:**
  - type: required, must be "inflow" or "outflow"
  - amount: required, numeric
  - category: optional
  - description: optional
  - date: required, date format

#### Create Multiple Entries

- **Endpoint:** `POST /entries/bulk`
- **Description:** Import multiple cash entries at once
- **Auth:** Required
- **Request Body:**
  ```json
  [
    {
      "type": "inflow",
      "amount": 1500.00,
      "category": "salary",
      "description": "Monthly salary",
      "date": "2024-05-01"
    },
    {
      "type": "outflow",
      "amount": 200.00,
      "category": "utilities",
      "description": "Electric bill",
      "date": "2024-05-02"
    }
  ]
  ```
- **Response:** `201 Created` (array of created entries)

#### Update Cash Entry

- **Endpoint:** `PUT /entries/:id`
- **Description:** Update a specific cash entry
- **Auth:** Required
- **URL Parameters:**
  - `id`: UUID of the cash entry
- **Request Body:** (same as create)
- **Response:** `200 OK` (updated entry)

#### Delete Cash Entry

- **Endpoint:** `DELETE /entries/:id`
- **Description:** Remove a cash entry
- **Auth:** Required
- **URL Parameters:**
  - `id`: UUID of the cash entry
- **Response:** `200 OK`
  ```json
  {
    "message": "Entry deleted successfully"
  }
  ```

#### Get 13-Week Forecast

- **Endpoint:** `GET /entries/forecast?startingCash=1000`
- **Description:** Generate a 13-week cash flow forecast based on historical entries and starting balance
- **Auth:** Required
- **Query Parameters:**
  - `startingCash`: Initial cash balance (required, numeric, defaults to 0)
- **Response:** `200 OK`
  ```json
  {
    "starting_cash": 1000.00,
    "weeks": [
      {
        "week": 1,
        "opening": 1000.00,
        "inflow": 1500.00,
        "outflow": 500.00,
        "closing": 2000.00,
        "warning": false
      },
      {
        "week": 2,
        "opening": 2000.00,
        "inflow": 0.00,
        "outflow": 1200.00,
        "closing": 800.00,
        "warning": false
      },
      {
        "week": 3,
        "opening": 800.00,
        "inflow": 0.00,
        "outflow": 1500.00,
        "closing": -700.00,
        "warning": true
      }
    ]
  }
  ```
- **Description:**
  - Groups all user entries by week relative to today
  - Week 1 = current week, Week 2 = next week, up to Week 13
  - Calculates running balance with opening and closing for each week
  - `warning` flag indicates weeks with negative closing balance
  - Past dates are grouped into their respective historical weeks

## Project Structure

```
cash-flow-forecast-backend/
├── main.go                  # Application entry point
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── Dockerfile               # Docker build configuration
├── .dockerignore             # Docker ignore rules
├── Makefile                 # Development tasks
├── .env.example             # Environment template
├── .gitignore               # Git ignore rules
├── README.md                # This file
├── internals/
│   ├── api/
│   │   └── routes.go        # API route definitions
│   ├── controllers/
│   │   ├── user.go          # User auth handlers
│   │   ├── cash.go          # Cash entry handlers
│   │   └── forecast.go      # Forecast logic and handlers
│   ├── db/
│   │   └── db.go            # Database connection
│   ├── helpers/
│   │   ├── auth.go          # JWT and auth utilities
│   │   └── password-helper.go # Password hashing
│   ├── middlewares/
│   │   ├── auth.go          # JWT validation middleware
│   │   └── cors.go          # CORS middleware
│   ├── migrate/
│   │   └── migrate.go       # Database migrations
│   └── models/
│       ├── user-model.go    # User schema
│       ├── cash-model.go    # CashEntry schema
│       └── forecast-model.go # Forecast schema
```

## Available Commands

```bash
# Build
make build                  # Compile the application
make clean                  # Remove build artifacts

# Development
make run                    # Run the application
make dev                    # Run with hot reload

# Testing
make test                   # Run all tests
make test-coverage          # Generate coverage report

# Code Quality
make fmt                    # Format code
make lint                   # Run linter
make deps-tidy              # Tidy dependencies

# Docker
make docker-build           # Build Docker image
make docker-run             # Run in Docker
make docker-compose-up      # Start with docker-compose
make docker-compose-down    # Stop docker-compose services
```

## Error Handling

The API returns standard HTTP status codes:

- `200 OK` — Successful request
- `201 Created` — Resource created
- `400 Bad Request` — Invalid input
- `401 Unauthorized` — Missing/invalid authentication
- `404 Not Found` — Resource not found
- `409 Conflict` — Resource already exists (e.g., duplicate email)
- `500 Internal Server Error` — Server error

Error responses include a descriptive message:

```json
{
  "error": "Invalid user ID"
}
```

## Security Considerations

- Passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- Auth cookies are HTTP-only to prevent XSS
- CORS is configured to prevent unauthorized cross-origin requests
- UUIDs are used for collision-resistant identifiers
- Environment variables store sensitive data (JWT_SECRET, DB_URL)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit changes (`git commit -am 'Add your feature'`)
4. Push to branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License — see LICENSE file for details.

## Support

For issues, questions, or suggestions, please open an issue on the repository or contact the development team.
