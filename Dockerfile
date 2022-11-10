from golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY internal/*.go ./

RUN go build -o /docker-card-quiz

EXPOSE 9999:9999

CMD ["/docker-card-quiz"]