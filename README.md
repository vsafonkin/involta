1. Run reindexer docker container

$ docker run -p9088:9088 -p6534:6534 --name reindex -it reindexer/reindexer

2. Build and run app

$ git clone https://github.com/vsafonkin/involta
$ cd involta
$ make
$ ./bin/main -config ./config.yaml

3. Test via curl

all documents:
$ curl -w'\n' localhost:8080/testdb/docs

document by id:
$ curl -w'\n' localhost:8080/testdb/docs?id=3

update (if `id` exists) or insert new:
$ curl -w'\n' http://localhost:8080/testdb/docs \
--request "POST" \
--data '{"id":1,"sort":453,"content":[{"id":0,"content":"Bob"},{"id":1,"content":"Alice"}]}'
