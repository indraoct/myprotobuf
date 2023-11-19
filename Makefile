GO_MODULE := myprotobuf

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: verify
verify:
	go mod verify

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
else
	rm -fR ./protogen
endif

.PHONY: protoc
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*proto

.PHONY: build
build: clean protoc tidy vendor verify

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run