# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [1.4.0] - 2025-09-03

### Added

- A simple check in the cache for when a dropped command fails silently because the collection isn't found.

### Changed

- The Yacc grammaer to allow for incoming values as JSON (either quoter or not).
- The lexical analyzer (lexer) that will parse the incoming token for JSON values, now more stable and robust.

### Removed

- The old Lexer system that used to parse word-by-word for operations (like SET/GET).


## [1.3.0] - 2025-08-30

### Added

- A WAL mechanism.
- A configuration mechanism that takes in place a TOML file (`simplecache.toml` or what specified) with configuration for the WAL.
- A storing and retrieving mechanism for WAL.
- A recovery mechanism that reads from the WAL and reapplies the information to the cache.

### Changed

- `pkg/tcp/connection.go` to host the new functionality for WAL.

### Removed

- The whole command system with [Cobra](github.com/spf13/cobra) and migrated to a syscall/exec system that will host systemd execution (with SIGTERM).
- An old attempt to telemetry.
- Unused packages.


## [1.2.1] - 2025-08-24

### Added

- Mutexes and lockers for operations on the cache that should be locked and force mutual exclusion (like SETting a key).

### Changed

- `pkg/scl/tokenizer.go` was chaned to host the related flow for parsing incoming statement and making them globally available.


## [1.2.0] - 2025-08-23

### Added

- The initial structure of the Btree, which is what will hold the data for our caching mechanism.

### Changed

- Commands and move them to the root structure of `cmd`.
- Naming and better functions related to starting/stopping the server.


## [1.1.2] - 2025-08-21

### Added

- TTI to system (parsed) to manage Invalidation Time.

### Changed

- Version + improved Readme and changelogs.


## [1.1.1] - 2025-08-21

### Added

- Compiler via `Yacc`.
- Added Lexer and Parser to analyse structure.
- `Tokenzer` to manage incoming data to then parse from Parsers and manage from Lexer.

### Removed

- An attempt to a WAL system.


## [1.0.1] - 2025-08-16

### Added

- Added the CLI Structure to connect to the cache via SSH.
- Added the Plain connection mechanism.


## [1.0.0] - 2025-08-15

### Added

- Project setup via `go mod init`.
- Added both `readme` and `changelogs`.
- Added a versioning JSON with the version info.
- Added the cobra CLI strucutre.