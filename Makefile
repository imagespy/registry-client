.PHONY: test_e2e
test_e2e:
	godog

.PHONY: test_unit
test_unit:
	go test

.PHONY: test
test: test_unit test_e2e
