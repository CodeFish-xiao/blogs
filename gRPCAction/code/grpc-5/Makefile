API_PROTO_FILES=$(shell find pb -name *.proto)

api:
	protoc --proto_path=./ \
	       --proto_path=./third_party \
	       --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       --grpc-gateway_out . \
           --grpc-gateway_opt logtostderr=true \
           --grpc-gateway_opt paths=source_relative \
           --grpc-gateway_opt generate_unbound_methods=true \
           --openapiv2_out . \
           --openapiv2_opt logtostderr=true \
	       $(API_PROTO_FILES)

