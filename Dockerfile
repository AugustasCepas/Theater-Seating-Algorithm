
FROM golang:1.16-stretch as builder

WORKDIR /src

COPY . .

ARG github_user
ENV github_user=$github_user
ARG github_personal_token
ENV github_personal_token=$github_personal_token

RUN git config \
    --global \
    url."https://${github_user}:${github_personal_token}@@github.com".insteadOf \
    "https://github.com"

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o main

FROM alpine:3.8

COPY --from=builder /src/main /src/main

ENTRYPOINT ["src/main"]