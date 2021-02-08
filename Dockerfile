FROM go:1.15-alpine3.13
LABEL maintainer="NuID Developers <dev@nuid.io>"
WORKDIR /nuid/sdk-go
ADD . .
RUN apk add git nodejs npm
RUN go mod download
RUN npm install
ENV PATH=$PATH:/nuid/sdk-go/node_modules/.bin
CMD /bin/sh
