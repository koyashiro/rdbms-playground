FROM golang:latest as development
WORKDIR /go/src/github.com/koyashiro/postgres-playground
COPY go.mod go.sum ./
RUN go mod download
ENV PORT=80
COPY . .

FROM golang:latest as build
WORKDIR /go/src/github.com/koyashiro/postgres-playground
COPY . .
COPY --from=development /go/pkg /go/pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/static:latest as production
WORKDIR /root
ENV PORT=80
COPY --from=build /go/src/github.com/koyashiro/postgres-playground/app /usr/bin/app
EXPOSE 80
CMD ["/usr/bin/app"]
