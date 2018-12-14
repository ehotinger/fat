FROM golang:1.11.2-alpine AS go-base
RUN apk add --no-cache

FROM go-base AS base
WORKDIR /go/src/github.com/ehotinger/fat
COPY . .
RUN go build ./... && mv fat /usr/bin/fat && mv file.txt /usr/bin/file.txt

FROM scratch
COPY --from=base /usr/bin/fat /usr/bin/fat
COPY --from=base /usr/bin/file.txt /usr/bin/file.txt
ENTRYPOINT [ "fat" ]
CMD [ ]