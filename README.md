# webass
go webassembly

> GOOS=js GOARCH=wasm go build -o lib.wasm main.go

> cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .