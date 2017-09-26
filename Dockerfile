FROM golang:1.9.0 as build
COPY . /go/src/github.com/seemethere/misto
WORKDIR /go/src/github.com/seemethere/misto
RUN CGO_ENABLED=0 go build -a -o /misto -ldflags '-extldflags "-static"' misto.go

FROM scratch
COPY --from=build /misto /misto
ENTRYPOINT [ "/misto" ]
