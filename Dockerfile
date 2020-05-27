FROM golang:1.14 as builder

WORKDIR /earmuffs
COPY . /earmuffs

RUN pwd
RUN go get -d -v

# Statically compile our app for use in a distroless container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o app .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

# Copy over binary and words dir
COPY --from=builder /earmuffs/app /app
COPY --from=builder /earmuffs/data/words data/words

ENTRYPOINT ["/app"]
