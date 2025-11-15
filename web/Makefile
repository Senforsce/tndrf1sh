install-wasmtime:
	curl https://wasmtime.dev/install.sh -sSf | bash

dev:
	t1 generate --watch --proxy="http://localhost:3000" --cmd="make up"

up:
	orch8rion go build .
	orch8rion go run .

regenup:
	@make -C ../ regen
	t1 generate
	@make up

genup:
	@make -C ../../fh generate
	@make -C ../../owl generate
	@make -C ../../htmx generate

	t1 generate
	@make up

build-wasm:
	t1 generate
	GOOS=js GOARCH=wasm go build -o cmd/main.wasm .
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
	

