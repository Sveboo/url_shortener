# Url Shortener

🌍 **[English](../README.md) ∙ [Русский](local_docs/README_ru.md)**

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Alpine Linux](https://img.shields.io/badge/Alpine_Linux-%230D597F.svg?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

Этот модуль реализует простой REST API, который принимает полный URL-адрес ресурса
и возвращает короткую ссылку для доступа к тому же ресурсу.
Используется кодирование с помощью Base62.

## Table of Contents
- [Url Shortener](#url-shortener)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Direct Usage](#direct-usage)
    - [Docker](#docker)
    - [Allowed Usage](#allowed-usage)
  - [Endpoints](#endpoints)
  - [Swagger](#swagger)
  - [Debug Build](#debug-build)
  - [How It Works](#how-it-works)
  - [Roadmap](#roadmap)
  - [Issues](#issues)

## Getting Started

### Direct Usage
1. Установите go версии 1.21 или новее. Посмотрите [официальную документацию](https://go.dev/doc/install).
2. Установите зависимости. Запустите `go mod tidy`.
3. Запустите приложения. Посмотрите [разрешённые режимы](#allowed-usage).
4. Используйте [эндпоинты](#endpoints).

### Docker
1. Установите docker engine. Посмотрите [официальную документацию](https://docs.docker.com/engine/install/).
2. Соберите образ. Запустите `docker build -t url_shortener:latest .` или используйте существующий образ на [Dockerhub](https://hub.docker.com/).
3. Запустите приложения. Посмотрите [разрешённые режимы](#allowed-usage).
4. Используйте [эндпоинты](#endpoints).

### Allowed Usage
 - Запускается без флагов. Данные для URL-адресов будут храниться непосредственно в памяти (Не рекомендуется для production использования).
 - Запускается с флагом `-d` и установленной переменной окружения `PATH_TO_DB`. Данные для URL-адресов будут храниться в базе данных PostgreSQL, доступной по ссылке из переменной окружения.

## Endpoints
| Request type                         | Endpoint | Parameters type | Parameters            | Returns               |
|--------------------------------------|----------|-----------------|-----------------------|-----------------------|
|<span style="color:yellow">POST</span>| /        | body(json)      | url - URL ресурса     | Короткий хэш          |
|<span style="color:green">GET</span>  | /${hash} | query string    | короткий хэш          | URL ресурса           |
|<span style="color:green">GET</span>  | /docs    | -               | -                     | Документация swagger  |

## Swagger

Приложение также предоставляет документацию swagger, доступную с помощью эндпоинта `/docs`. Посмотрите [эндпоинты](#endpoints).

## Debug Build

> [!CAUTION]
> **Не используйте debug build на production!**

Для сборки с тегом `debug` приложение также обслуживает эндпоинт `/debug/pprof` с метриками для профиля pprof.

## How It Works

Когда вы отправляете URL внешнего ресурса, он проверяется и хэшируется с использованием случайного числа и кодировки base62.
Если такой хэш уже существует, предпринимается разумное количество попыток пересоздать его.

> [!NOTE]
> Не рекомендуется использовать хранилище URL в памяти, так как не существует политики удаления устаревших URL-адресов.
> Пожалуйста, используйте внешнюю базу данных с настроенной политикой удаления, чтобы избежать нехватки ресурсов хранилища.

После получения короткого имени вы можете отправить запрос GET на сервер и получить полный URL исходного ресурса.

## Roadmap
- [x] Покрыть код тестами
- [ ] Сделать Docker-image доступным публично
- [ ] В будущем планируется добавить возможность автоматического перенаправления на исходный ресурс вместо возврата его в теле ответа

## Issues

Если вы заметили какую-либо ошибку или уязвимость, не стесняйтесь создавать Issue с полным описанием найденной проблемы.
Кроме того, будет здорово, если вы предложите свое решение проблемы через Pull Request.