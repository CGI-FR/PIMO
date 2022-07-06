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

[![](https://mermaid.ink/img/pako:eNpVUEtrwzAM_iuuTi20lF1zKAy2ncYGbdklCsOLlMbgR3DshdH2v8-J05XpIvl7WegMtSOGAhrthrqVPojXPVqRaiGJyu0j0bYSm81OWB5eFGvKbB-_Tl527R_82Z8yM9YNHI0XhBSyfONBfEgdeYVwubsesontnHsnpk8XLoYuht371Kp5M2XTo1xObZW3a0bT0e3ZuG_Osn9QTvPTXCLOOqxgDYa9kYrSEc6jESG0bBihSCNxI6MOCGivSRo7koGfSQXnoWik7nkNMgZ3-LE1FMFHvomelEwHMrPq-gvit3by)](https://mermaid.live/edit/#pako:eNpVUEtrwzAM_iuuTi20lF1zKAy2ncYGbdklCsOLlMbgR3DshdH2v8-J05XpIvl7WegMtSOGAhrthrqVPojXPVqRaiGJyu0j0bYSm81OWB5eFGvKbB-_Tl527R_82Z8yM9YNHI0XhBSyfONBfEgdeYVwubsesontnHsnpk8XLoYuht371Kp5M2XTo1xObZW3a0bT0e3ZuG_Osn9QTvPTXCLOOqxgDYa9kYrSEc6jESG0bBihSCNxI6MOCGivSRo7koGfSQXnoWik7nkNMgZ3-LE1FMFHvomelEwHMrPq-gvit3by)

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
