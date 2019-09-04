SRCS=$( shell find . -type f -name '*.go' )

ap7900: $(SRCS)
	go build -o $@ .

clean:
	rm -f ap7900

.PHONY: clean
