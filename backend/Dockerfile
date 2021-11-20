FROM golang:1.17.3 as development
WORKDIR /go/src/github.com/koyashiro/rdbms-playground
COPY go.mod go.sum ./
RUN go mod download
RUN go get -u github.com/cosmtrek/air
COPY . .
CMD ["air"]

FROM development as build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/static:latest as production
WORKDIR /root
COPY --from=build /go/src/github.com/koyashiro/rdbms-playground/app /usr/bin/app
EXPOSE 80
CMD ["/usr/bin/app"]
