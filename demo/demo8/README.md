# Handle complex structures

The `pipe` mask is a powerfull tool to be able to mask data that is nested in a complex multi-level array structure.

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

The truth is by using only the `template` mask (or any other except `pipe`), it is impossible to have the correct expected result. That's why the mask `pipe` was created.
