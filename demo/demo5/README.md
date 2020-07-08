# Fifth demo - Multiple successive masking

This folder cointains the file for the fifth demo. This demo shows how to mask the same data multiple times. The `masking.yml` file contains the following configuration:

* A randDate mask for the `birth` field. This mask create a random date during the 2010's.
* A duration mask with "-P20Y" for the `birth` field. This mask will substract 20 year to the `birth` field.

The first mask will create a random date and the second will change this date. The result of those 2 successive masks should be a date during the 90's.

To use this demo, be sure to be inside the demo5 folder and use the following command line: `pimo < data.jsonl > dataout.jsonl`.
