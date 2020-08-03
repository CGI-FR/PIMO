# Pimo : Private Input, Masked Output

PIMO is a tool for data masking. It can mask datas from JSONline and return other JSONline thanks to a masking configuration contained in a yaml file.

## Configuration file needed

Pimo requires a yaml configuration file to works. This file must be named `masking.yml` and be placed in the working directory. The file must respect the following format :

```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "exemple.exemple"
    mask:
      type: "argument"
```

`version` is the version of the masking file.  
`seed` is to give every random mask the same seed, it is optional and in case it is not informed, the seed is chosen with the time to look more random.  
`masking` is used to define the masks that are going to be used by pimo.  
`selector` is made of a jsonpath and a mask.  
`jsonpath` defines the path of the entry that has to be masked in the json file.  
`mask` defines the mask that will be used for the entry defined by `selector`.  

## Possible masks

The following types of masks can be used :

* `regex` is to mask using a regular expression given in argument.
* `constant` is to mask the value by a constant value given in argument.
* `randomChoice` is to mask with a random value from a list in argument.
* `randomInt` is to mask with a random int from a range with arguments min and max.
* `command` is to mask with the output of a console command given in argument.
* `weightedChoice` is to mask with a random value from a list with probability, both given with the arguments `choice` and `weight`.
* `hash` is to mask with a value from a list by mashing the original value, allowing to mask a value the same way every time.
* `randDate` is to mask a date with a random date between `dateMin` and `dateMax`.
* `duration` is to mask a date by adding or removing a certain number of days.
* `randomDuration` is to mask a date by adding or removing a random time between `Min` and `Max`
* `incremental` is to mask datas with incremental value starting from `start` with a step of `increment`.
* `remplacement` is to mask a data with another data from the jsonline.
* `template` is to mask a data with a template using other values from the jsonline.
* `remove` is to mask a field by completely removing it.
* `add` is a mask to add a field to the jsonline.

A full `masking.yml` file exemple, using every kind of mask, is given with the source code.

In case two types of mask are entered in the same selector, the program can't extract the masking configuration and will return an error. The file `wrongMasking.yml` provided with the source illustrate that error.

## Usage

To use PIMO to mask a `data.json`, use in the following way :

```bash
./pimo <data.json >maskedData.json
```

This takes the `data.json` file, masks the datas contained inside it and put the result in a `maskedData.json` file. If datas are in a tables (for exemple multiples names), then each field of this table will be masked using the given mask. The following flags can be used:

* `--repeat=N` This flag will make pimo mask every input N-times.
* `--skip-line-on-error` This flag will totally skip a line if an error occurs masking a field.
* `--skip-field-on-error` This flag will return output without a field if an error occurs masking this field.
* `--empty-input` This flag will give PIMO a `{}` input, usable with `repeat` flag.
* `--config=filename.yml` This flag allow to use another file for config that `masking.yml`.

## Exemple

This section will give exemples for every type of masking.

### regex

```yaml
  - selector:
      jsonpath: "phone"
    mask:
      regex: "0[1-7]( ([0-9]){2}){4}"
```

This exemple will mask the `phone` field of the input jsonlines with a random string respecting the regular expression.

### constant

```yaml
  - selector:
      jsonpath: "name"
    mask:
      constant: "Toto"
```

This exemple will mask the `name` field of the input jsonlines with the value of the `constant` field.

### randomChoice

```yaml
  - selector:
      jsonpath: "name"
    mask:
      randomChoice:
       - "Mickael"
       - "Mathieu"
       - "Marcelle"
```

This exemple will mask the `name` field of the input jsonlines with random values from the `randomChoice` list.

### randomChoiceInUri

```yaml
  - selector:
      jsonpath: "name"
    mask:
      randomChoiceInUri: "file://names.txt"
```

This exemple will mask the `name` field of the input jsonlines with random values from the list contained in the name.txt file. The different URI usable with this selector are : `pimo`, `file` and `http`/`https`.

### randonInt

```yaml
  - selector:
      jsonpath: "age"
    mask:
      randomInt:
        min: 25
        max: 32
```

This exemple will mask the `age` field of the input jsonlines with a random number between `min` and `max` included.

### randonDecimal

```yaml
  - selector:
      jsonpath: "score"
    mask:
      randomDecimal:
        min: 0
        max: 17,23
        precision: 2
```

This exemple will mask the `score` field of the input jsonlines with a random float between `min` and `max`, with the number of decimal chosen in the `precision` field.

### command

```yaml
  - selector:
      jsonpath: "name"
    mask:
      command: "echo Dorothy"
```

This exemple will mask the `name` field of the input jsonlines with the output of the given commande. In this case, `Dorothy`.

### weightedChoice

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

This exemple will mask the `surname` field of the input jsonlines with a random value in the `weightedChoice` list with a probability proportional at the `weight` field.

### hash

```yaml
  - selector:
      jsonpath: "town"
    mask:
      hash:
        - "Emerald City"
        - "Ruby City"
        - "Sapphire City"
```

This exemple will mask the `town` field of the input jsonlines with a value from the `hash` list. The value will be chosen thanks to a hashing of the original value, allowing the output to be always the same in case of identical inputs.

### hashInUri

```yaml
  - selector:
      jsonpath: "name"
    mask:
      hashInUri: "pimo://nameFR
```

This exemple will mask the `name` field of the input jsonlines with a value from the list nameFR contained in pimo, the same way as for `hash` mask. The different URI usable with this selector are : `pimo`, `file` and `http`/`https`.

### randDate

```yaml
  - selector:
      jsonpath: "date"
    mask:
      randDate:
        dateMin: "1970-01-01T00:00:00Z"
        dateMax: "2020-01-01T00:00:00Z"
```

This exemple will mask the `date` field of the input jsonlines with a random date between `dateMin` and `dateMax`. In this case the date will be between the 1st January 1970 and the 1st January 2020.

### duration

```yaml
  - selector:
      jsonpath: "last_contact"
    mask:
      duration: "-P2D"
```

This exemple will mask the `last_contact` field of the input jsonlines by decreasing its value by 2 days. The duration field should match the ISO 8601 standard for durations.

### randomDuration

```yaml
  - selector:
      jsonpath: "date"
    mask:
      randomDuration:
        min: "-P2D"
        max: "-P27D"
```

This exemple will mask the `date` field of the input jsonlines by decreasing its value by a random value between 2 and 27 days. The durations should match the ISO 8601 standard.

### incremental

```yaml
  - selector:
      jsonpath: "id"
    mask:
      incremental:
        start: 1
        increment: 1
```

This exemple will mask the `id` field of the input jsonlines with incremental values. The first jsonline's `id` will be masked by 1, the second's by 2, etc...

### replacement

```yaml
  - selector:
      jsonpath: "name4"
    mask:
      replacement: "name"
```

This exemple will mask the `name4` field of the input jsonlines with the field `name` of the jsonline. This selector must be placed after the `name` selector to be masked with the new value and it must be placed before the `name` selector to be masked by the previous value.

### template

```yaml
  - selector:
      jsonpath: "mail"
    mask:
      template: "{{.surname}}.{{.name}}@gmail.com"
```

This exemple will mask the `mail` field of the input jsonlines respecting the given template. In the `masking.yml` config fil, this selector must be placed after the fields contained in the template to mask with the new values and  before the other fields to be masked with the old values. In the case of a nested json, the template must respect the following exemple :

```yaml
  - selector:
      jsonpath: "user.mail"
    mask:
      template: "{{.user.surname}}.{{.user.name}}@gmail.com"
```

The format for the template should respect the `text/template` package : <https://golang.org/pkg/text/template/>

The template mask can format the fields used. The following exemple will create a mail address without accent or upper case:

```yaml
  - selector:
      jsonpath: "user.mail"
    mask:
      template: "{{.surname | NoAccent | upper}}.{{.name | NoAccent | lower}}@gmail.com"
```

Available functions for templates come from <http://masterminds.github.io/sprig/>. The function NoAccent can also be used.

### Remove

```yaml
  - selector:
      jsonpath: "useless-field"
    mask:
      remove: true
```

This field will mask the `useless-field` of the input jsonlines by completely deleting it.

### Add

```yaml
  - selector:
      jsonpath: "newField"
    mask:
      add: "newvalue"
```

This exemple will create the field `newField` containing the value `newvalue`. This value can be a string, a number, a boolean...

The field will be created in every input jsonline that doesn't already contains this field.

### FluxURI

```yaml
  - selector:
      jsonpath: "id"
    mask:
      fluxURI: "file://id.csv"
```

This exemple will create an `id` field in every output jsonline. The values will be the ones contained in the `id.csv` file in the same order as in the file. If the field already exist on the input jsonline it will be replaced and if every value of the file has already been assigned, the input jsonlines won't be modified.
