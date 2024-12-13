################## ビルドステージ ##################
FROM golang:1.23.4 as build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ .
RUN go build -tags lambda.norpc -o main main.go

################## 実行ステージ ##################
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]
