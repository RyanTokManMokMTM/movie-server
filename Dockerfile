FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download  # download all package

COPY . .
# server.yaml
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/movieservice ./movieservice.go


FROM scratch
#there is nothing -> we need to copy all the thing that we need form builder

# copy all built think into this image
COPY --from=builder /usr/share/zoneinfo/Asia/Hong_Kong /usr/share/zoneinfo/Asia/Hong_Kong
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

#set time zone
ENV TZ Asia/Taipei

#set workdir
WORKDIR /app

#copy all file from builder
COPY --from=builder /app/movieservice /app/movieservice
COPY --from=builder /app/etc /app/etc

# starting the server after started container
CMD ["./movieservice","-f","etc/movieservice.yaml"]

#FROM golang:alpine AS goBuild
#
#LABEL stage=gobuilder
#
#ENV CGO_ENABLED 0
#ENV GOPROXY https://goproxy.cn,direct
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#
#RUN apk update --no-cache && apk add --no-cache tzdata
#
#WORKDIR /build
#
#ADD go.mod .
#ADD go.sum .
#RUN go mod download
#COPY . .
#COPY ./etc /app/etc
#RUN go build -ldflags="-s -w" -o /app/movieservice ./movieservice.go
#
#
#FROM scratch
#
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
#ENV TZ Asia/Shanghai
#
#WORKDIR /app
#COPY --from=builder /app/movieservice /app/movieservice
#COPY --from=builder /app/etc /app/etc
#
#CMD ["./movieservice", "-f", "etc/movieservice.yaml"]
