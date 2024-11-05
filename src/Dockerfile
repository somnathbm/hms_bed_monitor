FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./main.go
COPY ./api ./api/
COPY ./db ./db/
COPY ./bed_stats ./bed_stats/
RUN go build -o ./dist/bed_stats.sh

FROM golang:alpine as RUN
WORKDIR /app
COPY --from=build /app/dist/bed_stats.sh ./
EXPOSE 8080
CMD ["./bed_stats.sh"]
