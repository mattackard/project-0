
#build stage
FROM golang:alpine AS builder 
RUN mkdir /app 
ADD . /app
WORKDIR /app
RUN apk add --no-cache git
RUN go get -d
RUN go build -o GoNotesd ./server.go


FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 5555
CMD [ "./GoNotesd" ]