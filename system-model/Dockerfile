FROM golang:1.9.1-alpine
EXPOSE 9080
RUN mkdir -p ${GOPATH}/system-model;
WORKDIR ${GOPATH}/system-model/
COPY system-model.go ${GOPATH}/system-model/
RUN go build -o system-model .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/system-model/system-model .
CMD [ "/app/system-model" ]

