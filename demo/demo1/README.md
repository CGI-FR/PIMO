# First demo - Simple use of PIMO

This folder contains the first demo of PIMO.

This demo shows how to use PIMO to mask simple data contained in a file. Data that will be used in this demo are contained in the `data.jsonl` file, the 5 lines from this file are identical ( `{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}` ) to see effects of the different masks.

#### **`data.jsonl`**
```json
{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}
{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}
{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}
{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}
{"age": 28, "name":"Dupont", "surname":"Martin", "town": "Nantes"}
```

The `masking.yml` file contains the following configuration:

* A constant mask to replace the `name` field by "Benjamin".
* A randomChoice mask to replace the `surname` field by either "Dupont" or "Dupond".
* A hash mask to replace the `town` field by either "Emerald City", "Ruby City" or "Sapphire City".
* A randomInt mask to replace `age` field by a random number contained between 18 and 90.

#### **`masking.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "name"
    mask:
      constant: "Benjamin"
  - selector:
      jsonpath: "surname"
    mask:
      randomChoice:
        - Dupont
        - Dupond
  - selector:
      jsonpath: "town"
    mask:
      hash:
        - "Emerald City"
        - "Ruby City"
        - "Sapphire City"
  - selector:
      jsonpath: "age"
    mask:
      randomInt:
        min: 18
        max: 90
```

To use it, make sure you are in the right folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---

Masked data are written in the new file `dataout.jsonl`. This exemple show the use of some masks (constant, hash, randomInt and randomChoice), but the `masking.yml` file can be changed to test other features (changing masks, removing or changing the seed, etc...). It is recommended to change the seed first (or to remove it) to see its effect. Launching PIMO multiple time with the same seed should keep the same results every time but removing the seed should change the results of every use of PIMO. After that you can change the town field on some of the input json to see the effects of the hash on different values.
