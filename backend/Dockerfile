FROM golang:1.22.4

WORKDIR /app

# copy go package manager
COPY go.mod ./
COPY  go.sum ./

RUN go mod download

# copy all project files
COPY . .

# run to compile go program
RUN go build -v -o main .

CMD [ "/app/main" ]