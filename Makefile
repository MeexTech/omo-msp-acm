.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/menu.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/role.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/user.proto
