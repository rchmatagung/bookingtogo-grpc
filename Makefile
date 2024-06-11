generate-protoc-master_hotel:
	protoc -I./grpc/protobuf \
	--go_out=.//grpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=.//grpc/pb --go-grpc_opt=paths=source_relative \
	grpc/protobuf/master_hotel/*.proto