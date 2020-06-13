TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

start_couchdb:
	docker run -d -p 5984:5984 --rm -e COUCHDB_USER=admin -e COUCHDB_PASSWORD=admin --name couchdb couchdb:2.3.1
	sleep 2
	curl -X PUT http://admin:admin@localhost:5984/_users
	curl -X PUT http://admin:admin@localhost:5984/_replicator


testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

integration: start_couchdb fmtcheck
	COUCHDB_ENDPOINT=http://localhost:5984 COUCHDB_USERNAME=admin COUCHDB_PASSWORD=admin TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
	-docker rm -f couchdb

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./aws"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile

