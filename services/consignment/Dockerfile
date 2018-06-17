# Initialize and name a new build stage,
# Set the 'golang:1.9.0' image for subsequent instructions
FROM golang:1.9.0 as builder

# Set the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD 
# instructions that follow it
WORKDIR /go/src/github.com/hoflish/shippy/services/consignment

# Copy the current code into our workdir (filesystem of container)
COPY . .

# Here we're pulling in godep, which is a dependency manager tool,
# we're going to use dep instead of go get, to get around a few
# quirks in how go get works with sub-packages.
RUN go get -u github.com/golang/dep/cmd/dep

# Create a dep project, and run `ensure`, which will pull in all 
# of the dependencies within this directory.
RUN dep init && dep ensure

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 
RUN CGO_ENABLED=0 go build -a -installsuffix cgo .

# Here we're using a second FROM statement,
# this tells Docker to start a new stage with this image.
FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/hoflish/shippy/services/consignment/consignment .

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependencies and
# run time libraries.
CMD ["./consignment"]