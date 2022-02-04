# Fourth demo - Using other data in mask

This folder contains files for the fourth demo. This demo shows how PIMO masks data thanks to values contained in other fields by replacing or creating new values.

#### **`data.jsonl`**
```json
{"replacement": "Some better value", "fieldToReplace": "Some value"}
{"field1": "First field data", "field2": "Second field data", "compositeField": "Unimportant value"}
{"replacement": "Value1", "fieldToReplace": "Value0", "field1": "data1", "field2": "data2", "compositeField": "Unimportant value"}
{"field1": "AàEéè", "field2": "äâGêë", "compositeField": "Unimportant value"}
```

The `masking.yml` file contains the following configuration :

* A replacement mask that replace the `fieldToReplace` field with the `replacement` field.
* A template mask that replace the `compositeField` with the expression `{{.field1}}+{{.field2}}`, this expression will replace the original data with a created data containing the data from `field1` and `field2` fields separated by a + sign.

[![](https://mermaid.ink/img/pako:eNqtUj1PwzAQ_SvpTaloKyVjhk7AhITUdsMVMvalsYjtyLGFUOr_jvNRUCxQB3rLnfXevftyB0xzhALKWn-wihqbPO2ISoIthGqcfUkHtzwm6_U2KQXW_KB32NSU4d88pmWjW2HxsU-4opddwfMRb93bydCmipp4bU8j3tsc6iXOBKaXRGVT8xMvCZxjrakVVDyqOR9oVnMOTTUPKIOixbTrNuOQ3t9d4tz7oXgkmkVzZDfUyv-v9b2UeGfDpRba2XC37fPgjiMzFvqNCSuQaCQVPPzCrs8jYKtwIQJFCDk17wSI8oHnGh5af-DCagNFSesWV0Cd1ftPxaCwxuGFdC9oOJycWP4Lagn6Rg)](https://mermaid.live/edit/#pako:eNqtUj1PwzAQ_SvpTaloKyVjhk7AhITUdsMVMvalsYjtyLGFUOr_jvNRUCxQB3rLnfXevftyB0xzhALKWn-wihqbPO2ISoIthGqcfUkHtzwm6_U2KQXW_KB32NSU4d88pmWjW2HxsU-4opddwfMRb93bydCmipp4bU8j3tsc6iXOBKaXRGVT8xMvCZxjrakVVDyqOR9oVnMOTTUPKIOixbTrNuOQ3t9d4tz7oXgkmkVzZDfUyv-v9b2UeGfDpRba2XC37fPgjiMzFvqNCSuQaCQVPPzCrs8jYKtwIQJFCDk17wSI8oHnGh5af-DCagNFSesWV0Cd1ftPxaCwxuGFdC9oOJycWP4Lagn6Rg)

#### **`masking.yml`**
```yaml
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "fieldToReplace"
    mask:
      replacement: "replacement"
  - selector:
      jsonpath: "compositeField"
    mask:
      template: "{{.field1}}+{{.field2}}"
```

To use it, be sure to be inside the demo4 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---

The template mask can be used to remove lower case, upper case or accents on a text. Replacing the `template: "{{.field1}}+{{.field2}}"` line in the `masking.yml` by `template: "{{.field1 | ToLower | NoAccent}}+{{.field2 | ToUpper | NoAccent}}"` will remove every accent and format the text from field1 in lower case and from field2 in upper case.
