install-wasmtime:
	curl https://wasmtime.dev/install.sh -sSf | bash

dev:
	t1 generate --watch --proxy="http://localhost:3000" --cmd="make up"

up:
	go build .
	go run .

genup:
	t1 generate
	@make up

build-wasm:
	t1 generate
	GOOS=js GOARCH=wasm go build -o cmd/main.wasm .
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
	

