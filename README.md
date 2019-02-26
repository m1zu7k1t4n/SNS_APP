# Go言語を使ってバックエンド作ってみる

## Commands

起動

```sh
docker-compose up -d
```

MySQLサーバーに接続

```sh
mysql -hlocalhost -utest -ptest test
```

それぞれのコンテナに接続

```sh
docker-compose exec {serviceName} sh
```
