FROM golang:1.20

WORKDIR /usr/src/app

# Download all the dependencies
RUN go install github.com/14Artemiy88/termPaint@latest && export PATH=$GOPATH/bin

# Run the executable
CMD ["termPaint"]
