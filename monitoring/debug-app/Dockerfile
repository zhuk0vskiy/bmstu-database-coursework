FROM golang:1.22
WORKDIR /app
COPY go.mod ./
COPY ../go.sum ./
COPY main.go ./
CMD [ "go", "run", "./main.go" ]