# ───────────────────────── Makefile for C++ gRPC ──────────────────────────────
PROJECT := {{ cookiecutter.project_slug }}
BUILD   := build

.PHONY: all build test clean docker-build docker-run

all: build

build:
	mkdir -p $(BUILD)
	cmake -S src/cpp -B $(BUILD) -DCMAKE_BUILD_TYPE=Release
	cmake --build $(BUILD) --target $(PROJECT) -- -j$(shell nproc)

test: build
	cmake --build $(BUILD) --target echo-unit-tests -- -j$(shell nproc)
	cd $(BUILD) && ctest --output-on-failure

clean:
	rm -rf $(BUILD)

docker-build:
	docker build -f Dockerfile.cpp -t $(PROJECT)-cpp:local .

docker-run:
	docker run --rm -p 50051:50051 $(PROJECT)-cpp:local
