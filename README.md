# Pimo : Private Input, Masked Output

PIMO is a tool for data masking. It can mask datas from JSONline and return other JSONline thanks to a masking configuration contained in a yaml file.

## Configuration file needed

Pimo requires a yaml configuration file to works. This file must be named `masking.yml` and be placed in the working directory. The file must respect the following format :

```
version: "1"
seed: 42
masking:
  - selector:
      jsonpath: "exemple.exemple"
    mask:
      type: "argument"
```

`version` is the version of the masking file.  
`seed` is to give every random mask the same seed, it is optional and in case it is not informed, the seed is chosen with the time to look more random.  
`masking` is used to define the masks that are going to be used by pimo.  
`selector` is made of a jsonpath and a mask.  
`jsonpath` defines the path of the entry that has to be masked in the json file.  
`mask` defines the mask that will be used for the entry defined by `selector`.  

## Possible masks

The following types of masks can be used : 

* `regex` is to mask using a regular expression given in argument.
* `constant` is to mask the value by a constant value given in argument.
* `randomChoice` is to mask with a random value from a list in argument.
* `randomInt` is to mask with a random int from a range with arguments min and max.
* `command` is to mask with the output of a console command given in argument.
* `weightedChoice` is to mask with a random value from a list with probability, both given with the arguments `choice` and `weight`.
* `hash` is to mask with a value from a list by mashing the original value, allowing to mask a value the same way every time.
* `randomDate` is to mask a date with a random date between `dateMin` and `dateMax`.

A full `masking.yml` file exemple, using every kind of mask, is given with the source code.

In case two types of mask are entered in the same selector, the program can't extract the masking configuration and will return an error. The file `wrongMasking.yml` provided with the source illustrate that error.

## Usage
To use PIMO to mask a `data.json`, use in the following way : 
```
./pimo <data.json >maskedData.json
```

This takes the `data.json` file, masks the datas contained inside it and put the result in a `maskedData.json` file.