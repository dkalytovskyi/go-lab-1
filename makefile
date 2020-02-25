default: out/example

clean:
	rm -rf out

test:
	go vet && go test

out/example: implementation.go main.go
	mkdir out
	myVersion=$(git describe)
	echo "package main
	var buildVersion = \"${myVersion}\"" > buildVersion.go
	go build -o out/example.exe .