FROM golang:1.4.2
MAINTAINER Xabier Larrakoetxea <slok69@gmail.com>

# Create the user/group for the running stuff
RUN groupadd -g 1000 baton
RUN useradd -m -u 1000 -g 1000 baton
RUN chown baton:baton -R /go
USER baton

# Install dependencies
RUN go get github.com/tools/godep
COPY . /tmp/daton
RUN cd /tmp/daton && godep restore

WORKDIR /go/src/github.com/slok/baton