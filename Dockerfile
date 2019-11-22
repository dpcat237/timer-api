# Create builder image
FROM golang:alpine as builder
ARG TIMER_GITLAB_TOKEN

WORKDIR /go/src/gitlab.com/dpcat237/timer-api

# Download dependencies
RUN apk update && apk upgrade && apk add git
RUN git config --global url."http://dpcat237:${TIMER_GITLAB_TOKEN}@gitlab.com/".insteadOf "https://gitlab.com/"
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only -v

# Build the binary
COPY . .
RUN go install .
EXPOSE 3000 5000
RUN addgroup ustimer && adduser -S -G ustimer ustimer
USER ustimer
ENTRYPOINT ["/go/bin/timer-api"]
