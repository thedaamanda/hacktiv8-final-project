# Hacktiv8 Golang Final Project - Scalable Web Service with Golang
This is Final Project for Scalable Web Service With Golang by Hacktiv8.
This project is created using Golang Programming Language with Echo Framework, GORM for object relational mapping for PostgreSQL database.

## Usage
* Define the application configuration first using `.env` file (see `.env.example`). You can copy the `.env.example` file and fill in the blanks..
* Run with `go run main.go`
* You're ready to go!

## Endpoints
|  Method | URL | Description |
| ------------ | ------------ | ------------ |
| POST | /users/register | Register a new `user` |
| POST | /users/login | Login `user` |
| PUT | /users/:id | Update an `user` |
| DELETE | /users | Delete an `user` |
| POST | /photos | Create a new `photo` |
| GET | /photos | Get all `photo` |
| PUT | /photos/:id | Update a `photo` |
| DELETE | /photos/:id | Delete a `photo` |
| POST | /comments | Create a new `comment` |
| GET | /comments | Get all `comment` |
| PUT | /comments/:id | Update a `comment` |
| DELETE | /comments/:id | Delete a `comment` |
| POST | /socialmedias | Create a new `social media` |
| GET | /socialmedias | Get all `social media` |
| PUT | /socialmedias/:id | Update a `social media` |
| DELETE | /socialmedias/:id | Delete a `social media` |
