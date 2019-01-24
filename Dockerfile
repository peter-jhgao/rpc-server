FROM swr.cn-north-1.myhuaweicloud.com/hwstaff_g00420596/golang:v1

WORKDIR /go/src

ADD . ./

RUN go build

CMD ["./src"]
