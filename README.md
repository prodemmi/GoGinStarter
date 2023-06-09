# GoGinStarter

GoGinStarter is a starter project for building web applications using the Go programming language and Gin web framework.

It provides a basic structure and set of features to kickstart your project.

# Features
- [Wire](https://github.com/google/wire]) for Dependency Injection
- Logger
- Cache
- Paginator
- Session and Cookies
- Command Lines using [cobra](https://github.com/spf13/cobra)
- Event/Listeners
- Localization
- Notifications
- API Responses
- Authentication with OTP and JWT
- Scheduling Tasks with [gocron](https://github.com/go-co-op/gocron)
- Configuration management for yaml files
- Database migrations and seeders with GORM

# Supported Drivers
- Database: MySql, PostgresSql
- Cache: File, Redis
- Notification: sms, email

# Built With

- Go - The programming language used
- Gin - The web framework used
- Gorm - The database ORM tool used

# Installation
- Clone the repository:

```git clone https://github.com/prodemmi/GoGinStarter.git```

- Install the dependencies:

```go mod download```
- Make yaml config file:

```cp .yaml.example .yaml```

- Serve the web application:

```go run main.go serve```


# Contributing
Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request.

# License
GoGinStarter is licensed under the MIT License. See LICENSE for more information.
