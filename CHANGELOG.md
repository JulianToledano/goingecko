<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry is required to include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"API Breaking" for breaking API function signatures.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## Unreleased

### Features

* [47](https://github.com/JulianToledano/goingecko/pull/47) Adds a rate limited client with configurable request limits and exponential backoff retry policy.

## [v3.0.3](https://github.com/JulianToledano/goingecko/releases/tag/v3.0.3) - 2025-04-18

### Improvements

* [#44](https://github.com/JulianToledano/goingecko/pull/44) Return API errors.

## [v3.0.2](https://github.com/JulianToledano/goingecko/releases/tag/v3.0.2) - 2025-03-01

### Bug Fixes

* [#42](https://github.com/JulianToledano/goingecko/pull/42) Update `categories` endpoint.

## [v3.0.1](https://github.com/JulianToledano/goingecko/releases/tag/v3.0.1) - 2025-02-07

### Bug Fixes

* [#39](https://github.com/JulianToledano/goingecko/pull/39) Change CodeStats types from int16 to int64.
* [#41](https://github.com/JulianToledano/goingecko/pull/41) Use int64 and float64 in all types.

## [v3.0.0](https://github.com/JulianToledano/goingecko/releases/tag/v3.0.0) - 2025.01-25

### Improvements

* [#31](https://github.com/JulianToledano/goingecko/pull/31) [32](https://github.com/JulianToledano/goingecko/pull/32) Updated api functions signature to use options for optional api arguments.
* [#33](https://github.com/JulianToledano/goingecko/pull/33) Update module path to v3.

### API Breaking

* [#31](https://github.com/JulianToledano/goingecko/pull/31) [32](https://github.com/JulianToledano/goingecko/pull/32) Updated api functions signature to use options for optional api arguments.
