# Auth Example API

Welcome to the Auth Example API, a sample server designed to illustrate user authentication and management. This API is built with Golang and utilizes JWT for securing endpoints.

## Overview

The Auth Example API provides a straightforward authentication system, including user signup, authentication, and user management functionalities. It's designed to be a starting point for building secure applications with user authentication.

### API Version

1.0

### License

This project is licensed under the Apache 2.0 License - see the [LICENSE](http://www.apache.org/licenses/LICENSE-2.0.html) file for details.

### Contact Information

- **Name:** Javier Salvador
- **Email:** javier@jsrf.com.br
- **Website:** [www.jsrf.com.br](http://www.jsrf.com.br)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

```bash
make dependencies
```


### Tests

```bash
make test
```

### Installing

1. Clone the repo
2. Install dependencies
3. Start the server

```bash
make
```

The server should start and be accessible at `localhost:8080`.

## API Endpoints

All API requests are made to the base URL: `/api/v1`. Below is a summary of available endpoints.

### Authentication

- **POST `/auth`**: Authenticate with email and password to receive a token.
- **POST `/signup`**: Sign up a new user with email, name, and password.

### User Management (JWT Security)

- **GET `/users/`**: List all users.
- **GET `/users/:userId`**: Get details of a specific user by userID.

## External Documentation

For more detailed API documentation, visit our [Swagger UI](https://localhost:8080/swagger/index.html).

## Acknowledgments

- GORM for ORM support
- Gorilla Mux for routing
- Swaggo for Swagger documentation
