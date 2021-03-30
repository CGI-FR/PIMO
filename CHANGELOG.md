# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- `Changed` improve masking.yml readability and conciceness (possible solutions: merge frequent mask combination into single use e.g. add with template=>add with template syntax support ; include external yaml to deduplcate, parameterized masks, ...).

## [1.3.0 (Unreleased)]

- `Added` ff1 mask to meet the requirement of re-identification from a single secret key.
- `Fixed` use same random mask on different path produce the same value.
- `Fixed` masking values in nested arrays generate panic error.

## [1.2.1]

- `Added` First public version released
