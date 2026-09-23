# .patchy

Configuration for [patchy](https://github.com/devthenet-labs/patchy), the pipeline that investigates and fixes this
repository's code-scanning alerts. Its agent runs in a sandboxed pod with no network, so the image it runs in has to
carry the Go toolchain and every module the build needs.

- [`Dockerfile`](Dockerfile): the agent image. Patchy's `agent-base` (pinned by digest) plus Go 1.24 and a module
  cache baked from `go.mod`/`go.sum`, with `GOPROXY=off` so builds never reach for the network.
- [`.github/workflows/agent-image.yml`](../.github/workflows/agent-image.yml): builds that Dockerfile for `linux/amd64`
  on every push to `main` that changes it, `go.mod` or `go.sum` (or on demand), and pushes it to ECR as
  `:sha-<short commit>` through GitHub OIDC. No stored secrets.
- [`agent.yaml`](agent.yaml): declares the image patchy runs the agent in. `image` is the only key.

Patchy uses the declared image only when its operator has enabled repository images and allowlisted this ECR path;
otherwise the agent runs in patchy's default image and this directory changes nothing. Patchy resolves the tag to a
digest once per finding, so a later push cannot change the environment of a finding already in flight.

## Updating the image

1. Change `Dockerfile`, `go.mod` or `go.sum` and push to `main`. The workflow pushes a new `:sha-<short commit>` tag.
2. Point `agent.yaml` at that tag in a follow-up commit (editing `agent.yaml` or this README does not rebuild).

ECR tags are immutable, so running the workflow by hand on a commit that already has an image fails at the push; that
is expected. Bump the Go version in the Dockerfile together with the `go` directive in `go.mod`.

To check an image the way patchy will, with the workstation CLI:

```sh
patchy check image <image from agent.yaml> --allow 377946145366.dkr.ecr.us-east-1.amazonaws.com/patchy/app-envs/ --run
```
