Hello world! I'm a gopher

The tiniest go server
Buld it like this

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
run in docker

```
docker build . -t hello-world
docker run -ti -p 3333:3000 hello-world
```
