### We specify the base image we need for our ### go application
FROM golang:alpine as builder

## Create an /app directory within our ## image that will hold our application source
RUN mkdir /app

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

#Ruun go build to compile the binary executable of our Go program
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

WORKDIR /root/

## Our start command which kicks off our newly created binary executable
CMD ["/app/main"]
