# Go Server API

This Go server provides an API to manage users, including operations to get, add, update, and delete users, as well as add hours worked. It uses the `gorilla/mux` package for routing and `gorilla/handlers` for enabling CORS.

## Installation Requirements

Before you can run the server, you will need the following installed on your machine:

- [Go 1.18+](https://golang.org/dl/)
- [Git](https://git-scm.com/)
- [curl](https://curl.se/) or any other API testing tool like [Postman](https://www.postman.com/)

### On Mac:

1. Ensure you have [Go](https://golang.org/dl/) installed. You can check if Go is installed by running:
   ```bash
   go version
   ```

2. If you don't have Go installed, you can install it using Homebrew:
   ```bash
   brew install go
   ```

3. Install necessary Go modules for the project:
   ```bash
   go mod tidy
   ```

4. Install the required Go packages:
   ```bash
   go get -u github.com/gorilla/mux
   go get -u github.com/gorilla/handlers
   ```

### On Windows:

1. Ensure you have [Go](https://golang.org/dl/) installed. You can check if Go is installed by running:
   ```bash
   go version
   ```
   If Go is not installed, download and install Go from here - https://go.dev/dl/


### Running the Server:

1. In the project directory, run the following command to start the server:
   ```bash
   go run main.go
   ```

2. The server will start on port 5004 by default. You should see output like:
   ```bash
   Starting server on port 5004
   ```

### API Endpoints:

The following endpoints are available:

- GET /users: Get a list of all users.
- GET /users/{id}: Get a user by ID.
- POST /users: Add a new user.
- PUT /users/{id}: Update a user by ID.
- PATCH /users/{id}/hours: Add hours worked for a user.
- DELETE /users/{id}: Delete a user by ID.
- DELETE /users: Delete all users.

