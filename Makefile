.PHONY: check-upgrade-deps
check-upgrade-deps:
	go list -u -m all

.PHONY: test-upgrade-deps
test-upgrade-deps:
	go get -t -u ./...

.PHONY: upgrade-deps
upgrade-deps:
	go get -u ./...

.PHONY: test
test:
	go test all
