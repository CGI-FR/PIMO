# Fourth demo - Using other datas in mask

This folder cointains the file for the fourth demo. This demo shows how PIMO masks datas thanks to datas contained in other fields by replacing or creating other datas.

The `masking.yml` file contains the following configuration :

* A replacement mask that replace the `fieldToReplace` field with the `replacement` field.
* A template mask that replace the `compositeField` with the expression "{{.field1}}+{{.field2}}", this expression will replace the original data with a created data containing the datas from `field1` and `field2` fields separated by a + sign.

To use it, be sure to be inside the demo4 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.

The template mask can be used to remove lower case, upper case or accents on a text. Replacing the `template: "{{.field1}}+{{.field2}}"` line in the `masking.yml` by `template: "{{.field1 | ToLower | NoAccent}}+{{.field2 | ToUpper | NoAccent}}"` will remove every accent and format the text from field1 in lower case and from field2 in upper case.
