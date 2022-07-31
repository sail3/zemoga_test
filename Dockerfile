FROM golang:alpine3.13 as builder
RUN apk update && apk upgrade && apk add build-base git make sed
RUN go get github.com/silenceper/gowatch
# install migrate to perform migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# swagger addition via swagger-ui inyection
WORKDIR /go/src/github.com/sail3/zemoga_test/swagger
COPY ./oas/oas.yml ./swagger.yml
RUN git clone https://github.com/swagger-api/swagger-ui && \
 cp -r swagger-ui/dist/. . && rm -r swagger-ui/ && sed -i 's+https://petstore.swagger.io/v2/swagger.json+/swagger/swagger.yml+g' index.html

WORKDIR /go/src/github.com/sail3/zemoga_test
COPY . .

RUN GIT_COMMIT=$(git rev-parse --short HEAD) && \
  go build -o service -ldflags "-X 'github.com/sail3/zemoga_test/internal/config.serviceVersion=$GIT_COMMIT'" ./cmd/api

FROM alpine:3.13 

COPY --from=builder /go/src/github.com/sail3/zemoga_test/service /
COPY --from=builder /go/src/github.com/sail3/zemoga_test/swagger /swagger
COPY --from=builder /go/src/github.com/sail3/zemoga_test/migrations /migrations

ENTRYPOINT [ "./service" ]
