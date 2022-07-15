# PIMO : Private Input, Masked Output

PIMO is a tool for data masking. It can mask data from a JSONline stream and return another JSONline stream thanks to a masking configuration contained in a yaml file.

## Configuration file needed

PIMO requires a yaml configuration file to works. By default, the file is named `masking.yml` and is placed in the working directory. The file must respect the following format :

```yaml
version: "1"
seed: 42
functions:
    # Optional define functions
masking:
  - selector:
      jsonpath: "example.example"
    mask:
      type: "argument"
    # Optional cache (coherence preservation)
    cache: "cacheName"
    # Optional custom seed for this mask
    seed:
      field: "example.example"

  # another mask on a different location
  - selector:
      jsonpath: "example.example2"
    mask:
      type: "argument"
    preserve: "null"

caches:
  cacheName:
    # Optional bijective cache (enable re-identification if the cache is dumped on disk)
    unique: true
    # Use reverse cache dictionnary
    reverse: true
```

`version` is the version of the masking file.
`seed` is to give every random mask the same seed, it is optional and if it is not defined, the seed is derived from the current time to increase randomness.
`functions` is used to define the functions that can be used in the te mask `template`, `template-each`, `add`, and `add-transient`.
`masking` is used to define the pipeline of masks that is going to be applied.
`selector` is made of a jsonpath and a mask.
`jsonpath` defines the path of the entry that has to be masked in the json file.
`mask` defines the mask that will be used for the entry defined by `selector`.
`cache` is optional, if the current entry is already in the cache as key the associated value is returned without executing the mask. Otherwise the mask is executed and a new entry is added in the cache with the orignal content as `key` and the masked result as `value`. The cache have to be declared in the `caches` section of the YAML file.
`preserve` is optional, and is used to keep some values unmasked in the json file. Allowed `preserve` options are: `"null"` (null values), `"empty"` (empty string `""`), and `"blank"` (both `empty` and `null` values). Additionally, `preserve` can be used with mask [`fromCache`](#fromCache) to preserve uncached values. (usage: `preserve: "notInCache"`)

Multiple masks can be applied on the same jsonpath location, like in this example :

```yaml
  - selector:
      jsonpath: "example"
    masks:
      - add: "hello"
      - template: "{{.example}} World!"
      - remove: true
```

Masks can be applied on multiple selectors, like in this example:

```yaml
  - selectors:
      - jsonpath: "example"
      - jsonpath: "example2"
    mask:
      add: "hello"
```

It is possible to define functions and reuse them later in the masks, like in this example:

```yaml
functions:
  add20:
    params:
      - name: "i"
    body: |-
      return i + 20
  sub:
    params:
      - name: "x"
      - name: "y"
    body: |-
      return x - y
masking:
  - selector:
      jsonpath: "addValue"
    mask:
      template: '{{add20 5}}'
  - selector:
      jsonpath: "subValue"
    mask:
      template: '{{sub 10 5}}'
```

## Possible masks

The following types of masks can be used :

* Pure randomization masks
  * [`regex`](#regex) is to mask using a regular expression given in argument.
  * [`randomInt`](#randomInt) is to mask with a random int from a range with arguments min and max.
  * [`randomDecimal`](#randomDecimal) is to mask with a random decimal from a range with arguments min, max and precision.
  * [`randDate`](#randDate) is to mask a date with a random date between `dateMin` and `dateMax`.
  * [`randomDuration`](#randomDuration) is to mask a date by adding or removing a random time between `Min` and `Max`.
  * [`randomChoice`](#randomChoice) is to mask with a random value from a list in argument.
  * [`weightedChoice`](#weightedChoice) is to mask with a random value from a list with probability, both given with the arguments `choice` and `weight`.
  * [`randomChoiceInUri`](#randomChoiceInUri) is to mask with a random value from an external resource.
  * [`transcode`](#transcode) is to mask a value randomly with character class preservation.
* K-Anonymization
  * [`range`](#range) is to mask a integer value by a range of value (e.g. replace `5` by `[0,10]`).
  * [`duration`](#duration) is to mask a date by adding or removing a certain number of days.
* Re-identification and coherence preservation
  * [`hash`](#hash) is to mask with a value from a list by matching the original value, allowing to mask a value the same way every time.
  * [`hashInUri`](#hashInUri) is to mask with a value from an external resource, by matching the original value, allowing to mask a value the same way every time.
  * [`fromCache`](#fromCache) is a mask to obtain a value from a cache.
  * [`ff1`](#ff1) mask allows the use of <abbr title="Format Preserving Encryption">FPE</abbr> which enable private-key based re-identification.
* Formatting
  * [`dateParser`](#dateParser) is to change a date format.
  * [`template`](#template) is to mask a data with a template using other values from the jsonline.
  * [`template-each`](#template-each) is like template but will apply on each value of an array.
  * [`fromjson`](#fromjson) is to convert string field values to parsed JSON, e.g. "[1,2,3]" -> [1,2,3].
* Data structure manipulation
  * [`remove`](#remove) is to mask a field by completely removing it.
  * [`add`](#add) is a mask to add a field to the jsonline.
  * [`add-transient`](#add-transient) same as `add` but the field is not exported in the output jsonline.
* Others
  * [`constant`](#constant) is to mask the value by a constant value given in argument.
  * [`command`](#command) is to mask with the output of a console command given in argument.
  * [`incremental`](#incremental) is to mask data with incremental value starting from `start` with a step of `increment`.
  * [`fluxUri`](#fluxUri) is to replace by a sequence of values defined in an external resource.
  * [`replacement`](#replacement) is to mask a data with another data from the jsonline.
  * [`pipe`](#pipe) is a mask to handle complex nested array structures, it can read an array as an object stream and process it with a sub-pipeline.
  * [`luhn`](#luhn) can generate valid numbers using the Luhn algorithm (e.g. french SIRET or SIREN).
  * [`markov`](#markov) can generate pseudo text based on a sample text.

A full `masking.yml` file example, using every kind of mask, is given with the source code.

In case two types of mask are entered with the same selector, the program can't extract the masking configuration and will return an error. The file `wrongMasking.yml` provided with the source illustrate that error.

## Usage

To use PIMO to mask a `data.json`, use in the following way :

```bash
./pimo <data.json >maskedData.json
```

This takes the `data.json` file, masks the data contained inside it and put the result in a `maskedData.json` file. If data are in a table (for example multiple names), then each field of this table will be masked using the given mask. The following flags can be used:

* `--repeat=N` This flag will make pimo mask every input N-times (useful for dataset generation).
* `--skip-line-on-error` This flag will totally skip a line if an error occurs masking a field.
* `--skip-field-on-error` This flag will return output without a field if an error occurs masking this field.
* `--empty-input` This flag will give PIMO a `{}` input, usable with `--repeat` flag.
* `--config=filename.yml` This flag allow to use another file for config than the default `masking.yml`.
* `--load-cache cacheName=filename.json` This flag load an initial cache content from a file (json line format `{"key":"a", "value":"b"}`).
* `--dump-cache cacheName=filename.json` This flag dump final cache content to a file (json line format `{"key":"a", "value":"b"}`).
* `--verbosity <level>` or `-v<level>` This flag increase verbosity on the stderr output, possible values: none (0), error (1), warn (2), info (3), debug (4), trace (5).
* `--debug` This flag complete the logs with debug information (source file, line number).
* `--log-json` Set this flag to produce JSON formatted logs ([demo9](demo/demo9) goes deeper into logging and structured logging)
* `--mask` Declare a simple masking definition in command line (minified YAML format: `--mask "value={fluxUri: 'pimo://nameFR'}"`, or `--mask "value=[{add: ''},{fluxUri: 'pimo://nameFR'}]"` for multiple masks). For advanced use case (e.g. if caches needed) `masking.yml` file definition will be preferred.
* `--repeat-until <condition>` This flag will make PIMO keep masking every input until the condition is met. Condition format is using [Template](https://pkg.go.dev/text/template). Last output verifies the condition.
* `--repeat-while <condition>` This flag will make PIMO keep masking every input while the condition is met. Condition format is using [Template](https://pkg.go.dev/text/template).

### PIMO Play

The `play` command will start a local website, where you will find commented examples and a playground to play with the masking configuration.

```console
$ pimo play
⇨ http server started on [::]:3010
```

Then go to [http://localhost:3010/](http://localhost:3010/) in your browser.

![PIMO Play screenshot](assets/pimo-play.png)

## Examples

This section will give examples for every types of mask.

Please check the [demo folder](demo) for more advanced examples.

### Regex

```yaml
  - selector:
      jsonpath: "phone"
    mask:
      regex: "0[1-7]( ([0-9]){2}){4}"
```

This example will mask the `phone` field of the input jsonlines with a random string respecting the regular expression.

[Return to list of masks](#possible-masks)

### Constant

```yaml
  - selector:
      jsonpath: "name"
    mask:
      constant: "Bill"
```

This example will mask the `name` field of the input jsonlines with the value of the `constant` field.

[Return to list of masks](#possible-masks)

### RandomChoice

```yaml
  - selector:
      jsonpath: "name"
    mask:
      randomChoice:
       - "Mickael"
       - "Mathieu"
       - "Marcelle"
```

This example will mask the `name` field of the input jsonlines with random values from the `randomChoice` list.

[Return to list of masks](#possible-masks)

### RandomChoiceInUri

```yaml
  - selector:
      jsonpath: "name"
    mask:
      randomChoiceInUri: "file://names.txt"
```

This example will mask the `name` field of the input jsonlines with random values from the list contained in the name.txt file. The different URI usable with this selector are : `pimo`, `file` and `http`/`https`.

A value can be injected in URI with the template syntax. For example, `file://name{{.gender}}.txt` select a line in `name_F.txt` if the current jsonline is `{gender : "F"}`.

[Return to list of masks](#possible-masks)

### RandomInt

```yaml
  - selector:
      jsonpath: "age"
    mask:
      randomInt:
        min: 25
        max: 32
```

This example will mask the `age` field of the input jsonlines with a random number between `min` and `max` included.

[Return to list of masks](#possible-masks)

### RandomDecimal

```yaml
  - selector:
      jsonpath: "score"
    mask:
      randomDecimal:
        min: 0
        max: 17.23
        precision: 2
```

This example will mask the `score` field of the input jsonlines with a random float between `min` and `max`, with the number of decimal chosen in the `precision` field.

[Return to list of masks](#possible-masks)

### Command

```yaml
  - selector:
      jsonpath: "name"
    mask:
      command: "echo -n Dorothy"
```

This example will mask the `name` field of the input jsonlines with the output of the given command. In this case, `Dorothy`.

[Return to list of masks](#possible-masks)

### WeightedChoice

```yaml
  - selector:
      jsonpath: "surname"
    mask:
      weightedChoice:
        - choice: "Dupont"
          weight: 9
        - choice: "Dupond"
          weight: 1
```

This example will mask the `surname` field of the input jsonlines with a random value in the `weightedChoice` list with a probability proportional at the `weight` field.

[Return to list of masks](#possible-masks)

### Hash

```yaml
  - selector:
      jsonpath: "town"
    mask:
      hash:
        - "Emerald City"
        - "Ruby City"
        - "Sapphire City"
```

This example will mask the `town` field of the input jsonlines with a value from the `hash` list. The value will be chosen thanks to a hashing of the original value, allowing the output to be always the same in case of identical inputs.

[Return to list of masks](#possible-masks)

### HashInUri

```yaml
  - selector:
      jsonpath: "name"
    mask:
      hashInUri: "pimo://nameFR"
```

This example will mask the `name` field of the input jsonlines with a value from the list nameFR contained in pimo, the same way as for `hash` mask. The different URI usable with this selector are : `pimo`, `file` and `http`/`https`.

[Return to list of masks](#possible-masks)

### RandDate

```yaml
  - selector:
      jsonpath: "date"
    mask:
      randDate:
        dateMin: "1970-01-01T00:00:00Z"
        dateMax: "2020-01-01T00:00:00Z"
```

This example will mask the `date` field of the input jsonlines with a random date between `dateMin` and `dateMax`. In this case the date will be between the 1st January 1970 and the 1st January 2020.

[Return to list of masks](#possible-masks)

### Duration

```yaml
  - selector:
      jsonpath: "last_contact"
    mask:
      duration: "-P2D"
```

This example will mask the `last_contact` field of the input jsonlines by decreasing its value by 2 days. The duration field should match the ISO 8601 standard for durations.

[Return to list of masks](#possible-masks)

### DateParser

```yaml
  - selector:
      jsonpath: "date"
    mask:
      dateParser:
        inputFormat: "2006-01-02"
        outputFormat: "01/02/06"
```

This example will change every date from the date field from the `inputFormat` to the `outputFormat`. The format should always display the following date : `Mon Jan 2 15:04:05 -0700 MST 2006`. Either field is optional and in case a field is not defined, the default format is RFC3339, which is the base format for PIMO, needed for `duration` mask and given by `randDate` mask. It is possible to use the Unix time format by specifying `inputFormat: "unixEpoch"` or `outputFormat: "unixEpoch"`.

[Return to list of masks](#possible-masks)

### RandomDuration

```yaml
  - selector:
      jsonpath: "date"
    mask:
      randomDuration:
        min: "-P2D"
        max: "-P27D"
```

This example will mask the `date` field of the input jsonlines by decreasing its value by a random value between 2 and 27 days. The durations should match the ISO 8601 standard.

[Return to list of masks](#possible-masks)

### Incremental

```yaml
  - selector:
      jsonpath: "id"
    mask:
      incremental:
        start: 1
        increment: 1
```

This example will mask the `id` field of the input jsonlines with incremental values. The first jsonline's `id` will be masked by 1, the second's by 2, etc...

[Return to list of masks](#possible-masks)

### Replacement

```yaml
  - selector:
      jsonpath: "name4"
    mask:
      replacement: "name"
```

This example will mask the `name4` field of the input jsonlines with the field `name` of the jsonline. This selector must be placed after the `name` selector to be masked with the new value and it must be placed before the `name` selector to be masked by the previous value.

[Return to list of masks](#possible-masks)

### Template

```yaml
  - selector:
      jsonpath: "mail"
    mask:
      template: "{{.surname}}.{{.name}}@gmail.com"
```

This example will mask the `mail` field of the input jsonlines respecting the given template. In the `masking.yml` config file, this selector must be placed after the fields contained in the template to mask with the new values and before the other fields to be masked with the old values. In the case of a nested json, the template must respect the following example :

```yaml
  - selector:
      jsonpath: "user.mail"
    mask:
      template: "{{.user.surname}}.{{.user.name}}@gmail.com"
```

The format for the template should respect the `text/template` package : <https://golang.org/pkg/text/template/>

The template mask can format the fields used. The following example will create a mail address without accent or upper case:

```yaml
  - selector:
      jsonpath: "user.mail"
    mask:
      template: "{{.surname | NoAccent | upper}}.{{.name | NoAccent | lower}}@gmail.com"
```

Available functions for templates come from <http://masterminds.github.io/sprig/>.

[Return to list of masks](#possible-masks)

### Template each

```yaml
  - selector:
      jsonpath: "array"
    mask:
      template-each:
        template: "{{title .value}}"
        item: "value"
```

This will affect every values in the array field. The field must be an array (`{"array": ["value1", "value2]`).
The `item` property is optional and defines the name of the current item in the templating string (defaults to "it"). There is another optional property `index`, if defined then a property with the given name will be available in the templating string (e.g. : `index: "idx"` can be used in template with `{{.idx}}`).

The format for the template should respect the `text/template` package : <https://golang.org/pkg/text/template/>

See also the [Template mask](#template) for other options, all functions are applicable on template-each.

[Return to list of masks](#possible-masks)

### Fromjson

```yaml
  - selector:
      jsonpath: "targetfield"
    mask:
      fromjson: "sourcefield"
```

This example will mask the `targetfield` field of the input jsonlines with the parsed JSON from field `sourcefield` of the jsonline. This mask changes the type of the input string (`sourcefield`) :

* null : nil
* string: string
* number: float64
* array:  slice
* object: map
* bool: bool

[Return to list of masks](#possible-masks)

### Remove

```yaml
  - selector:
      jsonpath: "useless-field"
    mask:
      remove: true
```

This field will mask the `useless-field` of the input jsonlines by completely deleting it.

[Return to list of masks](#possible-masks)

### Add

```yaml
  - selector:
      jsonpath: "newField"
    mask:
      add: "newvalue"
```

This example will create the field `newField` containing the value `newvalue`. This value can be a string, a number, a boolean...

The field will be created in every input jsonline that doesn't already contains this field.

[Return to list of masks](#possible-masks)

### Add-Transient

```yaml
  - selector:
      jsonpath: "newField"
    mask:
      add-transient: "newvalue"
```

This example will create the field `newField` containing the value `newvalue`. This value can be a string, a number, a boolean... It can also be a [template](#template).

The field will be created in every input jsonline that doesn't already contains this field, and it will be removed from the final JSONLine output.

This mask is used for temporary field that is only available to other fields during the execution.

[Return to list of masks](#possible-masks)

### FluxURI

```yaml
  - selector:
      jsonpath: "id"
    mask:
      fluxURI: "file://id.csv"
```

This example will create an `id` field in every output jsonline. The values will be the ones contained in the `id.csv` file in the same order as in the file. If the field already exist on the input jsonline it will be replaced and if every value of the file has already been assigned, the input jsonlines won't be modified.

[Return to list of masks](#possible-masks)

### FromCache

```yaml
  - selector:
      jsonpath: "id"
    mask:
      fromCache: "fakeId"
  caches:
    fakeId :
      unique: true
      reverse: false
```

This example will replace the content of `id` field by the matching content in the cache `fakeId`. Cache have to be declared in the `caches` section.
Cache content can be loaded from jsonfile with the `--load-cache fakeId=fakeId.jsonl` option or by the `cache` option on another field.
If no matching is found in the cache, `fromCache` block the current line and the next lines are processing until a matching content go into the cache.
A `reverse` option is available in the `caches` section to use the reverse cache dictionary.

[Return to list of masks](#possible-masks)

### FF1

```yaml
  - selector:
      jsonpath: "siret"
    mask:
      ff1:
        radix: 10
        keyFromEnv: "FF1_ENCRYPTION_KEY"
```

This example will encrypt the `siret` column with the private key base64-encoded in the FF1_ENCRYPTION_KEY environment variable. Use the same mask with the option `decrypt: true` to re-identify the unmasked value.

Be sure to check [the full FPE demo](demo/demo7) to get more details about this mask.

[Return to list of masks](#possible-masks)

### Range

```yaml
  - selector:
      jsonpath: "age"
    mask:
      range: 5
```

This mask will replace an integer value `{"age": 27}` with a range like this `{"age": "[25;29]"}`.

[Return to list of masks](#possible-masks)

### Pipe

If the data structure contains arrays of object like in the example below, this mask can pipe the objects into a sub pipeline definition.

**`data.jsonl`**

```json
{
    "organizations": [
        {
            "domain": "company.com",
            "persons": [
                {
                    "name": "leona",
                    "surname": "miller",
                    "email": ""
                },
                {
                    "name": "joe",
                    "surname": "davis",
                    "email": ""
                }
            ]
        },
        {
            "domain": "company.fr",
            "persons": [
                {
                    "name": "alain",
                    "surname": "mercier",
                    "email": ""
                },
                {
                    "name": "florian",
                    "surname": "legrand",
                    "email": ""
                }
            ]
        }
    ]
}
```

**`masking.yml`**

```yaml
version: "1"
seed: 42
masking:
  - selector:
      # this path points to an array of persons
      jsonpath: "organizations.persons"
    mask:
      # it will be piped to the masking pipeline definition below
      pipe:
        # the parent object (a domain) will be accessible with the "_" variable name
        injectParent: "_"
        masking:
          - selector:
              jsonpath: "name"
            mask:
              # fields inside the person object can be accessed directly
              template: "{{ title .name }}"
          - selector:
              jsonpath: "surname"
            mask:
              template: "{{ title .surname }}"
          - selector:
              jsonpath: "email"
            mask:
              # the value stored inside the parent object is accessible through "_" thanks to the parent injection
              template: "{{ lower .name }}.{{ lower .surname }}@{{ ._.domain }}"
```

In addition to the `injectParent` property, this mask also provide the `injectRoot` property to inject the whole structure of data.

It is possible to simplify the `masking.yml` file by referencing an external yaml definition :

```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "organizations.persons"
    mask:
      pipe:
        injectParent: "domain"
        file: "./masking-person.yml"
```

Be sure to check [demo](demo/demo8) to get more details about this mask.

[Return to list of masks](#possible-masks)

### Luhn

The [Luhn](https://en.wikipedia.org/wiki/Luhn_algorithm) algorithm is a simple checksum formula used to validate a variety of identification numbers.

The `luhn` mask can calculate the checksum for any value.

```yaml
  - selector:
      jsonpath: "siret"
    mask:
      luhn: {}
```

In this example, the `siret` value will be appended with the correct checksum, to create a valid SIRET number (french business identifier).

The mask can be parametered to use a different universe of valid characters, internally using the [Luhn mod N](https://en.wikipedia.org/wiki/Luhn_mod_N_algorithm) algorithm.

```yaml
  - selector:
      jsonpath: "siret"
    mask:
      luhn:
        universe: "abcde"
```

[Return to list of masks](#possible-masks)


### Markov

[Markov chains](https://en.wikipedia.org/wiki/Markov_chain#Markov_text_generators) produces pseudo text based on an sample text.

**sample.txt**

```txt
I want a cheese burger
I need a cheese cake
```

**masking.yml**

```yaml
  - selector:
      jsonpath: "comment"
    mask:
      markov:
        max-size: 20
        sample: "file://sample.txt"
        separator: " "
```

This example will mask the surname comment of the input jsonlines with a random value comment generated by the markov mask with an order of `2`. The different possibilities generated from **sample.txt** will be :

```txt
I want a cheese burger
I need a cheese burger
I want a cheese cake
I need a cheese cake
```

The `separator` field defines the way the sample text will be split (`""` for splitting into characters, `" "` for splitting into words)

[Return to list of masks](#possible-masks)

### Transcode

This mask produce a random string by preserving character classes from the original value.

**masking.yml**

```yaml
- selector:
    jsonpath: "id"
  mask:
    transcode:
      classes:
      - input: "0123456789abcdefABCDEF"
        output: "0123456789abcdef"
```

This example will mask the original id value by replacing every characters from the `input` class by a random character from the `output` class.

```console
$ echo '{"id": "1ef619-90F"}' | pimo
{"id": "d8e203-a92"}
```

By default, if not specified otherwise, these classes will be used (input -> output):

* lowercase letters -> lowercase letters
* UPPERCASE LETTERS -> UPPERCASE LETTERS
* Digits -> Digits

```yaml
# this configuration:
- selector:
    jsonpath: "id"
  mask:
    transcode: {}
# is equivalent to:
- selector:
    jsonpath: "id"
  mask:
    transcode:
      classes:
        - input: "abcdefghijklmnopqrstuvwxyz"
          output: "abcdefghijklmnopqrstuvwxyz"
        - input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
          output: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        - input: "0123456789"
          output: "0123456789"
```

[Return to list of masks](#possible-masks)

## Flow chart

PIMO can generate a Mermaid syntax flow chart to visualize the transformation process.

for example the command `pimo flow masking.yml > masing.mmd` with that [masking.yml](masking.yml) file generate following chart :

[![](https://mermaid.ink/img/pako:eNqtltFumzAUhl_F9RXRSESArQoXVbVQaZMWbUo77SJElQsOWAMbga02SvNEfYy92GxDUHDoBdoQCvZ_Dieffx9HOcCYJRgGcJez5zhDFQff1hEF8roitBR8Y-nHZAum0xsQi5qzAlezMmMUN3m1eEorVGZG8LFOm7i6-iFV6jWCa5ziF8vZzKfXWwtYG2e62E4O7nFy8I-TCL6aBedNPUyTZmCGNeEVE1zy3nzXj-37S6GoMBegpB62ElrYJaM1R5Rbn0meazqdbTI14ngSdwDFvWBxT84hmrBimTESY2tF4t_oz1turxDPCBbyWcU4z3FH6Q5iumM5UWoaJpUeo5z3CL9Kv1aEBsD9aIMVegmA52oq9aLJpLXxznkDznkXznndNhaFRLNwnDEQsorxbN_55A365I2lqkU10Fyt2iNrtZbtFyZpxnHS7msoSkY5uAULWw8TOZxr1lMpk7bTR-5rklS4rmecPVNzg89C_Z0-C7T4X1CdWXfyPKI8AUvC9_ZaPO2b0T0qy4xUWM-aDjivfNEKveC41SSIm9YrqUevhLNGDeXUUh-6V-eLa2fqzOUNHCfQN_jgyAv8fFjaQOepVnYd9_08vUb9vebaGnF8n_sDfe5f9Lnf_baWOYpxgeUJVHrX4_5gj_tjiQpEcgNIST0eJbQ4D7iQPNLlw2HWdunxOJOTZnSbqtxZzArNqSvNL45Jw_jP1f5Tqc6_RhxnX45q_hjL441ibth4HurZeR5o6UNRIU4YtaY_PjmhRuy9bqL2g0PI0IbyBMslJfI_wUG9FUGeyT6KYCCHCd4hkfMIRvQoU0WpuvkuIZxVMNihvMY2RIKz-z2NYcArgU9JIUFyfUWbdfwLdee3Rg)](https://mermaid.live/edit/#pako:eNqtltFumzAUhl_F9RXRSESArQoXVbVQaZMWbUo77SJElQsOWAMbga02SvNEfYy92GxDUHDoBdoQCvZ_Dieffx9HOcCYJRgGcJez5zhDFQff1hEF8roitBR8Y-nHZAum0xsQi5qzAlezMmMUN3m1eEorVGZG8LFOm7i6-iFV6jWCa5ziF8vZzKfXWwtYG2e62E4O7nFy8I-TCL6aBedNPUyTZmCGNeEVE1zy3nzXj-37S6GoMBegpB62ElrYJaM1R5Rbn0meazqdbTI14ngSdwDFvWBxT84hmrBimTESY2tF4t_oz1turxDPCBbyWcU4z3FH6Q5iumM5UWoaJpUeo5z3CL9Kv1aEBsD9aIMVegmA52oq9aLJpLXxznkDznkXznndNhaFRLNwnDEQsorxbN_55A365I2lqkU10Fyt2iNrtZbtFyZpxnHS7msoSkY5uAULWw8TOZxr1lMpk7bTR-5rklS4rmecPVNzg89C_Z0-C7T4X1CdWXfyPKI8AUvC9_ZaPO2b0T0qy4xUWM-aDjivfNEKveC41SSIm9YrqUevhLNGDeXUUh-6V-eLa2fqzOUNHCfQN_jgyAv8fFjaQOepVnYd9_08vUb9vebaGnF8n_sDfe5f9Lnf_baWOYpxgeUJVHrX4_5gj_tjiQpEcgNIST0eJbQ4D7iQPNLlw2HWdunxOJOTZnSbqtxZzArNqSvNL45Jw_jP1f5Tqc6_RhxnX45q_hjL441ibth4HurZeR5o6UNRIU4YtaY_PjmhRuy9bqL2g0PI0IbyBMslJfI_wUG9FUGeyT6KYCCHCd4hkfMIRvQoU0WpuvkuIZxVMNihvMY2RIKz-z2NYcArgU9JIUFyfUWbdfwLdee3Rg)

## Visual Studio Code

To integrate with Visual Studio Code (opens new window), download the [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml).

Then, edit your Visual Studio Code settings `yaml.schemas` to containing the following configuration:

```json
{
  "yaml.schemas": {
    "https://raw.githubusercontent.com/CGI-FR/PIMO/main/schema/v1/pimo.schema.json": "/**/*masking*.yml"
  }
}
```

Using this configuration, the schema will be applied on every YAML file containing the word `masking`` in their name.

## Contributors

* CGI France ✉[Contact support](mailto:LINO.fr@cgi.com)
* Pole Emploi

## Licence

Copyright (C) 2021 CGI France

PIMO is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

PIMO is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
 along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

