.PHONY:
test:
	go test 

.PHONY:
dependencies:
	dep version || go get -u get github.com/golang/dep/cmd/dep
	dep ensure

