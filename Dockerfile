# Specify the base image
FROM golang:1.13.4-alpine3.10

# Create /apps directory in the image
# This dir holds app source files
RUN mkdir /apps

# Copy everything in the project
# root dir to the new dir /apps
ADD . /apps

# Specify that we would need to execute
# any commands inside the /apps directory
WORKDIR /apps

# Specify command(s):
# 'go build' to compile to binary
RUN go build -o main .

# Then run the program
CMD ["/apps/main"]