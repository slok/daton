FROM golang:1.4.2
MAINTAINER Xabier Larrakoetxea <slok69@gmail.com>

# Create the user/group for the running stuff
RUN groupadd -g 1000 daton
RUN useradd -m -u 1000 -g 1000 daton
RUN chown daton:daton -R /go
USER daton

# Install dependencies
RUN go get github.com/tools/godep
COPY . /tmp/daton
RUN cd /tmp/daton && godep restore

WORKDIR /go/src/github.com/slok/daton
