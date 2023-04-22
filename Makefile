APP_NAME = shopwebapi
BUILD_DIR = ./build
ROOT_DIR = ./cmd/$(APP_NAME)

clean:
	rm -rf ./build

build: clean
	go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(ROOT_DIR)/main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init -g $(ROOT_DIR)/main.go
