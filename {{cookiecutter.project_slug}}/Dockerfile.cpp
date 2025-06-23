# ─── build stage ───────────────────────────────────────────────────────────────
FROM debian:bookworm-slim AS builder
WORKDIR /src

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential cmake git ca-certificates pkg-config \
        libssl-dev zlib1g-dev \
        protobuf-compiler libprotobuf-dev \
        libgrpc++-dev grpc-tools \
    && rm -rf /var/lib/apt/lists/*

# copy proto & code, generate & build
COPY proto proto
COPY src/cpp src
RUN cmake -S src/cpp -B build -DCMAKE_BUILD_TYPE=Release && \
    cmake --build build --target {{ cookiecutter.project_slug }} -- -j$(nproc)

# ─── runtime stage ─────────────────────────────────────────────────────────────
FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        libssl3 zlib1g \
        libprotobuf32 libgrpc++1 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/build/{{ cookiecutter.project_slug }} /usr/local/bin/{{ cookiecutter.project_slug }}

EXPOSE 50051
CMD ["{{ cookiecutter.project_slug }}"]
