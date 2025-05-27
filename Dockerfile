FROM golang:1.24

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

WORKDIR /app

# Copy only go.mod and go.sum first to use Docker cache effectively
COPY go.mod go.sum ./

# Download dependencies (will be cached if go.mod/go.sum not changed)
RUN go mod download

# Now copy the rest of the code
COPY . .

CMD ["make", "run"]
