FROM golang:1.9.1-alpine
EXPOSE 9090
RUN mkdir -p ${GOPATH}/system-standards;
WORKDIR ${GOPATH}/system-standards/
COPY src/system-standards.go ${GOPATH}/system-standards/
RUN go build -o system-standards .
RUN apk --no-cache add ca-certificates
CMD [ "./system-standards" ]
