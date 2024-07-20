# Online Bank Platform Application

## Overview

This repository contains the source code and Docker configuration for the Online Bank Platform application. The application is built using Go and deployed as a Docker container.

## Directory Structure


- `cmd/auth-service/`: Contains the source code for the application.
  - `main.go`: Entry point of the Go application.
- `docker/`:
  - `auth-service.Dockerfile`: Docker configuration file for building the application container.
- `go.mod` and `go.sum`: Go modules files for managing dependencies.
- `README.md`: This file.
- `.gitignore`: Defines files and directories to be ignored by Git.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)
