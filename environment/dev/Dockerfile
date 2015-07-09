FROM golang:1.4.2
MAINTAINER Xabier Larrakoetxea <slok69@gmail.com>

# Create the user/group for the running stuff
RUN groupadd -g 1000 daton
RUN useradd -m -u 1000 -g 1000 daton
RUN chown daton:daton -R /go
RUN chmod 777 /usr/src/go/pkg/tool/linux_amd64/
USER daton

# Install handy dependencies/tools
RUN go get github.com/tools/godep
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/axw/gocov/gocov

# Install project dependencies
COPY . /tmp/daton
RUN cd /tmp/daton && godep restore

WORKDIR /go/src/github.com/slok/daton
