# Logging

PIMO has 6 level of logging :
- 0 (none) : log nothing at all
- 1 (error) : log only errors
- 2 (warn) : same as level 1 + warnings that should be checked by user
- 3 (info) : same as level 2 + information about what is processed
- 4 (debug) : same as level 3 + debugging information, to analyse what can cause an unexpected behavior
- 5 (trace) : same as level 4 + tracing of events in code (enter function, exit function, values...)

This can be set with the `-v` flag (long version `--verbosity`) :

```console
$ pimo --empty-input -vinfo
9:10PM INF Logger level set to info
9:10PM INF Start PIMO config=masking.yml dump-cache={} empty-input=true load-cache={} repeat=1 skipFieldOnError=false skipLineOnError=false
9:10PM INF Add mask config=masking.yml context=empty-input mask="constant Benjamin" path=name
9:10PM INF Add mask config=masking.yml context=empty-input mask="hash size=2" path=surname
9:10PM INF Add mask config=masking.yml context=empty-input mask="hash size=3" path=town
9:10PM INF Add mask config=masking.yml context=empty-input mask="randomInt min=18 max=90" path=age
9:10PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=name
9:10PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=surname
9:10PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=town
9:10PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=age
{}
9:10PM INF End PIMO config=masking.yml duration=12.3069ms input-line=1 output-line=2 return=0 stats={"ignoredPaths":4,"skippedFields":0,"skippedLines":0}
```

By default, the log level is set to error.

Logs are written on the standard error (stderr), so it will not mix with the data stream. You can still pipe the output in a file : `pimo -v5 > out.jsonl`.

Logs format is :

```
TIME LVL Message field1=value1 field2=value2 ...
```

Example :
```
2:57PM INF Start PIMO config=masking.yml dump-cache={} empty-input=true load-cache={} repeat=1 skipFieldOnError=false skipLineOnError=false
```

## Fields

Here is a list of some fields that can be logged with description.

Field | Description
--|--
input-line | Index of the line currently being read in stdin, 1-based (with `--empty-input` flag, always equals to 1)
output-line | Index of the line currently being written in stdout, 1-based
context | The current context, e.g. `stdin[5]` means the fifth line on the standard input<br>If a pipe mask is used, the current context will show the full path currently processed, e.g. `stdin[3]/persons[1]`
path | The current path as it appears in the mask definition
return | PIMO exits the process with this return code
stats | Global counters like the number of ignored fields or lines, in JSON format
duration | The total time spend in processing the pipeline

## Events

Here are some of the events that are logged (non exhaustive list).

Event | Log Level | Fields
--|--|--
PIMO starts | `INFO` | All flags used on the command line
PIMO ends successfully | `INFO` | `input-line`, `output-line`, `return`, `stats`, `duration`
PIMO ends unexpectedly | `WARN` | `input-line`, `output-line`, `return`, `duration`
A mask is applied | `DEBUG` | `input-line`, `output-line`, `context`, `path`
A jsonpath point to nothing | `WARN` | `input-line`, `output-line`, `context`, `path`
A line is skipped (`--skip-line-on-error` is used) | `WARN` | `input-line`, `output-line`, `context`, `path`
A field is skipped (`--skip-field-on-error` is used) | `WARN` | `input-line`, `output-line`, `context`, `path`

## Structured logging

With addition of the `--log-json` flag, the logs format will change to JSON.

Without `--log-json` flag :

```console
$echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo -vdebug --repeat 2
9:14PM INF Logger level set to debug
9:14PM INF Start PIMO config=masking.yml dump-cache={} empty-input=false load-cache={} repeat=2 skipFieldOnError=false skipLineOnError=false
9:14PM INF Add mask config=masking.yml context=stdin mask="constant Benjamin" path=name
9:14PM INF Add mask config=masking.yml context=stdin mask="hash size=2" path=surname
9:14PM INF Add mask config=masking.yml context=stdin mask="hash size=3" path=town
9:14PM INF Add mask config=masking.yml context=stdin mask="randomInt min=18 max=90" path=age
9:14PM DBG Mask constant config=masking.yml context=stdin[1] input-line=1 output-line=1 path=name
9:14PM DBG Mask randomChoice config=masking.yml context=stdin[1] input-line=1 output-line=1 path=surname
9:14PM DBG Mask hash config=masking.yml context=stdin[1] input-line=1 output-line=1 path=town
9:14PM DBG Mask randomInt config=masking.yml context=stdin[1] input-line=1 output-line=1 path=age
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":79}
9:14PM DBG Mask constant config=masking.yml context=stdin[1] input-line=1 output-line=2 path=name
9:14PM DBG Mask randomChoice config=masking.yml context=stdin[1] input-line=1 output-line=2 path=surname
9:14PM DBG Mask hash config=masking.yml context=stdin[1] input-line=1 output-line=2 path=town
9:14PM DBG Mask randomInt config=masking.yml context=stdin[1] input-line=1 output-line=2 path=age
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":53}
9:14PM INF End PIMO config=masking.yml duration=7.2262ms input-line=1 output-line=3 return=0 stats={"ignoredPaths":0,"skippedFields":0,"skippedLines":0}
```

With `--log-json` flag :

```json
$ echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo -vdebug --log-json --repeat 2
{"level":"info","message":"Logger level set to debug"}
{"level":"info","skipLineOnError":false,"skipFieldOnError":false,"repeat":2,"empty-input":false,"dump-cache":{},"load-cache":{},"config":"masking.yml","message":"Start PIMO"}
{"level":"info","path":"name","mask":"constant Benjamin","config":"masking.yml","context":"stdin","message":"Add mask"}
{"level":"info","path":"surname","mask":"hash size=2","config":"masking.yml","context":"stdin","message":"Add mask"}
{"level":"info","path":"town","mask":"hash size=3","config":"masking.yml","context":"stdin","message":"Add mask"}
{"level":"info","path":"age","mask":"randomInt min=18 max=90","config":"masking.yml","context":"stdin","message":"Add mask"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"name","message":"Mask constant"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"surname","message":"Mask randomChoice"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"town","message":"Mask hash"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"age","message":"Mask randomInt"}
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":79}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"name","message":"Mask constant"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"surname","message":"Mask randomChoice"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"town","message":"Mask hash"}
{"level":"debug","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"age","message":"Mask randomInt"}
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":53}
{"level":"info","stats":{"ignoredPaths":0,"skippedLines":0,"skippedFields":0},"return":0,"config":"masking.yml","duration":"5.7093ms","input-line":"1","output-line":"3","message":"End PIMO"}
```

## Structured logging - advanced usage

Structured JSON logging enable usages with other JSON-compatible tools like `jq` or `mlr`.

In this example, logs are pretty-printed with `mlr --opprint --barred`.

```console
$ echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo --log-json -vdebug --repeat 2 2> >(mlr --ijson --opprint --barred cat)
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":79}
{"name":"Benjamin","surname":"Dupont","town":"Ruby City","age":53}
+-------+---------------------------+
| level | message                   |
+-------+---------------------------+
| info  | Logger level set to debug |
+-------+---------------------------+

+-------+-----------------+------------------+--------+-------------+-------------+------------+
| level | skipLineOnError | skipFieldOnError | repeat | empty-input | config      | message    |
+-------+-----------------+------------------+--------+-------------+-------------+------------+
| info  | false           | false            | 2      | false       | masking.yml | Start PIMO |
+-------+-----------------+------------------+--------+-------------+-------------+------------+

+-------+---------+-------------------------+-------------+---------+----------+
| level | path    | mask                    | config      | context | message  |
+-------+---------+-------------------------+-------------+---------+----------+
| info  | name    | constant Benjamin       | masking.yml | stdin   | Add mask |
| info  | surname | hash size=2             | masking.yml | stdin   | Add mask |
| info  | town    | hash size=3             | masking.yml | stdin   | Add mask |
| info  | age     | randomInt min=18 max=90 | masking.yml | stdin   | Add mask |
+-------+---------+-------------------------+-------------+---------+----------+

+-------+-------------+----------+------------+-------------+---------+-------------------+
| level | config      | context  | input-line | output-line | path    | message           |
+-------+-------------+----------+------------+-------------+---------+-------------------+
| debug | masking.yml | stdin[1] | 1          | 1           | name    | Mask constant     |
| debug | masking.yml | stdin[1] | 1          | 1           | surname | Mask randomChoice |
| debug | masking.yml | stdin[1] | 1          | 1           | town    | Mask hash         |
| debug | masking.yml | stdin[1] | 1          | 1           | age     | Mask randomInt    |
| debug | masking.yml | stdin[1] | 1          | 2           | name    | Mask constant     |
| debug | masking.yml | stdin[1] | 1          | 2           | surname | Mask randomChoice |
| debug | masking.yml | stdin[1] | 1          | 2           | town    | Mask hash         |
| debug | masking.yml | stdin[1] | 1          | 2           | age     | Mask randomInt    |
+-------+-------------+----------+------------+-------------+---------+-------------------+

+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
| level | stats:ignoredPaths | stats:skippedLines | stats:skippedFields | return | config      | duration | input-line | output-line | message  |
+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
| info  | 0                  | 0                  | 0                   | 0      | masking.yml | 5.2588ms | 1          | 3           | End PIMO |
+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
```
