default: out/example

clean:
	rm -rf out

test:
	go vet && go test

out/example: implementation.go main.go
	mkdir out
	printf "package main\nvar buildVersion = \"$(git describe)\"" > buildVersion.go
	go build -o out/example.exe .