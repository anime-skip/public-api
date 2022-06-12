# Public API

The primary backend for Anime Skip, containing user preferences, episodes, shows... and of course **timestamps**!

Check out the [API Playground](http://api.anime-skip.com) to get started and read the docs.

## Version 2

This branch contains the refactored backend that is easier to maintain and uses dependency injection to decouple the code.

There are minor changes to the GraphQL schema, none of which should introduce breaking changes!

## Development

The project is written in Go. However, all builds are done in a docker container, so you only need Go installed for editor support.

### Install Tooling

- [`docker`](https://docs.docker.com/get-docker/)
- [`docker-compose`](https://docs.docker.com/compose/install/)
   > Make sure you have the `docker-compose` command, not `docker compose`. Create an alias for it if needed:
   >
   > ```bash
   > alias docker-compose="docker compose"
   > ```
- GNU `make` to execute the `Makefile`
- (Optional) [`go v18+`](https://golang.org/doc/install#download)

### Build Commands

```bash
make run       # Run everything (API and postgres database) locally
make watch     # Run everything, but restart when saving a file
make compile   # Compile with go outside of docker, much faster than...
make build     # Build the latest development docker image
make run-clean # Same as run, but start with a brand-new postgres database
```

### Editor Setup

Feel free to add a section for your editor if it's not listed!

#### VS Code

Install the [golang extension](https://marketplace.visualstudio.com/items?itemName=golang.go). Make sure to follow the quick start to install extra CLI tooling the extension relies on!
