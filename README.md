##

```bash
# develpment
go run main.go --port=3000

# production
go run main.go --port=3000 --mode=release

# docker
docker build -t gin_server_cli:v0.1 .
# run docker
docker run -it -p 3000:3000 --rm --name gin_server_cli gin_server_cli:v0.1
```
