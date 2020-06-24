# A simple guide to learn to use PIMO

This folder contains every file necessary for demos of PIMO program.

## First demo - Simple use of PIMO

The folder demo1 contains the file for the first demo. This demo shows how to use PIMO to mask simple datas contained in the `data.jsonl` file, the 5 json from this file are identical to see effects of the different masks.

To use it, go inside the demo1 folder and use the following command line: `../../bin/pimo < data.jsonl > dataout.jsonl`

Masked datas are written in the new file `dataout.jsonl`. This exemple show the use of some masks (constant, hash, randomInt and randomChoice), but the `masking.yml` file can be changed to test other features (changing masks, removing or changing the seed, etc...).

## Second demo - Masking nested datas

The folder demo2 contains the file for the second demo. This demo shows how to use PIMO to mask nested datas contained in the `data.jsonl` file.

To use it, go inside the demo2 folder and use the following command line: `../../bin/pimo < data.jsonl > dataout.jsonl`

Tests can be done by changing the `masking.yml`

## Third demo - Masking Tables of datas

The folder demo3 cointains the file for the third demo. This demo shows how PIMO masks datas contained in tables, you can see different table configurations in the `data.jsonl` file.

To see the effect of Pimo on those datas, go inside the demo3 folder and use the `../../bin/pimo < data.jsonl > dataout.jsonl` command line.

The program should mask every field of the different tables. It is possible to change the `masking.yml` file or the `data.jsonl` file to see how Pimo react with those.
