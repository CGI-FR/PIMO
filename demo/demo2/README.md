# Second demo - Masking nested datas

This folder contains the second demo of PIMO.

This demo shows how to use PIMO to mask datas in nested json. Those datas are contained in the `data.jsonl` file. They are composed of the following fields : a field `town`, and a field `identity` that contains 3 fields : `age`, `name` and `surname`. The `masking.yml` file contains the following configuration :

* A command mask with the command `echo Dorothy` for the `identity.name` field. This command always returns Dorothy, so every name should be replaced by Dorothy.
* A weightedChoice mask for the field `identity.surname`. This should return Dupont in three times out of for and Dupond 1 time out of four.
* A randomInt mask for the field `identity.age`.
* A randomChoice mask for the `town` field that choose between Paris, Nantes and Carquefou.

To use it, be sure to be inside the demo2 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

Tests of different masks can be done by changing the `masking.yml`. If you have `chance.js` installed (<https://github.com/chancejs/chance-cli>) you can replace the ligne `command: "echo Dorothy"` by `command: "chance first --nationality=fr"` to replace the name field by a random name.
