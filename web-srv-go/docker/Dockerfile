FROM golang:latest as builder

WORKDIR /app

# Retrieve application dependencies using go mod.
# Allows for the cache to not be invalidated unless changes to mod/sum files are made.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY ./src ./

# Build the binary.
# -o app specifies the name of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app


# Start from scratch
FROM scratch

WORKDIR /

ADD ./html /html
COPY --from=builder /app/app /

ENTRYPOINT ["/app"]

EXPOSE 80