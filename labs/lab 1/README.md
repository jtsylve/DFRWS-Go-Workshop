# Lab 1 – Windows FILETIME

## Table of Contents

* [Description](#description)
* [References](#references)

## Description

In Windows, a `FILETIME` is a 64-bit value that represents the number of 100-nanosecond intervals that have elapsed since 12:00 A.M. January 1, 1601 Coordinated Universal Time (UTC).

In Go, a `time.Time` represents an instant in time with nanosecond precision.

In this lab you will create a new package called `timestamp`, which contains the following defined `FileTime` type.  You must implement all methods as defined by the documentation below.

#### type FileTime

```go
type FileTime int64
```

FileTime is a 64-bit value that represents the number of 100-nanosecond
intervals that have elapsed since 12:00 A.M. January 1, 1601 Coordinated
Universal Time (UTC).

#### func (FileTime) Time

```go
func (t FileTime) Time() time.Time
```
Time returns t as a standard time.Time

#### func  ToFileTime

```go
func ToFileTime(t time.Time) FileTime
```
ToFileTime returns t as a FileTime

## References

* [File Times (Windows) - MSDN - Microsoft](https://msdn.microsoft.com/en-us/library/windows/desktop/ms724290(v=vs.85).aspx)
* [time - The Go Programming Language](https://golang.org/pkg/time/)

