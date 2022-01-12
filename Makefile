.PHONY: cli gen clean build run

APP_NAME = entexample
BUILD_DIR = $(PWD)/bin

cli:
	go get -d entgo.io/ent/cmd/ent

gen:
	go generate ./ent 

clean:
	rm -rf $(BUILD_DIR)

build: 
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: 
	$(BUILD_DIR)/$(APP_NAME)