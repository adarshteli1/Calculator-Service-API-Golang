# Calculator Service API

A lightweight RESTful Calculator Service built using **Golang** and **Gin**. This project demonstrates clean backend architecture by separating routing, HTTP handlers, and business logic into independent layers.

---

## Features

* Perform basic arithmetic operations:

  * Addition (`+`)
  * Subtraction (`-`)
  * Multiplication (`*`)
  * Division (`/`)
* JSON-based request and response
* Structured error handling
* Execution time included in responses
* Clean project architecture (Router → Handler → Service)

---

## Tech Stack

* **Language:** Go
* **Framework:** Gin
* **Architecture:** Layered Architecture
* **API Style:** REST

---

## Project Structure

```text
calculator/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── handler/
│   ├── routes/
│   └── service/
│
├── go.mod
├── go.sum
└── README.md
```

---

## API Endpoint

### Calculate

**POST** `/cal`

### Request Body

```json
{
    "operand1": 50,
    "operand2": 5,
    "operation": "+"
}
```

### Successful Response

```json
{
    "answer": 55,
    "operation": "+",
    "executionTime": "430.097µs"
}


<img width="1232" height="566" alt="Screenshot From 2026-07-10 02-28-45" src="https://github.com/user-attachments/assets/1190e1fc-3d62-4351-bb77-ba1d037f776a" />


```

---

## Supported Operations

| Operator | Description    |
| -------- | -------------- |
| `+`      | Addition       |
| `-`      | Subtraction    |
| `*`      | Multiplication |
| `/`      | Division       |

---

## Error Responses

### Division by Zero

```json
{
    "error": "Cannot Divide By Zero"
}
```

### Invalid Operator

```json
{
    "error": "Invalid operator"
}
```

### Invalid JSON

```json
{
    "error": "invalid character 'a' looking for beginning of value"
}
```

---

## Running the Project

Clone the repository:

```bash
git clone <repository-url>
cd calculator
```

Install dependencies:

```bash
go mod tidy
```

Run the server:

```bash
go run ./cmd/api
```

The server starts on:

```
http://localhost:8000
```

---

## Testing

You can test the API using:

* Postman
* Bruno
* Insomnia
* curl

Example:

```bash
curl -X POST http://localhost:8000/cal \
-H "Content-Type: application/json" \
-d '{
    "operand1":10,
    "operand2":20,
    "operation":"+"
}'
```

---

## Architecture

```text
Client
   │
HTTP Request
   │
Router
   │
Handler
   │
Service
   │
Business Logic
   │
Handler
   │
HTTP Response
   │
Client
```

### Responsibilities

#### Router

* Registers API endpoints.
* Maps incoming requests to handlers.

#### Handler

* Parses JSON requests.
* Validates incoming data.
* Calls the service layer.
* Handles errors.
* Returns JSON responses.

#### Service

* Contains business logic.
* Performs arithmetic calculations.
* Returns the result or an error.
* Has no dependency on HTTP or Gin.

---

## Future Improvements

* Scientific calculator operations
* Input validation using struct tags
* Unit tests
* Swagger/OpenAPI documentation
* Request logging middleware
* Docker support
* Configuration management
* API versioning
* Rate limiting
* Authentication

---

## Learning Outcomes

This project helped reinforce:

* REST API development with Gin
* Clean architecture principles
* Separation of concerns
* Request and response DTOs
* JSON binding
* Error handling in Go
* Layered backend design
* Measuring request execution time
* Building maintainable backend services

---

## Author

**Adarsh Teli**

Built as part of a Golang Backend Engineering learning journey focused on writing production-style Go services using clean architecture and engineering best practices.
