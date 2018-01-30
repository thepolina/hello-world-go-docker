Hello world! I'm a gopher

The tiniest go server
Buld it like this

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
run in docker

```
docker build . -t hello-world
docker run --rm -p 3000:3333 hello-world
```
