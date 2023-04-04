FROM golang:1.18.4 AS build

WORKDIR /app   

COPY main.go go.mod ./

RUN go mod download

# build runner
RUN go build -o loadjob

FROM scratch

WORKDIR /app

COPY --from=build /app/loadjob .

CMD [ "./loadjob" ]