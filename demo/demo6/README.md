# Sixth demo - Adding and removing fields

This folder contains files for the fifth demo. This demo shows how to add or remove fields from input jsonlines.

In this example, input data are not clean, every line has a different structure, there are unimportant fields on some lines and missing fields on other.

#### **`data.jsonl`**
```json
{}
{"fieldToRemove":"Unimportant Value"}
{"newField":"Some Value"}
{"fieldToRemove":"Unimportant Value","newField":"Some Value"}
```

The `masking.yml` file contains the following configuration:

* A add mask that add the `newField` mask if it isn't already present.
* A remove mask that removes the `fieldToDelete` if present.

This configuration will clean the input stream into a correct one.

#### **`masking.yml`**
```yaml
version: "1"
masking:
  - selector:
      jsonpath: "newField"
    mask:
      add: "New Value"
  - selector:
      jsonpath: "fieldToRemove"
    mask:
      remove: true
```

To try this demo, be sure to be inside the demo6 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---

After using this command, every jsonline in the output should have a `newField` field with either the old value if the field was present in the input or `newValue`.
