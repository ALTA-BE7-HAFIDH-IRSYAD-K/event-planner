FROM golang:1.17

WORKDIR /app

COPY . .

RUN go build -o event-planner

EXPOSE 8080

CMD ./event-planner
