FROM gitpod/workspace-full

USER gitpod

RUN brew install go-task/tap/go-task
RUN go install github.com/spf13/cobra-cli@latest
