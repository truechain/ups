FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/truechain/ups && \
  (cd truechain-engineering-code && make gups) && \
  cp truechain-engineering-code/build/bin/gups /gups && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /truechain-engineering-code && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/gups"]
