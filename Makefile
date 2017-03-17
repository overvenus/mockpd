GO=GO15VENDOREXPERIMENT="1" CGO_ENABLED=0 go

build:
	rm -rf vendor && ln -s _vendor/vendor vendor
	$(GO) build -o bin/pd-server main.go
	rm -rf vendor

update:
	which glide >/dev/null || curl https://glide.sh/get | sh
	which glide-vc || go get -v -u github.com/sgotti/glide-vc
	rm -rf vendor && mv _vendor/vendor vendor || true
	rm -rf _vendor
	glide update --strip-vendor --skip-test
	@echo "removing test files"
	glide vc --only-code --no-tests
	mkdir -p _vendor
	mv vendor _vendor/vendor
