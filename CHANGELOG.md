# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.4.1 - 2021-08-11
### Changed
* Changed module path from ```github.com/strongishllama/xlambda``` to ```github.com/gofor-little/xlambda```.

## v0.4.1 - 2021-08-07
### Changed
* Fixed a bug with ```NewProxyResponse``` and ```AccessControlAllowOrigin```.

## v0.4.0 - 2021-08-06
### Added
* Added ```UnmarshalDynamoDBEventAttributeValues``` function.

## v0.3.0 - 2021-07-31
### Added
* Added ```UnmarshalAndValidate``` function.
* Added ```Validatable``` interface.
* Added a changelog.
* Added a code of conduct.

### Changed
* Updated github.com/gofor-little/env from ```v0.4.4``` to ```v1.0.0```.
* Updated github.com/gofor-little/log from ```v0.3.6``` to ```v1.0.0```.


## v0.2.1 - 2021-06-28
### Changed
* Updated github.com/aws/aws-lambda-go from ```v1.24.0``` to ```v1.26.0```.

## v0.2.0 - 2021-06-22
### Added
* Added ```NewProxyRequest``` function.

## v0.1.0 - 2021-06-22
### Added
* Added ```NewProxyResponse``` function.
* Added ```ContentType``` string type alias.
* Added ```ContentTypeApplicationJSON``` and ```ContentTypeTextHTML``` content types.
