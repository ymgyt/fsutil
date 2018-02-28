.PHONY:
test:
	go test 

.PHONY:
dependencies:
	dep version || go -u get github.com/golang/dep/cmd/dep
	dep ensure

