# Handle complex structures

The `pipe` mask is a powerful tool to be able to mask data that is nested in a complex multi-level array structure.

## Motivation

Consider the following structure :

**`data.json`**
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

- `organisations` is an array of organisation objects. 
- each organisation contains a field `persons`, this field is an array of person objects.

How to mask the `email` field in each person to this format : `{{.person.name}}.{{.person.surname}}@{{.domain}}` ?

### A wrong approach

The first idea that might come to mind is something like: 

**`masking-wrong.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "organizations.persons.email"
    mask:
      # this go template syntax refer to a field that is not in a nested array
      template: "{{.organizations.persons.name}}.{{.organizations.persons.surname}}@{{.organizations.domain}}"
```

Here is the result of applying the above configuration.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

The command `jq -c "."` used below is to reformat an indented multiline json structure into a single line (jsonl).

---

**`oups!`**
```console
$ cat data.json | jq -c "."  | pimo -c masking-wrong.yml
template: template:1:16: executing "template" at <.organizations.persons.name>: can't evaluate field persons in type model.Entry
```

This error occur because the templating syntaxe used by the mask `template` is different as the syntax used in the `jsonpath` property. PIMO can handle arrays and with the path `.organizations.persons.name` it recognize the fields `.organizations[*].persons[*].name` are to be masked (all the names, for all persons, for all organization).

The template mask however wants to know exactly which value to use, and it can't do it with the provided path. Because this path does not point to a valid location in the structure.

### Another wrong approach

The second idea that might come to mind is to try to fix the template syntax.

The way to access an array in go template is :

**`masking-alsowrong.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "organizations.persons.email"
    mask:
      # this go template syntax refer to a single values of index 0 in each array
      # (and it's not very readable)
      template: "{{(index (index .organizations 0).persons 0).name}}.{{(index (index .organizations 0).persons 0).surname}}@{{(index .organizations 0).domain}}"
```

Here is the result of applying the above configuration.

**`uh?`**
```json
$ cat data.json | jq -c "."  | pimo -c masking-alsowrong.yml | jq
{
  "organizations": [
    {
      "domain": "company.com",
      "persons": [
        {
          "email": "leona.miller@company.com",
          "name": "leona",
          "surname": "miller"
        },
        {
          "email": "leona.miller@company.com",
          "name": "joe",
          "surname": "davis"
        }
      ]
    },
    {
      "domain": "company.fr",
      "persons": [
        {
          "email": "leona.miller@company.com",
          "name": "alain",
          "surname": "mercier"
        },
        {
          "email": "leona.miller@company.com",
          "name": "florian",
          "surname": "legrand"
        }
      ]
    }
  ]
}
```

The error is gone, but everyone has the email `leona.miller@company.com` which is not what we want.

The truth is, by using only the `template` mask (or any other except `pipe`), it is impossible to have the correct expected result. That's why the mask `pipe` was created.

## Using the `pipe` mask

This mask can process the persons objects like an independent stream of json.

The usecase exposed in the [previous chapter](#Motivation) is tackled in this part, in 2 steps.

### Step 1 : setting a sub-pipeline to process the persons

This mask can process the persons objects like an independent stream of json. Its content is another `masking` node defining a list of masks to apply.

**`masking-pipe-1.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "organizations.persons"
    mask:
      pipe:
        # starting here is the definition another masking pipeline, that applies on the persons objects
        masking:
          - selector:
              jsonpath: "email"
            mask:
              # in the template, name and surname can be accessed directly
              template: "{{.name}}.{{.surname}}"
```

Here is the result of applying the above configuration.

**`result`**
```json
$ cat data.json | jq -c "."  | pimo -c masking-pipe-1.yml | jq
{
  "organizations": [
    {
      "domain": "company.com",
      "persons": [
        {
          "email": "leona.miller",
          "name": "leona",
          "surname": "miller"
        },
        {
          "email": "joe.davis",
          "name": "joe",
          "surname": "davis"
        }
      ]
    },
    {
      "domain": "company.fr",
      "persons": [
        {
          "email": "alain.mercier",
          "name": "alain",
          "surname": "mercier"
        },
        {
          "email": "florian.legrand",
          "name": "florian",
          "surname": "legrand"
        }
      ]
    }
  ]
}
```

The name and surname parts are now correct. The next step is is to handle the domain part.

### Step 2 : handle the domain

The domain is not part of a person object, but stored in the parent object (organisation).

The parent object can be accessed with the `injectParent` property of the `pipe` mask. The value of the property will be used to name this field.

**`masking-pipe-2.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "organizations.persons"
    mask:
      pipe:
        # the parent of the person will be injected during the processing of the sub-pipeline, under the path ".org"
        # the name "org" is an example, any valid identifier can be chosen
        injectParent: "org"
        masking:
          - selector:
              jsonpath: "email"
            mask:
              # now the template can read the value of the organization domain with .org.domain
              template: "{{.name}}.{{.surname}}@{{.org.domain}}"
```

Here is the result of applying the above configuration.

**`result`**
```json
$ cat data.json | jq -c "."  | pimo -c masking-pipe-2.yml | jq
{
  "organizations": [
    {
      "domain": "company.com",
      "persons": [
        {
          "email": "leona.miller@company.com",
          "name": "leona",
          "surname": "miller"
        },
        {
          "email": "joe.davis@company.com",
          "name": "joe",
          "surname": "davis"
        }
      ]
    },
    {
      "domain": "company.fr",
      "persons": [
        {
          "email": "alain.mercier@company.fr",
          "name": "alain",
          "surname": "mercier"
        },
        {
          "email": "florian.legrand@company.fr",
          "name": "florian",
          "surname": "legrand"
        }
      ]
    }
  ]
}
```

## Advanced usage

### Inject root

The `pipe` mask also expose the `injectRoot` property, similar to `injectParent` except it will inject the whole current structure being processed.

### Externalize sub-pipeline

The sub-pipeline definition can be in another YAML file.

**`masking-root.yml`**
```yaml
version: "1"
masking:
  - selector:
      jsonpath: "organizations.persons"
    mask:
      pipe:
        injectParent: "org"
        file: "masking-org.yml"
```

**`masking-org.yml`**
```yaml
version: "1"
masking:
  - selector:
      jsonpath: "email"
    mask:
      template: "{{.name}}.{{.surname}}@{{.org.domain}}"
```

### Use pipes with caches

Pipes are compatible with caches.

If a cache mut be shared by the main pipeline and all the sub-pipelines, it must be declared at the root.

**`masking-cache.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "age"
    cache: "age"
    mask:
      randomInt:
        min: 0
        max: 100
  - selector:
      jsonpath: "related"
    mask:
      pipe:
        masking:
          - selector:
              jsonpath: "age"
            cache: "age"
            mask:
              randomInt:
                min: 0
                max: 100
# declared here, the cache will be shared by all sub-pipelines and the main pipeline
caches:
  age : {}
```

**`data-cache.jsonl`**
```json
{"age": 10, "related": [{"age":30}]}
{"age": 20, "related": [{"age":40}]}
{"age": 30, "related": [{"age":10}]}
{"age": 40, "related": [{"age":20}]}
```

Here is the result of applying the above configuration.

**`result`**
```json
$ cat data-cache.jsonl | jq -c "."  | pimo -c masking-cache.yml
{"age":91,"related":[{"age":55}]}
{"age":25,"related":[{"age":84}]}
{"age":55,"related":[{"age":91}]}
{"age":84,"related":[{"age":25}]}
```

---
**NOTE**

Pipes are currently **NOT compatible** with the `fromCache` mask.

The use of `fromCache` inside a pipe is **discouraged**, an the results might be unexpected.

However, an approach like the one presented just above, by referencing a cache in multiple position will give a correct and expected result. The only difference with `fromCache` is that the mask must be duplicated in two locations.

---
