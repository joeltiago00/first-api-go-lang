##### Go Version: 1.18

## Building Container

Start container:
```
docker compose up -d
```

Turn off container:
```
docker compose down
```

If you need in container as root, just run:
```
docker compose exec -uroot postgres bash
```

<hr>

## Setup Application

Create `.env` file:
```
cp .env.example .env
```

##### PS: Don't forget to put value in empty variables

Install project dependencies:
```
go mod tidy
```

Run local server:
```
go run main.go
```

<hr>

## Running Application Tests

Running all tests at same time:
```
go test -v ./tests/...
```

Running specif file test:
```
go test -v ./tests/path_of_file_test.go
```

Running specif test on file:
```
go test -v ./tests/path_of_file_test.go -run TestFunctionName
```

