# Golang Example Boilerplate

[![lint](https://github.com/saefullohmaslul/golang-example/workflows/lint/badge.svg?branch=master)](https://github.com/saefullohmaslul/golang-example/actions?query=workflow%3Alint) [![test](https://github.com/saefullohmaslul/golang-example/workflows/test/badge.svg?branch=master)](https://github.com/saefullohmaslul/golang-example/actions?query=workflow%3Atest) [![Coverage Status](https://coveralls.io/repos/github/saefullohmaslul/golang-example/badge.svg?branch=master)](https://coveralls.io/github/saefullohmaslul/golang-example?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/saefullohmaslul/golang-example)](https://goreportcard.com/report/github.com/saefullohmaslul/golang-example) [![build](https://github.com/saefullohmaslul/golang-example/workflows/build/badge.svg?branch=master)](https://github.com/saefullohmaslul/golang-example/actions?query=workflow%3Abuild)

Simple REST API with golang (**golang boilerplate**)

## How to run

There are two ways to run this application, with docker or without docker

```bash
# running with docker

# copy .env
cp .env.example .env

# running in development mode, you can use live-reload when safe file
make docker-dev

# running in production image
make docker-prod
docker logs --tail=100 -f golang_example_production # monitoring production container
docker exec -it golang_example_production sh # access bash on production container
```

```bash
# running in local/without docker

# copy .env
cp .env.example .env
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
