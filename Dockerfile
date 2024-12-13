################## ビルドステージ ##################
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine3.21 AS builder
WORKDIR /lambda

COPY ./go.mod ./go.sum ./
RUN go mod download -x

COPY ./cmd ./cmd
COPY ./cost ./cost
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app cmd/main.go && \
    chmod +x app

################## 実行ステージ ##################
FROM public.ecr.aws/lambda/provided:al2023 AS runner
COPY --from=builder /lambda/app ./app
RUN chmod +x app
ENTRYPOINT [ "./app" ]