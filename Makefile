.PHONY: install server client examples

all: server client examples

server:
	@echo "--> Generating go files"
	protoc -I proto/ --go_out=plugins=grpc:server/proto/ proto/video_streaming.proto
	@echo ""

client:
	@echo "--> Generating Python client files"
	python3 -m grpc_tools.protoc -I prot