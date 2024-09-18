FROM golang:1.22-alpine

WORKDIR .

COPY src src
COPY run_backend.sh run_backend.sh
COPY go.mod go.mod
COPY go.sum go.sum

# RUN go mod tidy

#CMD ["go", "run", "src/cmd/main.go"]
CMD ["go", "mod", "tidy"]
#CMD ["./run_backend.sh"]


