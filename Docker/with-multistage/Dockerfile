# build stage
FROM golang:latest
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# final stage
# Scratch is minimalistic image which provides security at container level
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
