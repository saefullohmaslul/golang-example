# Golang Example

![lint](https://github.com/saefullohmaslul/Golang-Example/workflows/lint/badge.svg?branch=master) ![test](https://github.com/saefullohmaslul/Golang-Example/workflows/test/badge.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/github/saefullohmaslul/Golang-Example/badge.svg?branch=master)](https://coveralls.io/github/saefullohmaslul/Golang-Example?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/saefullohmaslul/Golang-Example)](https://goreportcard.com/report/github.com/saefullohmaslul/Golang-Example) ![build](https://github.com/saefullohmaslul/Golang-Example/workflows/build/badge.svg?branch=master)

Simple REST Api with golang

## How to run

There are two ways to run this application, with docker or without docker

```bash
# running with docker

# running in development mode, you can use live-reload when safe file
docker-compose up development

# running in production image
docker-compose up -d production
docker-compose logs --tail=100 -f production # monitoring production container
docker-compose exec production sh # access bash on production container
```

```bash
# running in local/without docker
make install
make run
```

## Run tests

```bash
make test
```

## Stay in touch

* Author - [Saefulloh Maslul](https://linkedin.com/saefullohmaslul)

## License

Golang-Example is [MIT](LICENSE).
