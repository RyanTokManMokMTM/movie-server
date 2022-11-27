# FYP Project - Rebuilt by Go-ZERO

### Introduction
This is my backend project using Go-zero framework, and this also is my **Final Year Project**. I'm trying to rebuild this project using Golang instead **Vapor Swift 3**.

### Tech In Use
* Go-Zero Framework
* MySQL
* GORM
* Docker / Docker-compose
* SwaggerAPI
* Environment(.env)
* Redis *(May use as DB cache in Future)*

### How to run
clone the project 
```
git clone https://github.com/RyanTokManMokMTM/movie-server.git
```

run the server 
```
make run
```
or
```
go run movieservice.go 
```

If you want to run in docker
```
docker-compose -f docker-compose_env.yaml
```

### Working API
* User API
* User Post API
* Movie API
* Genre API
* Movie Likes API
* Post Likes API
* Movie List API
* Post Comment API
* Friend API
* Instant message System(**In progress**)