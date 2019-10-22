# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2019-10-22
### Changed
- Moved to go modules.
- DNS check now uses the Cloudflare resolver (1.1.1.1) directly instead of relying on the local resolver.
- Unless specified otherwise, the dns checker is used by default.

## [1.0.1] - 2018-06-11
### Added
- Added a timeout to the HTTP check

### Changed
- Changed user agent for HTTP checks

## [1.0.0] - 2018-03-28
- Initial Release
