# Second demo - Masking nested data

This folder contains the second demo of PIMO.

This demo shows how to use PIMO to mask data in nested json. Those data are contained in the `data.jsonl` file. They are composed of the following fields : a field `town`, and a field `identity` that contains 3 fields : `age`, `name` and `surname`.

#### **`data.jsonl`**
```json
{"town": "Nantes", "identity": {"age": 27, "name":"Jean", "surname": "Bertrand"}}
{"town": "Nantes", "identity": {"age": 58, "name":"Benjamin", "surname": "Martin"}}
{"town": "Rennes", "identity": {"age": 35, "name":"Benjamin", "surname": "Bertrand"}}
{"town": "Rennes", "identity": {"age": 28, "name":"Jean", "surname": "Martin"}}
```

The `masking.yml` file contains the following configuration :

* A command mask with the command `echo Dorothy` for the `identity.name` field. This command always returns Dorothy, so every name should be replaced by Dorothy.
* A weightedChoice mask for the field `identity.surname`. This should return Dupont in three times out of for and Dupond 1 time out of four.
* A randomInt mask for the field `identity.age`.
* A randomChoice mask for the `town` field that choose between Paris, Nantes and Carquefou.

[![](https://mermaid.ink/img/pako:eNqdk09Lw0AQxb_Kdk8ppGLxojkUofUgWJV68GBKGbOTZKHZjZtZamn73d1sUiGhFXUvO8x7PN4vf3Y80QJ5xNO13iQ5GGIPi1gxdwZSlZbeAn8Nl2w0mjApUJGk7YWCAhtbZd8zA2Xe1VZV1sj16Sh1zj7mU10UoESASa7ZTBtN-XYY830vZtykoBLN0FN9qYG25CpOnvy1_EX7ypqfAFr5NEMrthivKLOcUExzLRMMZrbUitgtuwr9KNw47mIdw8-SfRv-CQfZWTAnnYZyQgu0cC9FF_eKgrlUERtfh2wOnxG7uexi1FFnEbz4t_qkN6pXu1516taLTs32qT-DkVX4CIqwCqdgPiym2vq-PqPfs1me6sdDXqApQAr3S-xqd8wpR_ep8MiNAlOwa4p5rA7OaksBhHdCkjY8SmFdYcjBkn7ZqoRHZCweTTMJDqpoXYcvYuQoOw)](https://mermaid.live/edit/#pako:eNqdk09Lw0AQxb_Kdk8ppGLxojkUofUgWJV68GBKGbOTZKHZjZtZamn73d1sUiGhFXUvO8x7PN4vf3Y80QJ5xNO13iQ5GGIPi1gxdwZSlZbeAn8Nl2w0mjApUJGk7YWCAhtbZd8zA2Xe1VZV1sj16Sh1zj7mU10UoESASa7ZTBtN-XYY830vZtykoBLN0FN9qYG25CpOnvy1_EX7ypqfAFr5NEMrthivKLOcUExzLRMMZrbUitgtuwr9KNw47mIdw8-SfRv-CQfZWTAnnYZyQgu0cC9FF_eKgrlUERtfh2wOnxG7uexi1FFnEbz4t_qkN6pXu1516taLTs32qT-DkVX4CIqwCqdgPiym2vq-PqPfs1me6sdDXqApQAr3S-xqd8wpR_ep8MiNAlOwa4p5rA7OaksBhHdCkjY8SmFdYcjBkn7ZqoRHZCweTTMJDqpoXYcvYuQoOw)
#### **`masking.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "identity.name"
    mask:
      command: "echo Dorothy"
  - selector:
      jsonpath: "identity.surname"
    mask:
      weightedChoice:
        - choice: "Dupont"
          weight: 3
        - choice: "Dupond"
          weight: 1
  - selector:
      jsonpath: "identity.age"
    mask:
      randomInt:
        min: 18
        max: 90
  - selector:
      jsonpath: "town"
    mask:
      randomChoice:
        - "Paris"
        - "Nantes"
        - "Carquefou"
```

To use it, be sure to be inside the demo2 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---

Tests of different masks can be done by editing the `masking.yml`. If you have `chance.js` installed (<https://github.com/chancejs/chance-cli>) you can replace the ligne `command: "echo Dorothy"` by `command: "chance first --nationality=fr"` to replace the name field by a random name.
