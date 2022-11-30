# gokomodo-challenge
Full API Doc: [DOCUMENTATION](https://documenter.getpostman.com/view/6786432/2s8Yt1qovX) 📖

## What's going on here?

- [Stack used 🥞](#stack-used-)
- [Before Run ☕️](#before-run-%EF%B8%8F)
    * [Clone Repo 💾](#clone-repo-)
- [How to Run 👟](#how-to-run-)
    * [Run Locally 🏃](#run-locally-)
    * [Debugging 🕵️‍♂️](#debugging-%EF%B8%8F%EF%B8%8F)
    * [Dockering 🐳](#dockering-)
- [How to Use 💻](#how-to-use-)

## Stack used 🥞

- SQLite3
- Golang
- Echo Framework

## Before Run ☕️

❗️ This app needed `gcc` to run since sqlite3 library required that. Here is the article about that [here](https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image).

### Clone Repo 💾

Clone this repo using `Download` button or bash 👨‍💻

```bash
$ git clone https://github.com/hrz8/gokomodo-challenge.git
```

## How to Run 👟

### Run Locally 🏃

This command below will compiling the `main.go` and all used packaged into binary file at the first and run the binary right after that.

```bash
$ make run
```

It will running in the localhost with the `3000` port given `http://localhost:3000`.

### Debugging 🕵️‍♂️

Debug app using VsCode Debugger Tool

- Select your debugger to be set as `Launch server`
- Press `F5` to run the debugging
- Done!

### Dockering 🐳

- Image Builds (Example)

```bash
# build as docker image
$ make docker-build
# make sure docker image already registered
$ docker image ls
# start the container
# you will automatically see the logs of the app as well
$ make docker-start
# check endpoint
$ curl --location --request POST 'http://localhost:3000/seller/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "your name",
        "email": "your@email.com",
        "password": "password",
        "address": "your address"
    }'
```

- Docker Compose (Example)

```bash
# you will automatically see the logs of the app as well
$ make docker-compose
# check endpoint
$ curl --location --request POST 'http://localhost:3000/seller/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "your name",
        "email": "your@email.com",
        "password": "password",
        "address": "your address"
    }'
```

## How to Use 💻

This application based on RESTful API, so the usage will required to do the HTTP request to each available endpoint below. Full Documentation of how to use each endpoint is in the link right here: [DOCUMENTATION](https://documenter.getpostman.com/view/6786432/2s8Yt1qovX) 📖

## Test 🧪

Test unit.

```bash
$ make test
```

Test lint.

```bash
$ make lint
```

## Author ℹ️

Hirzi Nurfakhrian
