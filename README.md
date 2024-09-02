
Go-Bookstore API
================

This is a RESTful API for managing an online bookstore, built with Go using the Gorilla Mux router and MySQL as the database. The API supports operations for managing books, categories, user authentication, and profiles.

Features
--------

* CRUD operations for Books and Categories
* User authentication with JWT-based Bearer tokens
* Secure access to protected routes
* Swagger documentation for easy API testing and exploration

Prerequisites
-------------

Before you begin, ensure you have the following installed:

* [Go](https://golang.org/dl/) (1.18+)
* [MySQL](https://www.mysql.com/downloads/)
* [Git](https://git-scm.com/downloads)

Getting Started
---------------

1. **Clone the repository:**

   `git clone https://github.com/your-username/go-bookstore.git cd go-bookstore`
2. **Set up the environment variables:**

   Create a `.env` file in the root directory and add the following environment variables:

   `APP_ENV=development CONFIG_PATH=./config.json ERROR_PATH=./errorcode.json`
3. **Set up the MySQL database:**

   * Create a MySQL database.
   * Update the `config.json` file with your database credentials.
4. **Install dependencies:**

   `go mod tidy`
5. **Run the database migrations:**

   Ensure that your database is up-to-date with the necessary tables.
6. **Run the application:**

   `go run main.go`

   The API will be accessible at `http://localhost:8080`.
7. **Access Swagger UI:**

   Visit `http://localhost:8080/swagger/index.html` to view the API documentation and test endpoints.

API Endpoints
-------------

* `POST /api/v1/auth/signup` - Sign up a new user
* `POST /api/v1/auth/signin` - Log in a user and receive a Bearer token
* `GET /api/v1/books` - Get a list of all books
* `POST /api/v1/books` - Add a new book (protected)
* `GET /api/v1/books/{bookId}` - Get details of a specific book
* `PUT /api/v1/books/{bookId}` - Update a book (protected)
* `DELETE /api/v1/books/{bookId}` - Delete a book (protected)

Acknowledgements
----------------

* [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher for Go
* [GORM](https://gorm.io/) - The fantastic ORM library for Golang
* [Swagger](https://swagger.io/) - API documentation tool
