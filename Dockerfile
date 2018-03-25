FROM golang:1.9.4
WORKDIR /go/src/app
COPY ./authentication ./authentication
COPY ./config ./config
COPY ./generators ./generators
COPY ./migrations ./migrations
COPY ./models ./models
COPY ./service ./service
COPY ./usermanager ./usermanager
ADD main.go main.go
RUN go-wrapper download
RUN go-wrapper install 
CMD ["go-wrapper", "run"]
EXPOSE 80