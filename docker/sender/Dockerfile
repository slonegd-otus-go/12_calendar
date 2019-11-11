FROM golang:1.13
WORKDIR /app
COPY . .
RUN go build -o mycalendar .
ENTRYPOINT ["./mycalendar", "sender"]
