FROM golang:1.21

ENV SERVER_PORT=:3000 \
    DSN="host=dpg-cn8mn1ol5elc738uad1g-a.oregon-postgres.render.com user=root password=FL9QLWIcvOkx3vM1ndG7ozAL6tyUBoHJ dbname=gorm port=5432"

WORKDIR /app


COPY . .
RUN go get -d -v 
RUN go build -o main -v 
ENTRYPOINT [ "/app/main" ]