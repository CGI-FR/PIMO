# Logging

PIMO has 6 level of logging :
1. none : log nothing at all
2. error : log only errors
3. warn : same as level 1 + warnings that should be checked by user
4. info : same as level 2 + information about what is processed
5. debug : same as level 3 + debugging information, to analyse what can cause an unexpected behavior
6. trace : same as level 4 + tracing of events in code (enter function, exit function, values...)

This can be set with the `-v` flag (long version `--verbosity`) :

```console
$ pimo --empty-input -v5
2:57PM INF Logger level set to trace
2:57PM INF Start PIMO config=masking.yml dump-cache={} empty-input=true load-cache={} repeat=1 skipFieldOnError=false skipLineOnError=false
2:57PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=name
2:57PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=surname
2:57PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=town
2:57PM WRN Path not found config=masking.yml context=empty-input[1] input-line=1 output-line=1 path=age
{}
2:57PM INF End PIMO config=masking.yml duration=4.0093ms input-line=1 output-line=2 return=0 stats={"ignoredPaths":4,"skippedFields":0,"skippedLines":0}
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
A mask is applied | `INFO` | `input-line`, `output-line`, `context`, `path`
A jsonpath point to nothing | `WARN` | `input-line`, `output-line`, `context`, `path`
A line is skipped (`--skip-line-on-error` is used) | `WARN` | `input-line`, `output-line`, `context`, `path`
A field is skipped (`--skip-field-on-error` is used) | `WARN` | `input-line`, `output-line`, `context`, `path`

## Structured logging

With addition of the `--log-json` flag, the logs format will change to JSON. 

Without `--log-json` flag :

```console
$ echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo -vinfo --repeat 2
3:26PM INF Logger level set to info
3:26PM INF Start PIMO config=masking.yml dump-cache={} empty-input=false load-cache={} repeat=2 skipFieldOnError=false skipLineOnError=false
3:26PM INF Mask constant config=masking.yml context=stdin[1] input-line=1 output-line=1 path=name
3:26PM INF Mask randomChoice config=masking.yml context=stdin[1] input-line=1 output-line=1 path=surname
3:26PM INF Mask hash config=masking.yml context=stdin[1] input-line=1 output-line=1 path=town
3:26PM INF Mask randomInt config=masking.yml context=stdin[1] input-line=1 output-line=1 path=age
{"age":79,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
3:26PM INF Mask constant config=masking.yml context=stdin[1] input-line=1 output-line=2 path=name
3:26PM INF Mask randomChoice config=masking.yml context=stdin[1] input-line=1 output-line=2 path=surname
3:26PM INF Mask hash config=masking.yml context=stdin[1] input-line=1 output-line=2 path=town
3:26PM INF Mask randomInt config=masking.yml context=stdin[1] input-line=1 output-line=2 path=age
{"age":53,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
3:26PM INF End PIMO config=masking.yml duration=10.3316ms input-line=1 output-line=3 return=0 stats={"ignoredPaths":0,"skippedFields":0,"skippedLines":0}
```

With `--log-json` flag :

```json
$ echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo -vinfo --repeat 2
{"level":"info","message":"Logger level set to info"}
{"level":"info","skipLineOnError":false,"skipFieldOnError":false,"repeat":2,"empty-input":false,"dump-cache":{},"load-cache":{},"config":"masking.yml","message":"Start PIMO"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"name","message":"Mask constant"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"surname","message":"Mask randomChoice"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"town","message":"Mask hash"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"1","path":"age","message":"Mask randomInt"}
{"age":79,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"name","message":"Mask constant"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"surname","message":"Mask randomChoice"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"town","message":"Mask hash"}
{"level":"info","config":"masking.yml","context":"stdin[1]","input-line":"1","output-line":"2","path":"age","message":"Mask randomInt"}
{"age":53,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
{"level":"info","stats":{"ignoredPaths":0,"skippedLines":0,"skippedFields":0},"return":0,"config":"masking.yml","duration":"10.8207ms","input-line":"1","output-line":"3","message":"End PIMO"}
```

## Structured logging - advanced usage

Structured JSON logging enable usages with other JSON-compatible tools like `jq` or `mlr`.

In this example, logs are pretty-printed with `mlr --opprint --barred`.

```console
echo '{"name": "", "surname": "", "town": "", "age": ""}' | pimo --log-json -vinfo --repeat 2 2> >(mlr --ijson --opprint --barred cat)
{"age":79,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
{"age":53,"name":"Benjamin","surname":"Dupont","town":"Ruby City"}
+-------+--------------------------+
| level | message                  |
+-------+--------------------------+
| info  | Logger level set to info |
+-------+--------------------------+

+-------+-----------------+------------------+--------+-------------+-------------+------------+
| level | skipLineOnError | skipFieldOnError | repeat | empty-input | config      | message    |
+-------+-----------------+------------------+--------+-------------+-------------+------------+
| info  | false           | false            | 2      | false       | masking.yml | Start PIMO |
+-------+-----------------+------------------+--------+-------------+-------------+------------+

+-------+-------------+----------+------------+-------------+---------+-------------------+
| level | config      | context  | input-line | output-line | path    | message           |
+-------+-------------+----------+------------+-------------+---------+-------------------+
| info  | masking.yml | stdin[1] | 1          | 1           | name    | Mask constant     |
| info  | masking.yml | stdin[1] | 1          | 1           | surname | Mask randomChoice |
| info  | masking.yml | stdin[1] | 1          | 1           | town    | Mask hash         |
| info  | masking.yml | stdin[1] | 1          | 1           | age     | Mask randomInt    |
| info  | masking.yml | stdin[1] | 1          | 2           | name    | Mask constant     |
| info  | masking.yml | stdin[1] | 1          | 2           | surname | Mask randomChoice |
| info  | masking.yml | stdin[1] | 1          | 2           | town    | Mask hash         |
| info  | masking.yml | stdin[1] | 1          | 2           | age     | Mask randomInt    |
+-------+-------------+----------+------------+-------------+---------+-------------------+

+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
| level | stats:ignoredPaths | stats:skippedLines | stats:skippedFields | return | config      | duration | input-line | output-line | message  |
+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
| info  | 0                  | 0                  | 0                   | 0      | masking.yml | 2.1671ms | 1          | 3           | End PIMO |
+-------+--------------------+--------------------+---------------------+--------+-------------+----------+------------+-------------+----------+
```
