FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./main.go
COPY ./api ./api/
COPY ./db ./db/
RUN go build -o ./dist/hms-bed-monitor.sh

FROM golang:alpine AS run
WORKDIR /app
COPY --from=build --chmod=500 /app/dist/hms-bed-monitor.sh ./
EXPOSE 8080
CMD ["./hms-bed-monitor.sh"]
