FROM golang:1.18-alpine

WORKDIR /data

RUN go install golang.org/x/tools/cmd/present@latest

COPY . .

EXPOSE 8080

ENTRYPOINT ["/go/bin/present", "-base", "/go/pkg/mod/golang.org/x/tools@v0.1.11/cmd/present", "-use_playground", "-content", "/data", "-http", ":8080"]