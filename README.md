# Line Counter

This is a small command-line utility to count lines in text. That's it. That's all it does. Technically, it doesn't even care if it's text. 

There are some counting assumptions that I made. Basically, if it shows as a line in my text editor, I want to count it as a line. This means a zero-length file has a line count of 1. A blank line at the end of a file will count as 1. 

## How to Get

If you're new to Go, and especially Go on Windows, I've [written about what I've learned so far](https://hyr.mn/go-structure-windows/). 

You can download the source using `go get`. From the command line, run

```
> go get github.com/hyrmn/lc/cmd/lc
```

This will download the entire repository and place it within your `%GOPATH%\src` directory

Next, if you want, you can have Go compile the program and copy it to the `%GOPATH%\bin` directory

```
> go install github.com/hyrmn/lc/cmd/lc
```

Now you'll have an executable named `lc` (`lc.exe` on Windows).

The line counting logic is separated into a package if you want to use it in your own applications.

## How to Use

`lc` can either have information piped to it or it have a file path passed via the command line.

To read a file:

```
> lc "path/to/your/file.txt"
```

To read from stdin (information piped in):

```
> echo "Count the lines in this" | lc
```

## Runtime considerations

Using `wc` (Unix word count program) on my machine to parse a 1.6GB text file of lorem ipsum text, I get the following averages after an initial warmup call:

```
real    0m0.822s
user    0m0.156s
sys     0m0.655s
```

Using `lc` to parse the same file, I get the following averages after an initial warmup call:

```
real    0m0.625s
user    0m0.015s
sys     0m0.015s
```

I'm fairly happy with the performance and with getting to learn some Go

## Security Considerations

This program works by traversing a byte stream 32kb of data at a time. The lesser of 32kb of data _or_ the remainder of the stream is read into a buffer. If the final chunk of data read is less than 32kb then the buffer will contain both the data from the current read at the front of the buffer and the data from the previous read in the rest of the buffer.

The entire operation reuses the buffer for the duration of reading the stream. The buffer will stick around in memory with the final data in the buffer until it is garbage collected by Go.

All I/O operations are handled by Go's standard libraries. A buffered IO reader is used to read in Stdin. A file handle is used to read in from a file location.

## Caveats

I'm still learning Go and so some behaviors are either an issue with Windows or how I'm using the Go standard library. For example, running

```
echo "This is a line" | lc
```
returns an answer of `2`. 

This is because a carriage return terminates `bufio.NewReader(os.Stdin)` (at least on Windows). And, back to my line counting rules, this is the correct line count given the input; it just looks weird

Please note, due to how I've chosen to count lines, the output from `lc` may be one line higher than the output from `wc` (if you're hoping for exact compatibility, this behavior may surprise you)

## Where this project is going next

I think it needs to be containerized via Docker into a standalone image (because why not)