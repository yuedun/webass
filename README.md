# webass
go webassembly

将main.go编译为lib.wasm
> chmod u+x go2wasm.sh && ./go2wasm

或

> GOOS=js GOARCH=wasm go build -o build/lib.wasm main.go

执行一次即可，用于调用和执行wasm文件，已经存在wasm_exec.js文件，可不执行
> cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build

运行：
> ./go2wasm

> go run server.go