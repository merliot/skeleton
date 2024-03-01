# syntax=docker/dockerfile:1

FROM merliot/device:main

WORKDIR /app
COPY . .

RUN go generate ./...
RUN go build -tags prime -o /skeleton ./cmd

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/skeleton"]
