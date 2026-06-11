GO ?= go

run-all:
	@$(MAKE) -j3 run-command run-query run-worker

run-command:
	$(GO) -C packages/shop-command run .

run-query:
	$(GO) -C packages/shop-query run .

run-worker:
	$(GO) -C packages/shop-worker run .
