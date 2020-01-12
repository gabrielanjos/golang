FROM golang:1.9

# Add Maintainer Info
LABEL maintainer="Gabriel Anjos <gabriel@dosanjos.net.br>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src/ ./

# Build the Go app
RUN  go get -d -v github.com/jinzhu/gorm && go get -d -v github.com/gabrielanjos/golang/src  && go build -o main .

# Command to run the executable
#CMD ["./main"]

ENTRYPOINT ["./main"]