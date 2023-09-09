all: out/client out/server


client: out/client
	@./$< localhost 8080

server: out/server
	@./$< 8080

out/%: src/%/main.go
	go build -o $@ $<

clean:
	rm -rf out

.PHONY: client server