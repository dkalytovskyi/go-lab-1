default: out/example

clean:
	rm -rf out

test:
	go vet && go test

out/example: implementation.go main.go
	mkdir out
	git describe > buildVersion.txt
	go build -o out/example.exe .