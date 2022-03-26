[![Total alerts](https://img.shields.io/lgtm/alerts/g/TheBeachMaster/golang-ent.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/TheBeachMaster/golang-ent/alerts/)

# Golang ent(ity) framework example 

> This is a poor attempt to build a bootleg music streaming application with Golang using coding best practices. This code is for learning purposes only, please obtain a copy of the license before using part(s) of the code in production.

## Structure 

> See `README.md` in each folder for documentation

```
|-config  # All config parsing and loading happens here
  -----/config-local.yaml
  -----/*.go
|-containers # Docker-compose & K8s setup
|-ent
  ---/schema/*.go # Schema definitions
  ---/*.go
|-internal # Actual app logic
|-pkg/*.go # Utilities shared across the app
|-server/*.go # Bootstrapping the server
|-/.main.go # Application entry point
|-/.Makefile # Small scripts to help you build the app
```

## Starting the app (locally)

### Start containers 
> You might need to create a volume that Postgres needs to attach to first;execute the command `docker volume create --name entexample-volume -d local`. **PLEASE DO THIS ONLY ONCE**.

```bash 
cd ./containers && docker-compose up -d 
```

### Get Ent codegen tool (only do this once)

```bash
make cli
```

### Generate  Entities

```bash
make gen
```

### Build the app 

```bash
make build
```

### Run the app 

```bash
make run
```

## Making HTTP/REST requests

Assuming the app is running on `localhost` try adding a song by making the following `curl` request

```bash
curl --request POST \
  --url http://localhost:42069/api/v1/song/add \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "test",
	"fileUrl": "http://file.com/123.mp3"
}'
```