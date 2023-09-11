prepare:
	go install github.com/hajimehoshi/wasmserve@latest
	
run-web:
	wasmserve

open-browser:
	http://localhost:8080/