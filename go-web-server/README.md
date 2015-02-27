To build:
```sh
docker build -t go-web-server .
mkdir -p go-web-server && cd go-web-server
docker run go-web-server > go-web-server
chmod +x go-web-server
cat > Dockerfile << END
FROM scratch
ADD go-web-server /go-web-server
ENTRYPOINT ["/go-web-server"]
END
docker build -t go-web-server .
docker run -p8080:8080 go-web-server
```



