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

- `Added` sprig dependency v3 which bring new sprig functions (durationRound, numerous, toRawJson, htpasswd, duration, seq, randInt, fromJson, mustFromJson, bcrypt, randBytes, dig, regexQuoteMeta, osBase, osDir, osExt, osClean, osIsAbs, and, all, addf, add1f, subf, divf, mulf, maxf, and minf, chunk, and more...) and improve others (get)

## [1.3.0]

- `Added` ff1 mask to meet the requirement of re-identification from a single secret key.
- `Added` pipe mask to handle jsons with complex structure (nested arrays of objects)
- `Fixed` use same random mask on different path produce the same value.
- `Fixed` masking values in nested arrays generate panic error.

## [1.2.1]

- `Added` First public version released
