package cratedb

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate types -o types.gen.go -package cratedb cloud-api-openapi-v1.0.0.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -generate client -o client.gen.go -package cratedb cloud-api-openapi-v1.0.0.json
