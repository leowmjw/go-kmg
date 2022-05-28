run:
	@go run cmd/cli/main.go

gen:
	@go generate ./...

openapi:
	@oapi-codegen -package rootly -generate "types,client,chi-server,spec" ./OASv3/rootly/document.json > ./api/rootly/rootly.gen.go
