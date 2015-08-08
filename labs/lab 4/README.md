# Lab 4 – Batch File Processing

## Table of Contents

* [Description](#description)
* [References](#references)

## Description

In this lab you will modify the program created in [Lab 3](../lab%203) to add the additional following requirements:

* If the file passed as a command line argument is a ZIP file, all files contained in the archive should be processed.
* For all valid recycle bin files, Metadata should be displayed, or for invalid files, an appropriate error message.
* Files should be processed in parallel and results printed as they are available.

## References

* [zip - The Go Programming Language](https://golang.org/pkg/archive/zip/)
* [sync - The Go Programming Language](https://golang.org/pkg/sync/)

