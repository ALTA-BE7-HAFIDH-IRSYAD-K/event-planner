FROM golang:1.17

WORKDIR /app

COPY . .

RUN go build -o event-planner

EXPOSE 80

CMD ./event-planner
