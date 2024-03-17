# Url Shortener

🌍 **[English](README.md) ∙ [Русский](local_docs/README_ru.md)**

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Alpine Linux](https://img.shields.io/badge/Alpine_Linux-%230D597F.svg?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)


This module implements simple REST API that takes full URL to resource
and returns a short link to access the same resource.
Base62 encoding is used.

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
1. Install go version 1.21 or higher. See [official docs](https://go.dev/doc/install).
2. Install dependencies. Run `go mod tidy`.
3. Run app. See [allowed usage](#allowed-usage).
4. Use [endpoints](#endpoints).

### Docker
1. Install docker engine. See [official docs](https://docs.docker.com/engine/install/).
2. Build image. Run `docker build -t url_shortener:latest .` or use existing image at [Dockerhub](https://hub.docker.com/).
3. Run app. See [allowed usage](#allowed-usage).
4. Use [endpoints](#endpoints).

### Allowed Usage
 - Run without flags. The data for urls will be stored directly in memory (Not recommended for production use)
 - Run with `-d` flag and set `PATH_TO_DB` enviroment variable. The data for urls will be stored in PostgreSQL database accessible by reference from the environment variable.

## Endpoints
| Request type                         | Endpoint | Parameters type | Parameters            | Returns       |
|--------------------------------------|----------|-----------------|-----------------------|---------------|
|<span style="color:yellow">POST</span>| /        | body(json)      | url - link to resource| short url hash|
|<span style="color:green">GET</span>  | /${hash} | query string    | short url hash        | original url  |
|<span style="color:green">GET</span>  | /docs    | -               | -                     | swagger docs  |

## Swagger

Application also provides swagger documentation accessible with endpoint `/docs`. See [endpoints](#endpoints).

## Debug Build

> [!CAUTION]
>  **Do not use debug build on production!**

For build with tag `debug` application also serves endpoint `/debug/pprof` with metrics for pprof profiler.

## How It Works

When you send the URL of an external resource, it is validated and hashed using a random number and base62 encoding.
If such a hash already exists, a reasonable number of attempts are made to recreate it.

> [!NOTE]
> It is not recommended to use in-memory URLs storage, as there is no policy for deleting outdated URLs.
> Please use an external database with a configured deletion policy to avoid running out of storage resources.

After receiving the short name, you can make a GET request to the server and get the full URL of the original resource.

## Roadmap
- [x] Cover the code with tests
- [ ] Make Docker-image publicly available
- [ ] In the future, it is planned to add the ability to automatically redirect to the initial resource instead of returning it in the response body

## Issues

If you notice any bug or vulnerability, feel free to create an Issue with a full description of the problem found.
In addition, it will be great if you offer your solution to the problem through a Pull Request.