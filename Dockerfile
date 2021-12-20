FROM docker.io/golang:latest

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build .

FROM scratch

COPY --from=0 /build/member-club /member-club

CMD ["./member-club"]

EXPOSE 8080