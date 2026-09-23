# patchy agent image for patchy-target: patchy's agent-base plus the Go
# toolchain go.mod asks for and the module cache, so the agent can build and
# test with no network. Built and pushed by .github/workflows/agent-image.yml;
# to build it by hand, use the repository root as the context:
#
#   docker build --platform linux/amd64 -f .patchy/Dockerfile -t patchy-target-agent .
#
# Adapted from examples/agent-images/go in devthenet-labs/patchy.

# The base, pinned to a patchy release and to that release's index digest, so
# a moved tag can never change what this image is built on.
ARG AGENT_BASE=ghcr.io/devthenet-labs/patchy/agent-base:v0.11.7@sha256:e78486306386db6129b41c1f4065429d72d777780a61c866bbf138eeeab32026
FROM ${AGENT_BASE}

# apk needs root; the image drops back to 65532 at the end (the pod runs as
# 65532 whatever the image says).
USER 0:0

# The toolchain go.mod's go directive asks for (go 1.24). Wolfi builds it
# against glibc and puts go on /usr/bin. Bump this with the go directive.
# hadolint ignore=DL3018
RUN apk add --no-cache go-1.24

# GOTOOLCHAIN=local: never fetch another toolchain; fail fast instead.
# GOMODCACHE sits outside HOME (/workspace) and /tmp, which are emptyDirs in
# the pod, so the baked cache survives into it (read-only there).
ENV GOTOOLCHAIN=local \
    GOMODCACHE=/opt/go/pkg/mod

# Bake the module cache from go.mod and go.sum alone (go.sum is optional: the
# module is stdlib-only today). No source enters the image, and this layer is
# rebuilt only when dependencies change. The build cache is dropped; in the
# pod Go's build cache lives under XDG_CACHE_HOME (/tmp/.cache), writable.
WORKDIR /tmp/deps
COPY go.mod go.sum* ./
RUN go mod download \
  && mkdir -p "${GOMODCACHE}" \
  && rm -rf /tmp/deps "${XDG_CACHE_HOME}"

# Offline from here on: modules resolve from the baked cache only, and -mod=mod
# lets go build and go test record a requirement the cache already holds.
ENV GOPROXY=off \
    GOFLAGS=-mod=mod

# Describe this image, not the base (labels are inherited). CI adds revision,
# version and created.
LABEL org.opencontainers.image.source="https://github.com/devthenet-labs/patchy-target" \
      org.opencontainers.image.url="https://github.com/devthenet-labs/patchy-target/tree/main/.patchy" \
      org.opencontainers.image.documentation="https://github.com/devthenet-labs/patchy-target/blob/main/.patchy/README.md" \
      org.opencontainers.image.title="patchy-target agent image" \
      org.opencontainers.image.description="patchy agent image for patchy-target: Go 1.24 with the module cache baked in"

WORKDIR /workspace
USER 65532:65532
