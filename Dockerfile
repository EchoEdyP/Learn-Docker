FROM golang:1.19.1

COPY main.go /app/main.go

CMD ["go", "run", "/app/main.go"]