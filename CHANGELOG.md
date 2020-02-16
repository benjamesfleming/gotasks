# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.1]
### Fixed
- Changed ID data type to `vaarchar(255)` for mysql compatibility

## [0.2.0]
### Added
- [Gravatar](https://en.gravatar.com/) support for profile fallback image
- [rikschennink/fitty](https://github.com/rikschennink/fitty) to fit text to containers
- This CHANGELOG to keep track of changes between versions
- WarningPanel component to deduplicate styling

### Changed
- Get email from auth provider
- Only init Sortable.js when needed (for performance)
- Restyle application using a mobile first philosophy

### Fixed
- Don't use a cachebuster when in production

## [0.1.0] - 2019-11-22
This is the initial release with very basic functionality.
I suggest following the install guide in the README to get started.

[Unreleased]: https://github.com/benjamesfleming/gotasks/compare/v0.2.1...HEAD
[0.2.1]: https://github.com/benjamesfleming/gotasks/releases/tag/v0.2.1
[0.2.0]: https://github.com/benjamesfleming/gotasks/releases/tag/v0.2.0
[0.1.0]: https://github.com/benjamesfleming/gotasks/releases/tag/v0.1.0