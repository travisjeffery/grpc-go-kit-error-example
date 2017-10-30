compile:
	protoc *.proto --gogofast_out=plugins=grpc:. --proto_path=. --proto_path=vendor
