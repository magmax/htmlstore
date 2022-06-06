PROTO=$(wildcard api/*.proto)
API=$(patsubst %.proto,pkg/github.com/magmax/htmlstore/%.pb.go,$(PROTO))

all: $(API)

pkg/github.com/magmax/htmlstore/api/%.pb.go: api/%.proto
	mkdir -p pkg
	protoc --go_out=pkg --go-grpc_out=pkg $<

