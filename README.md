# airac
Auth Identify Resource and Access Controller

## API Spec

// TODO

## Configurations

### Windows

Get clientID and Secret Key each applications.

* [Google]()
* [Twitter]()
  * Permissions > Access permission > check "Read-only"
  * Permissions > Additional permissions > check on "Request email address from users"
* [Facebook]()
* [GitHub](https://github.com/settings/apps/new)
  * User authorization callback URL > "http://localhost:8000/github/callback"
  * User permissions > check on "User permissions" Access Read-only 
  * Where can this GitHub App be installed? > Any account


Set environment variables.

```sh
# Google Client ID
set GOOGLE_CLIENT_ID=<Your Client ID>
set GOOGLE_CLIENT_SECRET=<Your Client Secret>

# Twitter Consumer API keys
set TWITTER_CONSUMER_KEY=<Your Consumer Key>
set TWITTER_CONSUMER_SECRET=<Your Consumer Secret>

# Facebook
set FACEBOOK_CLIENT_ID=<Your Client ID>
set FACEBOOK_CLIENT_SECRET=<Your Client Secret>

# GitHub


# Optional(default 8000) 
set PORT=8000
```

### Mac/Linux


## Developer Guide

### Required

* go 1.12+

### Setup
       
1. git clone http://github.com/laqiiz/airac
2. cd airac
3. go mod download
4. go run main.go
5. curl http://localhost:8080/health
  * confirming return `ok`

### Install git pre-commit hook script before developing.

```bash
# Windows
cd airac
copy /Y .\.githooks\*.* .\.git\hooks

# Mac/Linux
cd airac
cp .githooks/* .git/hooks
chmod +x .git/hooks/pre-commit
```

### Build

Binary for Linux
 
```sh
# on Linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o main main.go

# on Windows
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags "-s -w" -a -installsuffix cgo -o main main.go
```

Docker
```sh
docker build -t airac .
docker run airac
```

## License

This project is licensed under the Apache License 2.0 License - see the [LICENSE](LICENSE) file for details
