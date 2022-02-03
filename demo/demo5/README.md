# Fifth demo - Multiple successive masking

This folder contains files for the fifth demo. This demo shows how to create a pipeline of masks that will update the same data multiple times.

#### **`data.jsonl`**
```json
{"birth": ""}
{"birth": ""}
{"birth": ""}
{"birth": ""}
{"birth": ""}
```

The `masking.yml` file contains the following configuration:

* A randDate mask for the `birth` field. This mask create a random date during the 2010's.
* A duration mask with "-P20Y" for the `birth` field. This mask will substract 20 year to the `birth` field.

The first mask will create a random date and the second will change this date. The result of those 2 successive masks should be a date during the 90's.

[![](https://mermaid.ink/img/pako:eNp1kE1rwzAMhv-Kq1PKEnBy9KGXdbeNjW47jLkUNXYaQ2IHR2Ybbf_7HCeDMZgQlnj9CH2coXZKg4Cmcx91i57Y_U5aFm1l7BDoPUthvWdFsWFH46mdv8dwPHkc2lk7jKdZniwpE3-RsEOrtkg6m54HYwWreMkLXkZnnIvk7IZHY68vtzlLHH5OXPU_t5ZwWTqXfxofyqX1Nngk42xWPFX87VdFNVdoq-ZkUdOGKxco7rt5TGEPOfTa92hUvNF5wiVQq3stQcRU6QZDRxKkvUY0DCpOf6cMOQ-iwW7UOWAg9_xlaxDkg_6Btgbj-fqFun4De6J3NA)](https://mermaid.live/edit/#pako:eNp1kE1rwzAMhv-Kq1PKEnBy9KGXdbeNjW47jLkUNXYaQ2IHR2Ybbf_7HCeDMZgQlnj9CH2coXZKg4Cmcx91i57Y_U5aFm1l7BDoPUthvWdFsWFH46mdv8dwPHkc2lk7jKdZniwpE3-RsEOrtkg6m54HYwWreMkLXkZnnIvk7IZHY68vtzlLHH5OXPU_t5ZwWTqXfxofyqX1Nngk42xWPFX87VdFNVdoq-ZkUdOGKxco7rt5TGEPOfTa92hUvNF5wiVQq3stQcRU6QZDRxKkvUY0DCpOf6cMOQ-iwW7UOWAg9_xlaxDkg_6Btgbj-fqFun4De6J3NA)

#### **`masking.yml`**
```yaml
version: "1"
masking:
  - selector:
      jsonpath: "birth"
    mask:
      randDate:
        dateMin: "2010-01-01T00:00:00Z"
        dateMax: "2020-01-01T00:00:00Z"
  - selector:
      jsonpath: "birth"
    mask:
      duration: "-P20Y"
```

To use this demo, be sure to be inside the demo5 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

---
**NOTE**

All command lines are listed in [demo.sh](demo.sh).

---
