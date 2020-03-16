VERSION?="0.0.1"
TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go')
WEBSITE_REPO=github.com/mortenoj/go-graphql-template
DOCKER_REPO=""
NAME="go-graphql-template"

default: test

test: fmt
	@TESTARGS=$(TESTARGS) sh -c "'$(CURDIR)/scripts/test.sh'"
	#go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=2m -parallel=4

tools:
	GO111MODULE=off go get -u golang.org/x/tools/cmd/cover
	GO111MODULE=off go get -u google.golang.org/grpc
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go


bin: fmt
	@DEV=false sh -c "'$(CURDIR)/scripts/build.sh'"

cover:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@TESTARGS=$(TESTARGS) COVER=1 sh -c "'$(CURDIR)/scripts/test.sh'"

fmt:
	@golangci-lint run --enable-all ./...

dev: fmt
	@DEV=true sh -c "'$(CURDIR)/scripts/build.sh'"

quickdev:
	@DEV=true sh -c "'$(CURDIR)/scripts/build.sh'"

run: dev
	@DEV=false sh -c "'$(CURDIR)/scripts/run.sh'"

run-dev: quickdev
	@DEV=true sh -c "'$(CURDIR)/scripts/run.sh'"

build: fmt bin
	@docker build -t $(NAME) . -f build/package/Dockerfile
	@docker tag $(NAME):latest $(NAME):$(VERSION)

acr: build
	@docker tag $(NAME) $(DOCKER_REPO)
	@docker push $(DOCKER_REPO)

release: build
	@docker tag $(NAME) $(DOCKER_REPO):$(VERSION)
	@docker push $(DOCKER_REPO)

taskrunner: quickdev
	@APP=taskrunner sh -c "'$(CURDIR)/scripts/dev.sh' $(PORT)"


.NOTPARALLEL:

.PHONY: bin cover default fmt fmtcheck test dev tools bin run build ecr release lists taskrunner
