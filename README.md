<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="120" alt="Go Logo" /></a>
</p>

<p align="center"><strong>A lightweight and efficient Golang project documenting the journey through core language fundamentals.</strong></p>

<p align="center">
<a href="https://go.dev/" target="_blank"><img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Go Language" /></a>
<a href="https://github.com/" target="_blank"><img src="https://img.shields.io/github/license/yourusername/your-repo.svg" alt="License" /></a>
<a href="https://github.com/" target="_blank"><img src="https://img.shields.io/github/stars/yourusername/your-repo?style=social" alt="GitHub stars" /></a>
<a href="https://twitter.com/golang" target="_blank"><img src="https://img.shields.io/twitter/follow/golang.svg?style=social&label=Follow" alt="Follow Go on Twitter"></a>
</p>

# AI PDF Chat

A modular Go application that allows users to upload and process PDF files, track processing status, and authenticate using JWT. Built with Go Fiber, PostgreSQL, and a clean layered architecture (MVC + hex-like structure).


## Structure
```
/cmd -> Application entry point
/internal -> Business logic (services, repositories)
/pkg -> Reusable packages (DTO responses, utilities)
/api -> Route handlers and request models
/config -> App configuration (DB, env, JWT setup)
/db -> Database models, migrations
/middleware -> Custom middleware (auth via JWT)

```

## Authentication
```
- JWT-based authentication
- Login returns token
- Token is required for all file actions
- `middleware.ExtractUserID()` is used to identify the current user from token

```

## Configure Environment Variables

Create a .env file and configure: (follow env.example strongly)


## Installation 
(follow go.mod strongly)


## Run Project
```
# run project
make run

# DB migration
make migration
```

## API Endpoints
```
* /api/login

* /api/file/list
* /api/file/upload
* /api/file/processed

* /api/chat/list
* /api/chat/ask-question
```