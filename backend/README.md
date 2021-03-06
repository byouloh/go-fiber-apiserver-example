# Fiber backend template for [Create Go App CLI](https://github.com/create-go-app/cli)

<img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/fiber-go-template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" />

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## ⚡️ Quick start

1. Create a new project with Fiber:

```bash
cgapp create

# Choose a backend framework:
#   net/http
# > fiber
```

2. Rename `.env.example` to `.env` and fill it with your environment values.
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:
- windows 10에서는 Docker Desktop을 설치해서 사용하는 것을 권장한다.
- [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations
- [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
- [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues

다음의 명령으로 필요한 3가지 모쥴을 설치한다.

```bash
go get github.com/golang-migrate/migrate/v4
go install github.com/swaggo/swag@latest
go install github.com/securego/gosec@latest
```

migrate, swag, securego가 제대로 설치되지 않는 경우, 해당 모쥴들이 추가된 이후에 설치된 디렉토리로 이동해서 다시 설치해야만 정상적으로 동작한다.
windows 10 커맨드 창에서,

```bash
cd C:\Users\USERID\go\pkg\mod\github.com\golang-migrate\migrate\v4@v4.14.1\cmd\migrate
go install .

cd C:\Users\USERID\go\pkg\mod\github.com\securego\gosec@v0.0.0-20200401082031-e946c8c39989\cmd\gosec
go install .

cd C:\Users\USERID\go\pkg\mod\github.com\swaggo\swag@v1.7.1\cmd\swag
go install .
```

제대로 설치되었는지 다음의 명령을 통해 확인한다.

```bash
swag -v
migrate -version
gosec -version
```

4. Run project by this command:

windows 10에서 Docker Desktop을 실행하고,

docker.run 명령중 migrate.up을 실행하기전에
backend/Makefile에서 작업 디렉토리를 자신의 디렉토리로 바꿔줘야한다(바꾸지 않으면 path 없다고 에러남)

```bash
MIGRATIONS_FOLDER = Your-working-directory/go-fiber-api-server/backend/platform/migrations
```

docker.run 실행

```bash
make docker.run
```

make docker.run을 실행시키면 migrate.up이 실패한다. make migrate.up을 실행해도 마찬가지.

migrate.up에서 unknown driver 에러가 발생하면, 다음의 명령을 명령창에서 실행한다.
(참고: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain)

database/sql/driver 모쥴은 postgresql의 driver를 실제로 설치하는 것은 아니다.
따라서, 별도로 postgresql의 driver를 설치해줘야 postgresql DB를 제어할 수 있다.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
```

그런다고 해도, 다시 아래와 같이 cgapp-postgres 호스트를 찾을 수 없다는 에러가 나온다.
migrate 명령에서 port:5432를 추가했음을 주의하자(@cgapp-postgres:5432, 생성된 템플릿에는 포트가 없다).

```bash
make migrate.up
migrate -path D:/works/go/go-fiber-api-server/backend/platform/migrations -database "postgres://postgres:password@cgapp-postgres:5432/postgres?sslmode=disable" up
error: dial tcp: lookup cgapp-postgres: no such host
make: *** [migrate.up] 오류 1
```

Windows나 MacOS 환경에서는 host ip나 host.docker.internal로 postgresql db host에 접근할 수 있다.
호스트 네임으로 데이타베이스 호스트를 찾지 못하는 경우, 도커가 실행되고 있는 머신의 ip 주소, 또는 host.docker.internal를 호스트네임에 넣으면 된다.

```bash
make migrate.up
migrate -path D:/works/go/go-fiber-api-server/backend/platform/migrations -database "postgres://postgres:password@host.docker.internal:5432/postgres?sslmode=disable" up
1/u create_init_tables (20.0987ms)
```

여기까지 에러를 수정하고 나면 make docker.run 명령이 에러없이 정상적으로 실행된다.

![screenshot](https://user-images.githubusercontent.com/3069673/131583232-5d3e336c-2228-4359-9c91-c3c6c8d079e3.PNG)


5. Go to API Docs page (Swagger): [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)

## 📦 Used packages

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.18.0` | core       |
| [gofiber/jwt](https://github.com/gofiber/jwt)                         | `v2.2.7`  | middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.17.0` | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`  | tests      |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt)                   | `v4.0.0`  | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.3.0`  | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)                       | `v1.3.4`  | database   |
| [jackc/pgx](https://github.com/jackc/pgx)                             | `v4.13.0` | database   |
| [go-redis/redis](https://github.com/go-redis/redis)                   | `v8.11.3` | cache      |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.7.1`  | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`  | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.9.0` | utils      |

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/repository` folder for describe `const` of your project
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/cache` folder with in-memory cache setup functions (by default, Redis)
- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS="dev"

# Server settings:
SERVER_HOST="0.0.0.0"
SERVER_PORT=5000
SERVER_READ_TIMEOUT=60

# JWT settings:
JWT_SECRET_KEY="secret"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15
JWT_REFRESH_KEY="refresh"
JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=720

# Database settings:
DB_HOST="cgapp-postgres"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="password"
DB_NAME="postgres"
DB_SSL_MODE="disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2

# Redis settings:
REDIS_HOST="cgapp-redis"
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB_NUMBER=0
```

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
