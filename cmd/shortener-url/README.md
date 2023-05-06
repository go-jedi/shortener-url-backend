# http API

```sh
# dev
http://localhost:8080/
```

### Добавление ссылки (POST)

````sh
Url: http://localhost:8080/a?url=
```js
{
  "query": {
    "url": string, <- ссылка
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное добавление ссылки" <- typeof string
  -result: "" <- typeof string
OR:
  -status: 200 <- typeof int
  -message: "ошибка добавления ссылки" <- typeof string
  -result: "" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в query";
  -message: "ошибка выполнения функции uid из базы данных, {err}";
  -message: "ошибка выполнения функции url_add из базы данных, {err}";
````

### Получение ссылки (GET)

````sh
Url: http://localhost:8080/s/:uid
```js
{
  "params": {
    "uid": string, <- uid ссылки
  }
}
```sh
RETURN:
  Редирект 302
OR:
  -status: 200 <- typeof int
  -message: "ошибка получения ссылки" <- typeof string
  -result: "" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в query";
  -message: "ошибка выполнения функции url_get из базы данных, {err}";
````
