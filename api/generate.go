//go:generate swagger generate server --target=../internal/web --spec=./swagger.yml --exclude-main
//go:generate ffjson ../internal/web/models/event.go
//go:generate protoc --go_out=plugins=grpc:../internal/grpc calendar.proto
package api

