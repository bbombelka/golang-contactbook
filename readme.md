# 📇 Golang Contact Book API
    
[![Go Version](https://img.shields.io/badge/Go-1.24%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-008ECF?style=flat&logo=gin)](https://gin-gonic.com/) 
[![MongoDB](https://img.shields.io/badge/MongoDB-Driver_v2-47A248?style=flat&logo=mongodb)](https://www.mongodb.com/)[![Swagger](https://img.shields.io/badge/Swagger-OpenAPI%202.0-85EA2D?style=flat&logo=swagger)](http://localhost:8080/swagger/index.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
    
A lightweight, performant RESTful CRUD API service for managing contacts and users, built with **Go** and the **Gin** web framework, backed by **MongoDB** (official Go Driver v2), and documented with interactive **Swagger / OpenAPI**.


 ## ✨ Features

 - **⚡ Blazing Fast Routing:** Powered by [Gin](https://github.com/gin-gonic/gin) for minimal overhead and high-speed HTTP processing.
 - **🍃 MongoDB Driver v2:** Uses the modern official `go.mongodb.org/mongo-driver/v2` with connection pooling and ping verification.
- **📖 Interactive API Docs:** Pre-configured Swagger UI available out of the box via `swaggo/gin-swagger`.
- **📄 Pagination:** Built-in `limit` and `skip` query parameters on contact list endpoints.
- **🛡️ Request Validation:** Strict body checks for user creation and partial-field updates on `PATCH` requests.
 - **🆔 UUID Generation:** Automatic unique identifier generation via `google/uuid`.
 - **📁 Modular Architecture:** Clean separation of concerns across handlers, repository operations, data models, and utilities.
 ---
