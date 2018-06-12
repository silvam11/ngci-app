FROM golang:1.9.1-alpine
EXPOSE 9080
RUN mkdir -p ${GOPATH}/system-model;
WORKDIR ${GOPATH}/system-model/
COPY system-model.go ${GOPATH}/system-model/
RUN go build -o system-model .
RUN apk --no-cache add ca-certificates
CMD [ "./system-model" ]
