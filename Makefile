SRCS=$( shell find . -type f -name '*.go' )

ap7900: $(SRCS) vendor
	go build -o $@ .

clean:
	rm -f ap7900

vendor:
	go mod vendor

.PHONY: clean
