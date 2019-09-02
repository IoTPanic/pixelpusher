FROM golang
ADD . /go/src/github.com/IoTPanic/pixelpusher
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd /go/src/github.com/IoTPanic/pixelpusher; dep ensure -v
RUN go install -v github.com/IoTPanic/pixelpusher
ENTRYPOINT [ "/go/bin/pixelpusher" ]
EXPOSE 8080
