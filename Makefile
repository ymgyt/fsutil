.PHONY:
test:
	go test 

.PHONY:
dependencies:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure

