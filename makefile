generate:
		protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
run:
	go run main.go