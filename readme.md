# API

<!-- Repo Badges for: Github Project, Slack, License-->

[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](./LICENSE)
[![Codecov](https://img.shields.io/codecov/c/github/datatogether/api.svg?style=flat-square)](https://codecov.io/gh/datatogether/api)


[1-3 sentence description of repository contents]

## License & Copyright

Copyright (C) <year> Data Together
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

## Usage

If you have docker & [docker-compose](https://docs.docker.com/compose/install/), clone this repo & run the following:
```shell
  # first, cd into the project directory, then run
  docker-compose up

  # this should respond with json, having an empty "data" array
  http://localhost:8080/tasks

  # this should respond with json, with meta.message : "task successfully enqueud"
  http://localhost:8080/ipfs/add?url=https://i.redd.it/5kwih5n5i58z.jpg

  # requesting this again should now show a task in the data array, including a "succeeded" timestamp:
  http://localhost:8080/tasks

  # congrats, you've put this url of a jpeg on ipfs: https://i.redd.it/5kwih5n5i58z.jpg
  # view it here:
  https://ipfs.io/ipfs/QmeDchVWNVxFcEvnbtBbio88UwrHSCqEAXpcj2gX3aufvv

  # connect to your ipfs server here:
  # click the "files" tab, and you'll see this hash: QmeDchVWNVxFcEvnbtBbio88UwrHSCqEAXpcj2gX3aufvv
  # this means you have a local ipfs node serving the image we just processed
  https://localhost:5001/webui
```

## Development

[Step-by-step instructions about how to set up a local dev environment and any dependencies]

## Deployment

[Optional section with deployment instructions]