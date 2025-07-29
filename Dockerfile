# ----- BUILD STAGE -----

FROM golang@sha256:adfbe17b774398cb090ad257afd692f2b7e0e7aaa8ef0110a48f0a775e3964f4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o student-api

# ----- RUN STAGE -----

FROM alpine@sha256:b3119ef930faabb6b7b976780c0c7a9c1aa24d0c75e9179ac10e6bc9ac080d0d AS runtime

WORKDIR /app

COPY --from=builder /app/student-api .

EXPOSE 8080

ENTRYPOINT ["./student-api"]