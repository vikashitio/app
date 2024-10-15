FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main ./main.go
RUN chmod +x main
EXPOSE 3000
CMD [ "./main" ]