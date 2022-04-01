FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dragod812/go-http-server-template

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies and install
RUN make install

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./bin/server"]

