# material-gakujo

gakujo を今風にしたいプロジェクト

## Requirements

- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [go](https://golang.org/doc/install)

## build & run

```console
$ pwd
/path/to/material-gakujo

$ go mod vendor # 正直なくていい気がする

$ docker-compose up -d --build
```

完了後、<http://localhost:5000> にアクセスする。
