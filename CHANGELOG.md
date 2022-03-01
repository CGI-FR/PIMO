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

## [1.12.0]

- `Added` possibility to pass multiple masking config to the command line
- `Added` markov mask to generate pseudo text based on a sample text
- `Added` command to export maskings as a mermaid flow chart
- `Added` possibility to use option `preserve: notInCache` with mask `fromCache`
- `Added` flag to mask input while a declared condition is met
- `Added` flag to mask input until a declared condition is met
- `Added` flag to declare a simple mask definition in command line
- `Added` possibility to declare multiple selectors for a masking definition
- `Fixed` pimo doesn't panic anymore with fromjson combined with pipe mask, or fromjson combined with another fromjson mask with nested selectors
- `Fixed` mask `replacement` with nested selectors
- `Fixed` using fromCache, the line is not deleted if the jsonpath in not present in the document

## [1.11.0]

- `Added` option preserve in masking configuration.
- `Fixed` cache with mask `fluxUri`.

## [1.10.0]

- `Added` luhn mask to generate valid checksums using the Luhn algorithm.

## [1.9.1]

- `Fixed` JSON Schema validation for `masks` property (only required if `mask` is not set).

## [1.9.0]

- `Added` possibility to use a template string with `randomChoiceInUri` mask
- `Added` new mask `add-transient` same as `add` but the field is not exported in the jsonline output
- `Added` possibility to use a template string directly with `add` mask
- `Added` possibility to chain multiple masks in YAML configuration on the same jsonpath with the `masks` property

## [1.8.0]

- `Added` new mask `fromjson` to convert a string to object model

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
