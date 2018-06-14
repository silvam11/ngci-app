FROM golang:1.9.1-alpine
EXPOSE 8080
RUN mkdir -p ${GOPATH}/vrealize-controller;
WORKDIR ${GOPATH}/vrealize-controller/
COPY src/vrealize-controller.go ${GOPATH}/vrealize-controller/
RUN go build -o vrealize-controller .
RUN apk --no-cache add ca-certificates
CMD [ "./vrealize-controller" ]
