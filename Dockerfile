FROM golang:latest as build
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

ENV PORT=6969

RUN go build  ./cmd/main.go


WORKDIR /app
COPY --from=build /root/main /app/

CMD ["./main"]