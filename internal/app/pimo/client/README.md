# PIMO Play client website

This folder should contains the PIMO Play website **before** compiling the `pimo` binary file.

This can be achieved in 2 ways :

- use the `neon build-web` command : this will build the static website and copy it into this directory
- use the `neon release` command : this will trigger the `neon build-web` then compile binay for production release

DO NOT COMMIT THE GENERATED STATIC WEBSITE
