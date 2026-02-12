# YAML Merge Tags & Parser Differential Research

This repository contains tooling and test cases used in the research described in the accompanying blog post:

👉 **YAML Merge Tags and Parser Differentials**
[https://blog.darkforge.io/yaml/merge/parser/differential/research/2026/02/11/YAML-Merge-Tags-and-Parser-Differentials.html](https://blog.darkforge.io/yaml/merge/parser/differential/research/2026/02/11/YAML-Merge-Tags-and-Parser-Differentials.html)

The goal of this project is to explore **parser differentials** across popular YAML implementations—specifically how merge keys (`<<`) and duplicate keys are handled differently by various language parsers.

---

## Repository Overview

This repo includes:

* Multiple YAML parsers (Go, Node.js, Ruby, Python, etc.)
* A shared test harness to run identical inputs across parsers
* Output comparison to highlight behavioural differences

The parsers are orchestrated via the `all_parsers.sh` script, which expects each parser binary or script to live in the project root.

---

## Building the Go Parser

Before testing the Go and Node parsers, you must first build the Go parser and place the resulting binary in the root directory so it can be picked up by `all_parsers.sh`.

```bash
cd goparse
go build .
cp goparse ../
```

This will produce a `goparse` binary in the project root.

---

## Setting Up the Node.js Parser

The Node.js parser is implemented in `nodeparse.js` and requires dependencies to be installed first.

## Running the Tests

Place the test YAML in data.yaml. Once all parsers are built and installed, you can run the full test suite using:

```bash
./all_parsers.sh
```

## Disclaimer

This project is intended for **research and educational purposes only**.
Do not use these techniques against systems you do not own or have permission to test.


