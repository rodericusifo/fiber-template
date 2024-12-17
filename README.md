# FIBER TEMPLATE

Application which can ...

## INSTALLATION

- Step 1: Install Make.

  - on Windows:

    - Install Chocolatey. Docs [Link](https://chocolatey.org/install)
    - Install Make using Chocolatey. Docs [Link](https://community.chocolatey.org/packages/make)

  - on Linux or MacOS:

    - Install Brew. Docs [Link](https://brew.sh/)
    - Install Make using Brew. Docs [Link](https://formulae.brew.sh/formula/make)

- Step 2: Install Docker. Docs [Link](https://docs.docker.com/get-docker/)

- Step 3: Copy `(env).application.env.example` to `(env).application.env` in `environment` folder

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

...
# DATABASE CACHE
...
DATABASE_CACHE_REDIS_USERNAME=""
DATABASE_CACHE_REDIS_PASSWORD=""
DATABASE_CACHE_REDIS_DATABASE=
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

...
# DATABASE CACHE
...
DATABASE_CACHE_REDIS_USERNAME="db_cache_username_example"
DATABASE_CACHE_REDIS_PASSWORD="db_cache_password_example"
DATABASE_CACHE_REDIS_DATABASE=0
...
```

## RUN APPLICATION

Run this command to start

```bash
make start ENV=(env)
```

Run this command to stop

```bash
make stop ENV=(env)
```

## DOCUMENTATION

- Postman

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10344918-27d85a45-7c41-4b84-9a31-9af5ab1e7a87%26entityType%3Dcollection%26workspaceId%3D667868fa-663b-45d5-a9ec-252ff52cb9c8#?env%5B%5BDOCKER%5D%20Go%20Fiber%20App%20Template%20Env%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDo4MDgxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6ODA4MSIsInNlc3Npb25JbmRleCI6MH1d)

## TEST COVERAGE

- Step 1: Install Mockery

```bash
go install github.com/vektra/mockery/v2@v2.20.0
```

- Step 2: Run this command

```bash
make test-cover
```
