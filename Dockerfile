FROM golang:1.12.4-alpine3.9 as builder
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod tidy && go mod vendor
COPY . .
RUN PORT=8888 BUILD_MODE=debug ginkgo -v -r
RUN go build . -o locrep-go.exe

FROM alpine:3.9.3
WORKDIR /
COPY --from=builder /go/src/app/locrep-go.exe /
ENTRYPOINT /locrep-go.exe