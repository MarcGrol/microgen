FROM golang
RUN go get github.com/tools/godep
RUN git clone https://github.com/MarcGrol/microgen.git /go/src/github.com/MarcGrol/microgen
RUN godep go install github.com/MarcGrol/microgen
# executable located in: /go/bin/microgen
