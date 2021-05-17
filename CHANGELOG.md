## [1.2.1](https://github.com/anime-skip/backend/compare/v1.2.0...v1.2.1) (2021-05-17)


### Bug Fixes

* cache result of `recentlyAddedEpisodes` for 20 min ([6bd2e6b](https://github.com/anime-skip/backend/commit/6bd2e6ba31885b04b742d91e0babc49c2903845e))



# [1.2.0](https://github.com/anime-skip/backend/compare/v1.1.0...v1.2.0) (2021-05-16)


### Features

* added show and season templates ([2f6ccef](https://github.com/anime-skip/backend/commit/2f6ccef2f11c3538742ceb999b5fc1addc69e203))



# [1.1.0](https://github.com/anime-skip/backend/compare/v1.0.9...v1.1.0) (2021-05-11)


### Bug Fixes

* Allow for passing not the entire preferences object to `savePreferences` ([68ac868](https://github.com/anime-skip/backend/commit/68ac868dbfe2207c72e25abb59f46ab5892b1d3d))


### Features

* added new preferences around hiding the toolbar and timeline ([bd4ca99](https://github.com/anime-skip/backend/commit/bd4ca99c67759eccc961cb31907296b260385728))



## [1.0.9](https://github.com/anime-skip/backend/compare/v1.0.8...v1.0.9) (2021-02-18)


### Bug Fixes

* Consume env.DISABLE_EMAILS ([f184245](https://github.com/anime-skip/backend/commit/f184245ac93cbf4421bd60d78e72b471372e7489))
* Internal error with the fetchEpsiodeByName query ([9389729](https://github.com/anime-skip/backend/commit/938972977923127d7343a3c6d9dc519d605c380a))
* Setup login timer to remove constant sleeps ([51b1cb3](https://github.com/anime-skip/backend/commit/51b1cb35e447160f1a827c14bb1d79f096ec8df0))
* updateTimestamps, better rollback utils ([43c1e63](https://github.com/anime-skip/backend/commit/43c1e63d89ff576135a0bd5e9cd8457c2f31815c))



## [1.0.8](https://github.com/anime-skip/backend/compare/v1.0.7...v1.0.8) (2021-02-18)


### Bug Fixes

* Don't accept empty strings for episode info ([81b1a7d](https://github.com/anime-skip/backend/commit/81b1a7d9401404ef31185843db5c2c2f0690f48f))



## [1.0.7](https://github.com/anime-skip/backend/compare/v1.0.6...v1.0.7) (2021-02-18)


### Bug Fixes

* Add CleanURL function and tests for future migration ([e5709af](https://github.com/anime-skip/backend/commit/e5709af42426e2bfa9bc691538512a900ec28d39))
* Don't query deleted episodes by name ([317973b](https://github.com/anime-skip/backend/commit/317973b30434c1ee98865effd43da07ff4c03ecc))



## [1.0.6](https://github.com/anime-skip/backend/compare/v1.0.5...v1.0.6) (2021-02-18)


### Bug Fixes

* Fix introspection ([5f8bf2d](https://github.com/anime-skip/backend/commit/5f8bf2d1f71f064f34a5a0a90c914e8d34123d3a))
* Update episodes mutations to accept durations ([d2cef92](https://github.com/anime-skip/backend/commit/d2cef9229c1889e5b83025688ef74d03c11de229))



## [1.0.5](https://github.com/anime-skip/backend/compare/v1.0.4...v1.0.5) (2021-02-18)


### Bug Fixes

* Added duration fields to graphql models and postgres entieies ([1353445](https://github.com/anime-skip/backend/commit/135344582250f8d0e74cb0f260d94cd1402ea29e))
* Added timestampOffset to EpisodeUrl ([0de0eb8](https://github.com/anime-skip/backend/commit/0de0eb887a8665998eba96366db0bf5fb76303fc))
* Allow rollbacks based on an environment variable ([88914d4](https://github.com/anime-skip/backend/commit/88914d4c1cacf98c5737e4ed85fd2e554ff3ba9a))
* Cleanup logs ([ec82150](https://github.com/anime-skip/backend/commit/ec821500ed60f7828a35f93ca05716b0f954567d))
* Fix old migartions to have a rollback strategy ([a61a71a](https://github.com/anime-skip/backend/commit/a61a71a3fafcd77fea04bcaf6373b631ece7760a))
* Use DATABASE_URL for connection string ([a69dd21](https://github.com/anime-skip/backend/commit/a69dd214a5fe6a60e6c2a6254bb31ef49ac93fbf))



## [1.0.4](https://github.com/anime-skip/backend/compare/v1.0.3...v1.0.4) (2021-02-18)


### Bug Fixes

* Added show data to thrid party episodes ([e1c64fb](https://github.com/anime-skip/backend/commit/e1c64fb205e38c7a09504c5d1e8bfdc2628c8bfe))



## [1.0.3](https://github.com/anime-skip/backend/compare/v1.0.2...v1.0.3) (2021-02-18)


### Bug Fixes

* Fixed query so that only 1 episode is returned per timestamp ([a81ac39](https://github.com/anime-skip/backend/commit/a81ac39192978cdbb358e5c9ba486ebd4aa2693d))



## [1.0.2](https://github.com/anime-skip/backend/compare/v1.0.1...v1.0.2) (2021-02-18)


### Bug Fixes

* Added query for recently added epiosdes with timestamps ([ceea872](https://github.com/anime-skip/backend/commit/ceea872b000a66b8e04727a07c5937301a60ddce))



## [1.0.1](https://github.com/anime-skip/backend/compare/v1.0.0...v1.0.1) (2021-02-18)


### Bug Fixes

* Add env bypass for disabling the admin directive ([5d5138f](https://github.com/anime-skip/backend/commit/5d5138f709de8fe1bd02ed3cac077620c55d58b8))
* Removed email account allowlist ([c935eeb](https://github.com/anime-skip/backend/commit/c935eeb289a2e4933aab7f0d58a954d9939ed8ba))



# 1.0.0 (2021-02-18)




## `v1.0.8`

### Breaking Changes

- None

### Enhancements

- None

### Fixes

- Replace empty episode values with `null`

## `v1.0.7`

### Breaking Changes

- None

### Enhancements

- Docs added about cleaning URLs before using them with `EpisodeUrl`

### Fixes

- Unexpected crash when an environment variable isn't passed

## `v1.0.6`

### Breaking Changes

- None

### Enhancements

- Added `updateEpisodeUrl` mutation
- Added `baseDuration` to `ThirdPartyEpisode`

### Fixes

- None

## `v1.0.5`

### Breaking Changes

- None

### Enhancements

- Added durations and timestamp offsets for episodes and episode urls
- Cleanup logging for development and production, added info level
- Rollback migration support via environment variable
- Use `DATABASE_URL` instead of separate host, user, password, etc when connecting to the DB

### Fixes

- Added rollback strategy to some old migrations

## `v1.0.4`

### Breaking Changes

- None

### Enhancements

- Added `showId` and `show` to `ThirdPartyEpisode`, providing resolution when multiple shows have
episodes with the same name

### Fixes

- None

## `v1.0.3`

### Breaking Changes

- None

### Enhancements

- none

### Fixes

- Recently added was returning an episode for each timestamp, now it just returns one

## `v1.0.2`

### Breaking Changes

- None

### Enhancements

- Added the `recentlyAddedEpisodes` query for the website

### Fixes

- None

## `v1.0.1`

Prep for public release

### Breaking Changes

- Remove account email allowlists - anyone can create an account

### Enhancements

- ENV var for enabling the `@isShowAdmin` directive

### Fixes

- None

## `v1.0.0`

Initial release
