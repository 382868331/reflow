FROM golang:1.26
WORKDIR /src
COPY . .
RUN go mod download
CMD ["go", "test", "./..."]
