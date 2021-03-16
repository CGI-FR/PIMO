# Third demo - Masking Tables of data

This folder contains files for the third demo. This demo shows how PIMO masks data contained in tables, you can see those different data configurations in the `data.jsonl` file.

#### **`data.jsonl`**
```json
{"name":["Benjamin", "Toto", "Louis"], "surname": ["Martin", "Bernard"]}
{"name":["Benjamin", "Toto", "Louis"], "surname": "Martin"}
{"name":"Benjamin", "surname": ["Martin", "Bernard"]}
{"name":"Benjamin", "surname": "Martin"}
{"identity":{"name":["Jeanne", "Louise"]}}
```

The `masking.yml` file contains the following configuration of masks:

* A randomChoice mask for the field `name`, choosing between Mickael, Toto and Benjamin.
* A regex mask for the field `surname`, respecting the regex "(Mar|Ber)t(rand|in)". This command is use to show an exemple of regex mask.
* A randomChoiceInUri mask for the field `identity.name`. This mask is used the same way a randomChoice mask does but allows to use an external list. All usable lists are shown in the global README.

#### **`masking.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "name"
    mask:
      randomChoice:
        - Mickael
        - Toto
        - Benjamin
  - selector:
      jsonpath: "surname"
    mask:
      regex: "(Mar|Ber)t(rand|in)"
  - selector:
      jsonpath: "identity.name"
    mask:
      randomChoiceInUri: "pimo://nameFR"
```

Be sure to be inside the demo3 folder and use the `pimo < data.jsonl > dataout.jsonl` command line. On the `dataout.jsonl` file, every data of every table should be masked thanks to the masks.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---

It is possible to edit the `masking.yml` file or the `data.jsonl` file to see how PIMO react in a case you're interested in.
