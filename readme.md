# Very nice code verifier

![GoLand total statements coverage 76.3%](https://img.shields.io/badge/coverage-76.3%25-green)

Сервис для генерации кодов с TTL и их верификации

<hr/>

Была использована KeyDB в силу наличия expire для HSET

Без лишних слов:

Установить зависимости `go mod download`

Запуск - `make run`

Test with coverage - `make test`

Build to executable - `make build` -> `./build/runner`

## Docker
```bash
$ make docker
$ make compose
```
Image:
https://hub.docker.com/r/supermantelecomworker/awesomecodeservice

manual тесты:
```bash
$ curl --location 'localhost:8080/api/v1/send' \
--header 'Content-Type: application/json' \
--data '{
    "number": "+7 (999) 888-77-66"
}'

> {"requestId":"86ae27ae-df99-11ed-bf70-2af8bc618b50","code":4386}
```

Positive:
```bash
$ curl --location 'localhost:8080/api/v1/verify' \
--header 'Content-Type: application/json' \
--data '{"requestId":"86ae27ae-df99-11ed-bf70-2af8bc618b50","code":4386}'

> {"verifiedAt":1682008671}
```

Negative:
```bash
$ curl --location 'localhost:8080/api/v1/verify' \
--header 'Content-Type: application/json' \
--data '{"requestId":"86ae27ae-df99-11ed-bf70-2af8bc618b50","code":4388}'

> {"error":"invalid code"}
```

Attempts exceeded:
```bash
$ curl --location 'localhost:8080/api/v1/verify' \
--header 'Content-Type: application/json' \
--data '{"requestId":"86ae27ae-df99-11ed-bf70-2af8bc618b50","code":4388}'

> {"error":"Verification attempts limit has been reached"}
```

# TODO
- max possible test coverage