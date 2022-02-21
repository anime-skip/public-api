# Timestamps Service

The primary backend for Anime Skip, containing accounts, auth, episodes, shows... and of course **timestamps**!

Check out the [API Playground](http://test.api.anime-skip.com/graphiql) to get started and read the docs.

## Version 2

This branch contains the refactored backend that is easier to maintain and uses dependency injection to decouple the code.

There are minor changes to the GraphQL schema, none of which should introduce breaking changes!

### Todo

- [x] ~~_Upgrade Go to 1.17_~~
  > I'm about a month early to upgrade to 1.18, but 1.17 is way newer than 1.14 like it was running before
- [x] ~~_Upgrade Gqlgen_~~
  > Library version was a few years old, so it needed upgraded. New features like [changesets](https://gqlgen.com/reference/changesets/) are nice to have as well
- [x] ~~_Dependency injection_~~
  > The main goal is to decouple the authentication logic so it can be hot-swapped when new authentication API is setup, but caching and data loaders should be much easier to include now
  >
  > This will also make testing easier when I'm ready to backfill the tests
- [x] ~~_Generate SQL methods_~~
  > Before, SQL logic was handwritten and copied from other functions using GORM.
  >
  > Instead, generate the methods so they are the exact same, removing possible human/copy/paste error. Also prevents me from forgetting to update the old methods when a bug is found
- [x] ~~_Drop GORM_~~
  > I'm, dropping GORM because I didn't really use it, and generating the SQL methods effectively implements most of the utils gorm provides, even unscoped requests for soft deleted data
  >
  > Instead, I'm using [sqlx](http://jmoiron.github.io/sqlx/) to help with the only thing generating code doesn't help with - scanning queries into structs
- [ ] Implement authorizer that matches v1 authorization
- [ ] Re-implement resolvers (delete as completed)
  - `mutationResolver.CreateAccount`
  - `mutationResolver.ChangePassword`
  - `mutationResolver.ResendVerificationEmail`
  - `mutationResolver.VerifyEmailAddress`
  - `mutationResolver.RequestPasswordReset`
  - `mutationResolver.ResetPassword`
  - `mutationResolver.DeleteAccountRequest`
  - `mutationResolver.DeleteAccount`
  - `queryResolver.Login`
  - `queryResolver.LoginRefresh`
  - `accountResolver.AdminOfShows`
  - `mutationResolver.CreateEpisodeURL`
  - `mutationResolver.DeleteEpisodeURL`
  - `mutationResolver.UpdateEpisodeURL`
  - `queryResolver.FindEpisodeURL`
  - `queryResolver.FindEpisodeUrlsByEpisodeID`
  - `episodeUrlResolver.Episode`
  - `mutationResolver.CreateEpisode`
  - `mutationResolver.UpdateEpisode`
  - `mutationResolver.DeleteEpisode`
  - `queryResolver.RecentlyAddedEpisodes`
  - `queryResolver.FindEpisode`
  - `queryResolver.FindEpisodesByShowID`
  - `queryResolver.SearchEpisodes`
  - `queryResolver.FindEpisodeByName`
  - `episodeResolver.Show`
  - `episodeResolver.Timestamps`
  - `episodeResolver.Urls`
  - `episodeResolver.Template`
  - `mutationResolver.CreateShowAdmin`
  - `mutationResolver.DeleteShowAdmin`
  - `queryResolver.FindShowAdmin`
  - `queryResolver.FindShowAdminsByShowID`
  - `queryResolver.FindShowAdminsByUserID`
  - `showAdminResolver.Show`
  - `mutationResolver.CreateShow`
  - `mutationResolver.UpdateShow`
  - `mutationResolver.DeleteShow`
  - `queryResolver.FindShow`
  - `queryResolver.SearchShows`
  - `showResolver.Admins`
  - `showResolver.Episodes`
  - `showResolver.Templates`
  - `showResolver.SeasonCount`
  - `showResolver.EpisodeCount`
  - `mutationResolver.AddTimestampToTemplate`
  - `mutationResolver.RemoveTimestampFromTemplate`
  - `templateTimestampResolver.Template`
  - `templateTimestampResolver.Timestamp`
  - `mutationResolver.CreateTemplate`
  - `mutationResolver.UpdateTemplate`
  - `mutationResolver.DeleteTemplate`
  - `queryResolver.FindTemplate`
  - `queryResolver.FindTemplatesByShowID`
  - `queryResolver.FindTemplateByDetails`
  - `templareResolver.Show`
  - `templareResolver.SourceEpisode`
  - `templareResolver.Timestamps`
  - `templareResolver.TimestampIds`
  - `thirdPartyEpisodeResolver.Timestamps`
  - `thirdPartyEpisodeResolver.Show`
  - `thirdPartyTimestampResolver.Type`
  - `mutationResolver.CreateTimestampType`
  - `mutationResolver.UpdateTimestampType`
  - `mutationResolver.DeleteTimestampType`
  - `queryResolver.FindTimestampType`
  - `queryResolver.AllTimestampTypes`
  - `mutationResolver.CreateTimestamp`
  - `mutationResolver.UpdateTimestamp`
  - `mutationResolver.DeleteTimestamp`
  - `mutationResolver.UpdateTimestamps`
  - `queryResolver.FindTimestamp`
  - `queryResolver.FindTimestampsByEpisodeID`
  - `timestampResolver.Type`
  - `timestampResolver.Episode`
  - `userResolver.AdminOfShows`

### Todo Before Publishing

- Added a migration, #20 so bump to that env var
- Env variable names are different (`ctrl+shift+F`, "os.Getenv")

## Development

The project is written in Go. However, all builds are done in a docker container, so you only need Go installed for editor support.

### Install Tooling

- [`docker`](https://docs.docker.com/get-docker/)
- [`docker-compose`](https://docs.docker.com/compose/install/)
- GNU `make`
- [`go v17+`](https://golang.org/doc/install#download)

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
