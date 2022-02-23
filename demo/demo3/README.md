# Third demo - Masking Tables of data

This folder contains files for the third demo. This demo shows how PIMO masks data contained in tables, you can see those different data configurations in the `data.jsonl` file.

#### **`data.jsonl`**
```json
{"name":["Benjamin", "Jean", "Louis"], "surname": ["Martin", "Bernard"]}
{"name":["Benjamin", "Jean", "Louis"], "surname": "Martin"}
{"name":"Benjamin", "surname": ["Martin", "Bernard"]}
{"name":"Benjamin", "surname": "Martin"}
{"identity":{"name":["Jeanne", "Louise"]}}
```

The `masking.yml` file contains the following configuration of masks:

* A randomChoice mask for the field `name`, choosing between Mickael, Jean and Benjamin.
* A regex mask for the field `surname`, respecting the regex "(Mar|Ber)t(rand|in)". This command is use to show an exemple of regex mask.
* A randomChoiceInUri mask for the field `identity.name`. This mask is used the same way a randomChoice mask does but allows to use an external list. All usable lists are shown in the global README.

[![](https://mermaid.ink/img/pako:eNqVkk1rwzAMhv9KqlMC6cquOfTQfUDHyiBjp6UMLVYSb7EdHJutNP3vc74GCdlhusi8eiU_Mj5DqhhBBFmpvtICtfEe40R6LlZcVta8-l0Kjt56vfUkCuqrtX3PNVZFJ73Vea-20QqtuUkgRsmUuCkUT8k_8PQTqQwfCGW4I_mBgssggaYfcd1PIMn6Qy92t66UNY5h-9Sl4994tdULhIM6gRy0kZNy-vb9A-pmRzowvnbgjaPr8Mb-OeGv_j9Izkgabk5XC6iT2gR4Ull43r18ifd-xYWKNpvWcx938NOJ8xVm1aVFIARBWiBn7pec27YETEEOHSJ3ZJShLU0Cibw4q60YGrpj3CgNUYZlTSGgNer5JFOIjLY0mm45uqXF4Lr8AG661yI)](https://mermaid.live/edit/#pako:eNqVkk1rwzAMhv9KqlMC6cquOfTQfUDHyiBjp6UMLVYSb7EdHJutNP3vc74GCdlhusi8eiU_Mj5DqhhBBFmpvtICtfEe40R6LlZcVta8-l0Kjt56vfUkCuqrtX3PNVZFJ73Vea-20QqtuUkgRsmUuCkUT8k_8PQTqQwfCGW4I_mBgssggaYfcd1PIMn6Qy92t66UNY5h-9Sl4994tdULhIM6gRy0kZNy-vb9A-pmRzowvnbgjaPr8Mb-OeGv_j9Izkgabk5XC6iT2gR4Ull43r18ifd-xYWKNpvWcx938NOJ8xVm1aVFIARBWiBn7pec27YETEEOHSJ3ZJShLU0Cibw4q60YGrpj3CgNUYZlTSGgNer5JFOIjLY0mm45uqXF4Lr8AG661yI)

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
        - Jean
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
