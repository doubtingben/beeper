FROM golang:1.8

WORKDIR /go/src/github.com/doubtingben/beeper
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["beeper"]
