# Changelog

## `[Unreleased]`

### Breaking Changes

- None

### Enhancements

- Added durations and timestamp offsets for episodes and episode urls (a8dea5cd53888e737f69d989a090680fa9f61332, eb5efc66b37fbadbcad8b4f3582fe89bf809d37c)
- Cleanup logging for development and production (fdcd5ea4aa1504781554fab9275d80baebd36742)
- Rollback migration support via environment variable (d3c61fb70fab5e8f1ee317e1efa8ded58c4c8bf6)
- Use `DATABASE_URL` instead of separate host, user, password, etc when connecting to the DB (377df0b9f4faccafb86791fc49adee1ff49e88ce)

### Fixes

- Added rollback strategy to some old migrations (1f343f9f766eaabe13b96cbd61e0540e947ee14f)

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
