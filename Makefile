check_quality:
	go test -cover ./... && \
	go vet ./... && \
	go fmt ./... && \
	codespell && \
	gocyclo . && \
	ineffassign ./...

watch_test:
	find . -name "*.go" | entr -c go test -cover ./...
