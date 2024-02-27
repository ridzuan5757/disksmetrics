.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -v -coverprofile=coverage.out

.PHONY: report
report:
	git clone https://github.com/gojp/goreportcard.git \
	&& cd goreportcard \
	&& make install \
	&& go install ./cmd/goreportcard-cli \
	&& goreportcard-cli
