# FIBER TEMPLATE

Application which can ...

## INSTALLATION FOR DOCKER

- Step 1: Install Docker. Docs [Link](https://docs.docker.com/get-docker/)

- Step 2: Copy `docker.application.env.example` to `docker.application.env` in `environment` folder

`INSTRUCTION` >> fill the empty values, from this:

```bash
# docker.application.env
...
# DATABASE SQL
...
DATABASE_MYSQL_NAME=""
DATABASE_MYSQL_USERNAME=""
DATABASE_MYSQL_PASSWORD=""
...
```

to this:

```bash
# docker.application.env
...
# DATABASE SQL
...
DATABASE_MYSQL_NAME="db_name_example"
DATABASE_MYSQL_USERNAME="db_username_example"
DATABASE_MYSQL_PASSWORD="db_password_example"
...
```

## INSTALLATION FOR LOCAL

- Step 1: Install Golang. Docs [Link](https://go.dev/doc/install)

- Step 2: Run These Commands

```bash
$ go mod download
$ go mod tidy
$ go mod verify
```

- Step 3: Install Wire

```bash
$ go install github.com/google/wire/cmd/wire@latest
```

- Step 4: Copy `(env).application.env.example` to `(env).application.env` in `environment` folder

`INSTRUCTION` >> fill the empty values, from this:

```bash
# (env).application.env
...
# DATABASE SQL
...
DATABASE_MYSQL_NAME=""
DATABASE_MYSQL_USERNAME=""
DATABASE_MYSQL_PASSWORD=""
...
```

to this:

```bash
# (env).application.env
...
# DATABASE SQL
...
DATABASE_MYSQL_NAME="db_name_example"
DATABASE_MYSQL_USERNAME="db_username_example"
DATABASE_MYSQL_PASSWORD="db_password_example"
...
```

## RUN APPLICATION ON DOCKER

Run this command to start

```bash
$ make docker-start
```

Run this command to stop

```bash
$ make docker-stop
```

## RUN APPLICATION ON LOCAL

Run this command to start

```bash
# on environment (development)
$ make start-dev
```

Run this command to stop

<kbd>Ctrl</kbd> + <kbd>C</kbd>

## DOCUMENTATION

- Postman (Docker Environment)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87%26entityType%3Dcollection%26workspaceId%3D667868fa-663b-45d5-a9ec-252ff52cb9c8#?env%5B%5BDOCKER%5D%20Go%20Fiber%20App%20Template%20Env%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDo4MDgxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6ODA4MSIsInNlc3Npb25JbmRleCI6MH1d)

- Postman (Local Environment)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87%26entityType%3Dcollection%26workspaceId%3D667868fa-663b-45d5-a9ec-252ff52cb9c8#?env%5B%5BLOCAL%5D%20Go%20Fiber%20App%20Template%20Env%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDo4MDgwIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6ODA4MCIsInNlc3Npb25JbmRleCI6MH1d)

## TEST COVERAGE

- Step 1: Install Mockery

```bash
$ go install github.com/vektra/mockery/v2@v2.20.0
```

- Step 2: Run this command

```bash
$ make test-cover
```
