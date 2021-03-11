FROM golang:1.14
LABEL maintainer="Silambarasan Karthikeyan"
LABEL name="go-restapi"
LABEL version="1"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/53555/restapi

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

#build
RUN go build 

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./restapi"]
