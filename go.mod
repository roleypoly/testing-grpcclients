module github.com/roleypoly/grpcdemo

go 1.12

require (
	github.com/roleypoly/rpc v0.0.0-20190809193116-ff425f7c11c0
	google.golang.org/grpc v1.22.1
)

replace github.com/roleypoly/rpc => ../rpc
