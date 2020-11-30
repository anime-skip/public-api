# Changelog

## `[Unreleased]`

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
