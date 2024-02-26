FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o hermes main.go

ENTRYPOINT [ "./hermes" ]
