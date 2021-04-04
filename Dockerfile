FROM golang
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . /usr/app
ENV GOPROXY="https://goproxy.io"
RUN go build  -i -v -ldflags '-w -s' -o ./bin/main
COPY ./public /usr/app/bin/public
COPY ./config/config.yaml /usr/app/bin/config
EXPOSE 3000
CMD ["./bin/main --port=3000 --mode=release"]