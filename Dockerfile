
FROM golang:1.19-alpine3.15

WORKDIR /projects/data/

COPY go.mod go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o se.exe ./cmd/

EXPOSE 8080

CMD [ "./se.exe" ]
