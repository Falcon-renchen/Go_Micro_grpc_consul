cd Services/protos
protoc --go_out=. --micro_out=. Models.proto
protoc --go_out=. --micro_out=. ProdService.proto
cd.. && cd..