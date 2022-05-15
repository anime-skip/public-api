# Public API

The primary backend for Anime Skip, containing user preferences, episodes, shows... and of course **timestamps**!

Check out the [API Playground](http://test.api.anime-skip.com/graphiql) to get started and read the docs.

## Version 2

This branch contains the refactored backend that is easier to maintain and uses dependency injection to decouple the code.

There are minor changes to the GraphQL schema, none of which should introduce breaking changes!

### Todo

- [x] ~~_Upgrade Go to 1.17_~~
  > I'm about a month early to upgrade to 1.18 and generics :/
- [x] ~~_Upgrade Gqlgen_~~
  > Library version was a few years old, so it needed upgraded. New features like [changesets](https://gqlgen.com/reference/changesets/) are nice to have as well.
- [x] ~~_Dependency injection_~~
  > The main goal is to decouple the authentication logic so it can be hot-swapped when new authentication API is setup, but caching and data loaders should be much easier to implement and include now.
  >
  > This will also make testing easier when I'm ready to backfill the tests
- [x] ~~_Generate SQL methods_~~
  > Before, SQL logic was handwritten and copied from other functions using GORM.
  >
  > Instead, generate the methods so they are the exact same, removing possible human/copy/paste error. Also prevents me from forgetting to update the old methods when a bug is found.
- [x] ~~_Drop GORM_~~
  > I'm, dropping GORM because I didn't really use it, and generating the SQL methods effectively implements most of the utils gorm provides, even unscoped requests for soft deleted data.
  >
  > Instead, I'm using [sqlx](http://jmoiron.github.io/sqlx/) to help with the only thing generating code doesn't help with - scanning queries into structs.
- [x] Implement authorizer that matches v1 authorization
  > The new injectable authorizor needs to match the existing logic so no-one gets logged out
- [x] Re-implement resolvers (delete as completed)
- [x] `ctrl+shift+f` "not implemented"
- [ ] `ctrl+shift+f` "TODO"
- [ ] Track Client IDs
- [ ] Implement repo layer

### Todo Before Publishing

- Added a migration, #20 so bump to that env var
- Env variable names are different (`ctrl+shift+F`, "os.Getenv")

## Development

The project is written in Go. However, all builds are done in a docker container, so you only need Go installed for editor support.

### Install Tooling

- [`docker`](https://docs.docker.com/get-docker/)
- [`docker-compose`](https://docs.docker.com/compose/install/)
- GNU `make`
- (Optional) [`go v17+`](https://golang.org/doc/install#download)

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
