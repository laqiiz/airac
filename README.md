# airac
Auth Identify Resource and Access Controller

## API Spec

// TODO

## Required Configurations

```sh
# Optional(default 8000) 
export PORT=8000
```


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
