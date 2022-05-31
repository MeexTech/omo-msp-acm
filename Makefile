.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/menu.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/role.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/user.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/cata.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/scene.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/acm/module.proto