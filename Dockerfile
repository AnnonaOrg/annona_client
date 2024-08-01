FROM alpine:3.18 as tdlib-builder

ENV LANG en_US.UTF-8
ENV TZ UTC

ARG TD_COMMIT

RUN apk update && \
    apk upgrade && \
    apk add --update \
        build-base \
        ca-certificates \
        ccache \
        cmake \
        git \
        gperf \
        linux-headers \
        openssl-dev \
        php \
        php-ctype \
        readline-dev \
        zlib-dev && \
    git clone "https://github.com/tdlib/td.git" /src && \
    cd /src && \
    git checkout ${TD_COMMIT} && \
    mkdir ./build && \
    cd ./build && \
    cmake \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_INSTALL_PREFIX:PATH=/usr/local \
        .. && \
    cmake --build . --target prepare_cross_compiling && \
    cd .. && \
    php SplitSource.php && \
    cd build && \
    cmake --build . --target install && \
    ls -lah /usr/local


FROM golang:alpine3.18 as go-builder

ENV LANG en_US.UTF-8
ENV TZ UTC

RUN set -eux && \
    apk update && \
    apk upgrade && \
    apk add \
        bash \
        build-base \
        ca-certificates \
        curl \
        git \
        linux-headers \
        openssl-dev \
        zlib-dev

WORKDIR /src

COPY --from=tdlib-builder /usr/local/include/td /usr/local/include/td/
COPY --from=tdlib-builder /usr/local/lib/libtd* /usr/local/lib/
COPY . /src

RUN go build \
    -a \
    -trimpath \
    -o annona_client \
    -ldflags "-s -w -buildid=" \
    "./cmd/annona_client" && \
    ls -lah


FROM alpine:3.18

#ENV TZ Asia/Shanghai
RUN apk upgrade --no-cache && \
    apk add --no-cache \
            ca-certificates \
            libstdc++ \
            tzdata \
            musl-locales musl-locales-lang
ENV LANG en_US.UTF-8
ENV TZ UTC
WORKDIR /app

COPY --from=go-builder /src/annona_client .

ENTRYPOINT ["./annona_client"]
#CMD ["./annona_client"]