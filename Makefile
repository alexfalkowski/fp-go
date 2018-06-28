dependencies:
	go get github.com/alecthomas/gometalinter
	gometalinter --install

	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega/...

	go get github.com/golang/dep/cmd/dep
	dep ensure -v

spec:
	go test -v -tags=spec -race -cover -coverprofile=coverage.out ./...

coverage:
	go tool cover -func=coverage.out

analysis:
	gometalinter --vendor ./... --disable-all --enable=gofmt --enable=golint --enable=vet --enable=errcheck
