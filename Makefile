PROTO=$(wildcard api/*.proto)
API=$(patsubst %.proto,%.pb.go,$(PROTO))

all: $(API)

%.pb.go: %.proto
	mkdir -p pkg
	protoc --go_out=. --go-grpc_out=. $<

