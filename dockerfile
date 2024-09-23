FROM golang:1.23 as build
WORKDIR /build
COPY . .
RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server

FROM scratch
COPY --from=build /build/api-server /
EXPOSE 8080
CMD ["/api-server"]