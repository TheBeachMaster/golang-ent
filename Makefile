.PHONY: clean build run

APP_NAME = entexample
BUILD_DIR = $(PWD)/bin

clean:
	rm -rf $(BUILD_DIR)

build: 
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: 
	$(BUILD_DIR)/$(APP_NAME)