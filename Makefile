all: out/client out/server


client: out/client
	@./$<

server: out/server
	@./$<

out/%: src/%/main.go
	go build -o $@ $<

clean:
	rm -rf out

.PHONY: client server