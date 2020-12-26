.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/scene.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/group.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/organization/common.proto
