//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types,gin-server,spec -o petstore.gen.go ../../petstore-expanded.yaml

package api
