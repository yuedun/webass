#!/bin/bash

GOOS=js GOARCH=wasm go build -o lib.wasm main.go