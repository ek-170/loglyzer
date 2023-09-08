.PHONY: setup
setup:
	bash ./script/setup.sh

.PHONY: serve
serve: air

.PHONY: openapi
openapi:
	docker run --rm -v ${PWD}:/local -v ${PWD}/spec:/spec openapitools/openapi-generator-cli generate -g go-server -i /spec/reference/codehub.yaml -o /local/gen --additional-properties=packageName=openapi,router=chi,sourceFolder=openapi
	golangci-lint run --no-config --enable=goimports --print-issued-lines=false --issues-exit-code=0 --fix --fast -v gen/...
