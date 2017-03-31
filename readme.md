# archivers-api
### An Archivers 2.0 joint

archivers-api is a json api for archivers 2.0 [proposed services](https://github.com/edgi-govdata-archiving/proposed-services). you can hit it at [api.archivers.space](https://api.archivers.space). It's written in [go](https://golang.org).

**Documentation is at [docs.archivers.space](https://docs.archivers.space)**

### Generating Documentation

1. Install [spectacle](https://github.com/sourcey/spectacle)
2. Dev with `spectacle -d open_api.yaml`, editing `open_api.yaml` to make dem changes
3. Generate Static docs with `spectacle open_api.yaml`
4. Commit. Rinse. Repeat.