# Public API

The primary backend for Anime Skip, containing user preferences, episodes, shows... and of course **timestamps**!

Check out the [API Playground](http://api.anime-skip.com) for example usage and read the docs.

## Contributing

The project is written in Go. However, all builds are done in a docker container, so you only need Go installed for editor support.

### Contributors

<a href="https://github.com/anime-skip/public-api/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=anime-skip/public-api" />
</a>

### Install Tooling

- [`go v18+`](https://golang.org/doc/install#download)
- [`jq`](https://golang.org/doc/install#download) to parse JSON for some run scripts
- [GNU `make`](https://www.gnu.org/software/make/) to execute the `Makefile`
- [`docker`](https://docs.docker.com/get-docker/) for building and running locally
- [`docker-compose`](https://docs.docker.com/compose/install/) for running a database when starting the app
  > Make sure you have the `docker-compose` command, not `docker compose`. Create an alias for it if needed:
  >
  > ```bash
  > alias docker-compose="docker compose"
  > ```
- [`ginkgo` CLI](https://onsi.github.io/ginkgo/#getting-started) for running BDD-style unit tests
  > ```bash
  > go install github.com/onsi/ginkgo/v2/ginkgo
  > ```

### Build Commands

```bash
make run       # Run API, postgres database, and other services locally
make run-clean # Same as run, but start with an empty postgres database
make watch     # Run everything, but restart when saving a file
make gen       # Generate GraphQL server code after changing api/*.graphqls
make compile   # Compile the application outside of docker to bin/server
make build     # Build the latest development docker image
make test      # Run unit tests
```

### Editor Setup

Feel free to add a section for your editor if it's not listed!

#### VS Code

Install the [golang extension](https://marketplace.visualstudio.com/items?itemName=golang.go). Make sure to follow the quick start to install extra CLI tooling the extension relies on!
