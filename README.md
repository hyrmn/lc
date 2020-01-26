# Line Counter

This is a small command-line utility to count lines in text. That's it. That's all it does. Technically, it doesn't even care if it's text. 

There are some counting assumptions that I made. I had originally chosen to have this match my editor's line count. That is, if Visual Studio Code shows `x` lines then my logic would also show `x` lines. However, I've chosen to follow the behavior of `wc -l`. I count carriage returns (`\n`). If a file does not end with a carriage return then the last line will not be counted.

While I'm not sure how I feel about this behavior, it is consistent with other tooling. A trailing carriage return is required to get an accurate count. Changing this is an exercise left to the reader.

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

The only output from `lc` will be the line count. This is because I want the ability to pipe this on to other programs easily.

So the full run might look like 

```
> lc "path/to/your/file.txt"
109
```

## Runtime Considerations

Using `time` (Unix timing utility) `wc` (Unix word count utility) on my machine to parse a 1.6GB text file of lorem ipsum text, I get the following averages after an initial warmup call:

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

Note: There is no timeout when waiting for piped input from stdin. If stdin never ends the stream then `lc` will hang until it is force-quit (`ctrl+c`).

## Security Considerations

This program works by traversing a byte stream 32kb of data at a time. The lesser of 32kb of data _or_ the remainder of the stream is read into a buffer. If the final chunk of data read is less than 32kb then the buffer will contain both the data from the current read at the front of the buffer and the data from the previous read in the rest of the buffer.

The entire operation reuses the buffer for the duration of reading the stream. The buffer will stick around in memory with the final data in the buffer until it is garbage collected by Go.

All I/O operations are handled by Go's standard libraries. A buffered IO reader is used to read in Stdin. A file handle is used to read in from a file location.

## Caveats

I'm still learning Go and so this utility is likely not idiomatic. It could probably be faster too!

The 32kb buffer size was chosen after profiling against several large files. While it's best on my poor little laptop, it may not be the most efficient on other platforms.

## Where this project is going next

I think it needs to be containerized via Docker into a standalone image (because why not)