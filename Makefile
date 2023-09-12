prepare:
	go install github.com/hajimehoshi/wasmserve@latest
	
run-web:
	wasmserve -http=:8080

open-browser:
	http://localhost:8080/