generate:
	rm gen -rf;
	cd protos; buf mod update; buf generate;
	rm gen/crud -r; rm gen/grpc -r; rm gen/modles -r; rm gen/main.go; rm gen/kafka_gateway -r; rm gen/kafka_service -r; rm gen/locker -r;
