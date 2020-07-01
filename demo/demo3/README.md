# Third demo - Masking Tables of datas

This folder cointains the file for the third demo. This demo shows how PIMO masks datas contained in tables, you can see those different datas configurations in the `data.jsonl` file.

The `masking.yml` file contains the following configuration of masks:

* A randomChoice mask for the field `name`, choosing between Mickael, Toto and Benjamin.
* A regex mask for the field `surname`, respecting the regex "(Mar|Ber)t(rand|in)". This command is use to show an exemple of regex mask.
* A randomChoiceInUri mask for the field `identity.name`. This mask is used the same way a randomChoice mask does but allows to use an external list. All usable lists are shown in the global README.

Be sure to be inside the demo3 folder and use the `pimo < data.jsonl > dataout.jsonl` command line. On the `dataout.jsonl` file, every data of every table should be masked thanks to the mask

It is possible to change the `masking.yml` file or the `data.jsonl` file to see how Pimo react in a case you're interested in.
