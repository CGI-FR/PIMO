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

## [1.7.0]

- `Added` new mask `template-each` to mask each value of an array with a template
- `Fixed` pimo doesn't crash anymore when nulls values occurs in the middle of a path

## [1.6.2]

- `Fixed` Protect masks against null values (dateparser, duration, hash, randdura, range)

## [1.6.1]

- `Fixed` Range over array in template mask

## [1.6.0]

- `Fixed` Pipe mask with repeat flag failed in panic (#34)
- `Added` jsonschema command to generate schema of masking.yml file
- `Added` flag to enable or disable coloring in output logs (--color [yes|no|auto])

## [1.5.0]

- `Changed` order of keys on JSON objects will now be preserved on the output

## [1.4.0]

- `Added` new functions in `template` mask via sprig dependency v3 (durationRound, numerous, toRawJson, htpasswd, duration, seq, randInt, fromJson, mustFromJson, bcrypt, randBytes, dig, regexQuoteMeta, osBase, osDir, osExt, osClean, osIsAbs, and, all, addf, add1f, subf, divf, mulf, maxf, and minf, chunk, and more...) and improve others (get)
- `Added` strutured logging with `-v` and `--log-json` flags
- `Added` debug option `--debug`, warning it's slow do not use in production
- `Fixed` flag `--skip-line-on-error`
- `Fixed` flag `--skip-field-on-error`
- `Fixed` handling of null values by the the `ff1` mask

## [1.3.0]

- `Added` ff1 mask to meet the requirement of re-identification from a single secret key.
- `Added` pipe mask to handle jsons with complex structure (nested arrays of objects)
- `Fixed` use same random mask on different path produce the same value.
- `Fixed` masking values in nested arrays generate panic error.

## [1.2.1]

- `Added` First public version released
