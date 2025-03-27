FROM cgr.dev/chainguard/go AS builder
WORKDIR /app 
COPY . /app
RUN CGO_ENABLED=0 go build -o go-greeter -ldflags="-s -w" . 


FROM cgr.dev/chainguard/static AS prod
COPY --from=builder /app/go-greeter /usr/bin/
ENTRYPOINT ["/usr/bin/go-greeter"]

