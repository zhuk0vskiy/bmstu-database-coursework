FROM golang:1.22

WORKDIR /app

ENV MONITORING_URL ${MONITORING_URL}

COPY monitoring_app.go monitoring_app.go
COPY go.mod go.mod

RUN go mod tidy

CMD ["go", "run", "monitoring_app.go"]