# Go-Station: Reminder



## Getting Started

Before trying the application, make sure you create a .env file in the root of the project.

```
# .env
POSTGRESHOST=your_db_host
POSTGRESPORT=your_db_port
POSTGRESUSER=your_db_user
POSTGRESPASSWORD=your_db_password
POSTGRESDATABASE=your_db_name
SERVERPORT=your_server_port
API_SECRET_KEY = jwt_secret_key
CONFIG_MAIL_HOST = email_server_url 
CONFIG_MAIL_PORT = email_server_port
CONFIG_AUTH_NAME = sender_name
CONFIG_AUTH_EMAIL = sender_email 
CONFIG_AUTH_PASS = sender_pass
```

This application uses SvelteKit as the Front-End Framework in the `web` directory.
To set it up:
```
$ > cd web
$ > npm install
```

## Running Up
```
$ > go run cmd/reminder/main.go
```

***

## Project Name
Reminder

## Description
We do have servers and we are humans with tendency to forget when the server will expire. Thus, this project.

## Support
Put some snacks on R&D table and you will see magic

## Contributing
Use standard GoLang style, make it tidy and we are good to go ;)

## Authors
Research and Development Team in Invosa Systems

## Project Status
Ongoing development
