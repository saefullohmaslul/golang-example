<p align="center">
  <a href="https://github.com/saefullohmaslul/golang-example" target="blank"><img src="https://raw.githubusercontent.com/nielsing/yar/master/images/yargopher3.png" width="200" alt="Go" /></a>
</p>

<h1 align="center">Golang Clean Boilerplate</h1>

## Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Feature

- Clean code
- Docker and docker-compose
- ORM with gorm
- Unit testing

## Setup

### Building and running the dockerized codebase

1. Clone the repository.

1. Build all containers using `docker-compose build --no-cache`.

1. Run all containers using `docker-compose up`.

   **Tips**:

   - Use `--build` in `docker-compose` to force update the docker image created, e.g. `docker-compose up --build`

1. The docker and docker compose will setup all requirements on the fly, and please provide coffee as this may take a few minutes.

1. If have finished it, that means you can use this server with all databases and data ready to use.
### Cleaning up

1. When you're done, `Ctrl-C` in the main `docker-compose up` window to terminate the running processes.

1. Run `docker-compose down` to stop and remove containers.

## Develop By

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/saefullohmaslul"><img src="https://avatars.githubusercontent.com/u/20754023" width="100px;" alt=""/><br /><sub><b>Saefulloh Maslul</b></sub></a></td>
  </tr>
</table>
<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

Copyright (c) 2021 Saefulloh Maslul
All rights reserved.