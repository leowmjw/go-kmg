run:
	@go run cmd/cli/main.go

gen:
	@go generate ./...

openapi:
	@oapi-codegen -package rootly -generate "types,client,chi-server,spec" ./OASv3/rootly/document.json > ./api/rootly/rootly.gen.go

ogenserver:
	@ogen -package rootly -target ./server/rootly -generate-tests -infer-types -debug.ignoreNotImplemented all -clean -v ./OASv3/rootly/document.json

tools:
	@go install github.com/ogen-go/ogen/cmd/ogen@latest

