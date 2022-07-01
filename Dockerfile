# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /go/src/github.com/aicyp/escape-dan-back

COPY . /go/src/github.com/aicyp/escape-dan-back

# ----------------------------------------------------------------------
# Lets get dependencies
# RUN go mod init github.com/aicyp/escape-dan-back

# build modules and get dependencies
# RUN go get github.com/go-chi/chi github.com/go-chi/render github.com/lib/pq 
# RUN go build github.com/aicyp/escape-dan-back/models
# RUN go build github.com/aicyp/escape-dan-back/controllers
# RUN go build github.com/aicyp/escape-dan-back/handlers

# RUN go mod tidy
# ----------------------------------------------------------------------

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./go/src/github.com/aicyp/escape-dan-back/bin/main .

# Finally our multi-stage to build a small image
# Start a new stage from scratch
# FROM scratch

# Copy the Pre-built binary file
# COPY --from=builder /go/src/github.com/aicyp/escape-dan-back/bin/main .

# ENTRYPOINT ["/bin/sh"]

# Run executable
CMD ["./go/src/github.com/aicyp/escape-dan-back/bin/main"]