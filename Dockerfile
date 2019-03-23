# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
ENV GOBIN /go/bin
# Copy the local package files to the container's workspace.
ADD . /go/src/goblot
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/goblot
#RUN go get -u github.com/golang/dep
#ENV PATH="/go/bin:${PATH}"
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN dep ensure
#RUN go install github.com/golang/example/outyet
RUN go install
# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/goblot

# Document that the service listens on port 8080.
EXPOSE 8080