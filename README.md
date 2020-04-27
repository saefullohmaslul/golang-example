# Golang Example

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
