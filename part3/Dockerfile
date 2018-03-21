FROM golang:1
RUN mkdir -p /go/src/github.com/dhopkins-va/cmpt-350/part2
COPY main.go /go/src/github.com/dhopkins-va/cmpt-350/part2/main.go
RUN GOOS=linux CGO_ENABLED=0 go build -o /main github.com/dhopkins-va/cmpt-350/part2

FROM scratch
COPY --from=0 /main /main
EXPOSE 80
ENTRYPOINT ["/main"]
