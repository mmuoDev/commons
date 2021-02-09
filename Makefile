OUTPUT = main # will be archived
SERVICE_NAME = commons


build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go
	
run: 
	go run main.go