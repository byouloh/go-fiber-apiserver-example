# go-fiber-apiserver-example
Go fiber webframework을 사용하여 web api server를 만드는 예제

참조할 글 : https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j

이 예제는 [create-go-app/cli](https://github.com/create-go-app/cli)로 만들어진 fiber web framework을 이용한 api server 예제다.
create-go-app 명령어를 이용하여 웹 api server를 만들어 실행해보는 과정에서 발생한 각종 오류를 기록해두고자 글을 쓰게 되었다.

Go를 선택한 이유는
- Tensorflow의 학습시킨 모델을 파이썬과 GO 형식으로 가져오기할 수 있다. 따라서, GO를 다룰 수 있으면 학습시킨 모델을 가져와서 데이타의 딥러닝 처리가 가능하다.
- 빠르다 : GO는 컴파일 언어로 속도가 스크립트 언어들보다 매우 빠르다.
- 서버 작업에 최적화된 언어 : 서버를 위한 언어라고 보면 된다.

Fiber web framework을 선택한 이유는
- 빠르다
- node.js의 express.js와 같은 형태로 사용하도록 만들어져서 node를 다룬 사람이면 쉽다
- go 언어를 먼저 학습하지 말고 바로 fiber를 공부하면서 go를 학습해도 된다(이건 fiber만 해당하는 사항은 아님)

이상의 이유로 fiber api server를 살펴보다가 우연히 github.com/create-go-app/cli 툴을 알게되었다.
이 툴을 사용할 때 다음과 같은 문제가 있었다(바로 backend/README.md 파일로 넘어가도 된다).

1. cgapp을 설치하기 위한 절차를 따름 : go 설치, go install github.com/create-go-app/cli@latest
2. 설치 후 cgapp create 명령으로 예제를 생성하는 경우 frontend가 생성되지 않는다면, npm이 7.x 버전 이상인지 확인할 것 :
``` bash
npm install -g npm@latest
```
3. go.mod 파일에서 모쥴 이름이 github.com/create-go-app/xxx와 같은 형태로 되어 있으니, 이와 관계된 항목들은 자신이 원하는 모쥴명으로 바꾼다(현재 게시된 예제에 바뀌어져있다: example.com/fiber-apiserver로 바꾸어서 썼다). 그리고, 모쥴명이 필요한 다른 파일들도 모두 바꾸어야한다.
4. .gitignore 파일에서 **/app/을 주석처리하는게 좋다(app을 git에 올리려면).
5. 이제 backend/README.md 파일로 넘어가자 : 도커로 돌리거나, 또는 그냥 실행시키기 위해서는 도커와 추가적으로 3개 이상의 모쥴을 설치해야만 한다. 또한, postgres driver와 migrate 문제도 있으니 정상적으로 실행하려면 꼭 읽어봐야한다. 문제가 되는 부분만 아래에 붙여놓겠다.

## backend/app 실행시 문제
1. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:
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

2. Run project by this command:

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
