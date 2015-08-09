# Lab 2 – Windows Recycle Bin

## Table of Contents

* [Description](#description)
* [References](#references)

## Description 

The Recycle Bin contains "Recycled" files. Moving files and directories to the Recycle Bin is also referred to as soft deletion, since the files are not removed from the file system.

###### $RECYCLE.BIN
The $Recycle.Bin is used as of Windows Vista.

Per user Recycle Bin folder in the form:
```
C:\$Recycle.Bin\%SID%\
```

Which contains:
* $I files; "Recycled" file metadata
* $R files; the original data

In this lab you will create a new package called `file/recycle`, which contains the following definitions.  You must implement all variables, types, and methods as defined by the documentation below.

You can test your implementation by running `go test` while in the directory of your `file/recycle` package.

#### Variables

```go
var ErrInvalid = errors.New("Invalid recycle bin file.")
```
ErrInvalid is returned when the data passed to a Decode function is not properly
formatted.

#### type Metadata

```go
type Metadata struct {
	Name    string    // Original File Name
	Size    int64     // Original File Size
	Deleted time.Time // The Date And Time The File Was Deleted
}
```

Metadata gives information about a deleted file

#### func  DecodeI

```go
func DecodeI(r io.Reader) (Metadata, error)
```
DecodeI takes a Windows Vista+ $RECYCLE.BIN $I file as r and returns associated
Metadata for the deleted file. ErrInvalid is returned if r is not a valid $I
file.


## References

* [Windows - Forensics Wiki](http://forensicswiki.org/wiki/Windows#RECYCLER)
* [The Forensic Analysis of the Microsoft Windows Vista Recycle Bin](http://www.forensicfocus.com/downloads/forensic-analysis-vista-recycle-bin.pdf)
* [binary - The Go Programming Language](https://golang.org/pkg/encoding/binary/#Read)
