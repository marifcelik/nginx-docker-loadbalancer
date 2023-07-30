FROM golang
WORKDIR /app
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
RUN ls

FROM scratch
WORKDIR /app
COPY --from=0 /app/main .
CMD [ "./main" ]
