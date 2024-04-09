.PHONY=default run build test docs clean run-with-docs

# Variables
APP_NAME=Gopportunies

#Tasks
deafult: run-with-docs
run-with-docs:
	@swag init
	@go run main.go

run: 
	@go run main.go

build: 
	@go build -o $(APP_NAME) main.go 

test: 
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs
