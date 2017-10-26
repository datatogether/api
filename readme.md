# Data Together JSON API

[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](./LICENSE)
[![Codecov](https://img.shields.io/codecov/c/github/datatogether/api.svg?style=flat-square)](https://codecov.io/gh/datatogether/api)
[![CircleCI](https://img.shields.io/circleci/project/github/datatogether/api.svg?style=flat-square)](https://circleci.com/gh/datatogether/api)


Data Together API is a JSON api for all of Data Together's centralized services. Full documentation of the API is on the server itself by running the server & visiting the `/docs/` endpoint. More instructions in the Usage section below.

## License & Copyright

Copyright (C) 2017 Data Together
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, version 3.0.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

See the [`LICENSE`](./LICENSE) file for details.

## Getting Involved

We would love involvement from more people! If you notice any errors or would like to submit changes, please see our [Contributing Guidelines](./CONTRIBUTING.md). 

We use GitHub issues for [tracking bugs and feature requests](./issues) and Pull Requests (PRs) for [submitting changes](./pulls)

## Usage

If you have docker & [docker-compose](https://docs.docker.com/compose/install/), clone this repo & run the following:
```shell
  # first, cd into the project directory, then run
  docker-compose up

  # this should take you to the api documentation:
  http://localhost:8080/docs/
```

## API Documentation

The Data Together API is documented using the [Open API Specification](https://github.com/OAI/OpenAPI-Specification), 

#### Generating Documentation

All API documentation is generated from the `open_api.yaml` file, to update docs:

1. Install [spectacle](https://github.com/sourcey/spectacle)
2. Dev with `spectacle -d open_api.yaml`, editing `open_api.yaml` to make changes
3. Generate Static docs with `spectacle open_api.yaml`
4. Commit. Rinse. Repeat.

## Development

Coming Soon!
