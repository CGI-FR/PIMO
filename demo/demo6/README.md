# Sixth demo - Adding and removing fields

This folder cointains the file for the fifth demo. This demo shows how to add or remove fields from input jsonlines. The `masking.yml` file contains the following configuration:

* A add mask that add the `newField` mask if it isn't already present.
* A remove mask that removes the `fieldToDelete` if present.

To try this demo, be sure to be inside the demo6 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

After using this command, every jsonline in the output should have a `newField` field with either the old value if the field was not present in the input or `newValue`
