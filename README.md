# Asterix-GO!
how to run
in cmd/command prompt run syntax 
install dependencies
```bash
go mod download
```

run go
```bash
go run .
```

if go.mod turn to yellow, just run (To remove any dependencies that are no longer needed)
```bash
go mod tidy
```

you can build and run your Go application as usual:
```bash
go build
```

you can build and up with docker, just run:
```bash
docker-compose build
```
and
```bash
docker-compose up
```

how to run with hot reload
in cmd/command prompt run syntax
```bash
go install github.com/cosmtrek/air@latest
```
then
```bash
air -c .air.toml
```
last
```bash
air
```

run golang with hot reload just 
```bash
air
```