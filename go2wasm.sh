#!/bin/bash

GOOS=js GOARCH=wasm go build -o build/lib.wasm main.go