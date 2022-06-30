FROM golang AS base

FROM base AS dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/github.com/aicyp/escape-dan-back/

CMD ["air"]