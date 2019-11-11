FROM golang:1.13
WORKDIR /app
COPY . .
RUN go build -o mycalendar .
ENTRYPOINT ["./mycalendar", "api" , "--host=0.0.0.0"]
