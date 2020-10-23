FROM golang:1.15-alpine as BUILD

WORKDIR /go/src/github.com/kakanghosh/file-server

COPY *.go .
COPY app app/
COPY *.mod .
COPY *.sum .

RUN go get
RUN go mod tidy
RUN go build

FROM alpine

WORKDIR /app
RUN mkdir -p static-files
COPY templates/ templates/
COPY assets/ assets/
COPY --from=BUILD /go/src/github.com/kakanghosh/file-server/fileserver fileserver

EXPOSE 8080
CMD [ "./fileserver" ]
