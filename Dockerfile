FROM golang:1.16-alpine

RUN apk update
RUN apk add curl
RUN apk add make
RUN apk add tar
RUN apk add gzip
RUN apk add vim

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# 安裝 golang migrate
# 參考資料 https://stackoverflow.com/a/69478562
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 $GOPATH/bin/migrate

CMD air