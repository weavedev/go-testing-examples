.PHONY: all

GOBIN = $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint
GOTEST = $(GOBIN)/gotest

$(GOTEST):
	GO111MODULE=off go get -u github.com/rakyll/gotest
$(GOLANGCI_LINT):
	(cd ~; GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.21.0)

lint: | $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --deadline=30m --exclude-use-default=false -v

test: | $(GOTEST) 
	@mkdir -p reports
	LOGFORMAT=ASCII gotest -covermode=count -p=10 -coverprofile reports/codecoverage_all.cov `go list ./...`
	@go tool cover -func=reports/codecoverage_all.cov > reports/functioncoverage.out
	@go tool cover -html=reports/codecoverage_all.cov -o reports/coverage.html
	@echo "View report at $(PWD)/reports/coverage.html"
	@tail -n 1 reports/functioncoverage.out

funcCoverage:
	@cat reports/functioncoverage.out

coverageReport:
	@open $(PWD)/reports/coverage.html

fibTest:
	go test ./cmd/fibonacci-test-example -v -cover

fibBench:
	go test ./cmd/fibonacci-test-example -bench=.

parallelTest: 
	go test ./cmd/parallel-test-example  -v

# Forcing parallel suites to run synchronously
notParallelTest:
	go test ./cmd/parallel-test-example -v -parallel=1

signinTest:
	go test ./cmd/api-test-example -run TestAPITestSuite/TestSignInAPICallEmailError -v

signinAsyncTest:
	go test ./cmd/api-test-example -run TestAPITestSuite/TestSignInAPICallEmailError -v 
