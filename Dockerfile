#FROM golang:1.14
#
#WORKDIR /go/src/dumbstored
#COPY . .
#
#RUN go get -d -v ./...
#RUN go install -v ./...
#
#CMD ["dumbstored"]

FROM debian:10

COPY ./dumbstored /bin/dumbstored

CMD ["/bin/dumbstored"]
