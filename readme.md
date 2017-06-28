# datatogether json api

### Running locally

Ok, this is a work-in-progress, but if you have docker & docker-compose installed you should be able to clone this repo & run:
`docker-compose build && docker-compose up` from the project directory, and prove it works by visiting `http://localhost:3200/v0/urls` from a browser, where you should get a JSON response of urls.

Right now modifying & updating code is a huge pain, but this is at least a start.

### Generating Documentation

1. Install [spectacle](https://github.com/sourcey/spectacle)
2. Dev with `spectacle -d open_api.yaml`, editing `open_api.yaml` to make dem changes
3. Generate Static docs with `spectacle open_api.yaml`
4. Commit. Rinse. Repeat.