CGO_ENABLED=0 GOOS=linux go build -o printwebhook .

docker build -t printwebhook:v0.1  .

rm printwebhook