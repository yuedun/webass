# webass
go webassembly

将main.go编译为lib.wasm
> GOOS=js GOARCH=wasm go build -o lib.wasm main.go

执行一次即可，用于调用和执行wasm文件
> cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

运行：
> ./js2wasm
> go run server.go