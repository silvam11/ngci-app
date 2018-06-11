FROM golang:1.9.1-alpine
EXPOSE 8080
RUN mkdir -p ${GOPATH}/web-gateway;
WORKDIR ${GOPATH}/web-gateway/
COPY web-gateway.go ${GOPATH}/web-gateway/
RUN go build -o web-gateway .
RUN apk --no-cache add ca-certificates
CMD [ "./web-gateway" ]

