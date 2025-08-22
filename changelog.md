# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



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