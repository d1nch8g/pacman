pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test:
	docker build -t fmnx.su/core/pacman .
	docker run --rm -it fmnx.su/core/pacman go test ./...
