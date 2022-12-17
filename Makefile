# flatc --binary api/fbs/req.fbs cmd/client/ip.json
# flatc --binary api/fbs/res.fbs cmd/server/op.json
.PHONY: gen server

gen:
	flatc --go --grpc --gen-object-api --go-module-name shiv -o api/ api/fbs/*.fbs

server:
	cd cmd/server && go run *.go

client:
	cd cmd/client && go run main.go

