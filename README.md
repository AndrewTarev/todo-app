# TodoList

## Описание
Это приложение служит для создания TODO списков. В нем есть регистрация и авторизация, возможность создавать списки, а в списках заводить отдельные задачи.

## Структура проекта
```
.
├── README.md
├── cmd
│   └── main.go
├── configs                                # Конфиг
│   └── config.yml
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── handler                            # Обработчик запросов (эндпоинты)
│   │   ├── auth.go
│   │   ├── handler.go
│   │   ├── item.go
│   │   └── list.go
│   ├── repository                         # Работа с БД
│   │   ├── postgres.go
│   │   └── repository.go
│   └── service                            # Бизнес логика
│       └── service.go
├── schema                                 # БД миграции
│   ├── 000001_init.down.sql
│   ├── 000001_init.up.sql
│   └── info.txt
├── server.go                              # Настройки запуска и остановки сервера
├── todo.go                                # Модели списков задач
└── user.go                                # Модель Юзера

```
## Используемые технологии

1. Gin - веб фреймворк
2. PostgreSQL - База данных
3. sqlx - для работы с БД
4. viper - работа с конфигом