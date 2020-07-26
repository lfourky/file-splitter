# file-splitter
A small Go utility application for splitting a large file (think gigabytes) into smaller ones.

## Available flags
```
-b uint
    Customize buffer size. (default 10000000)
-d string
    Provide a directory path where partial files will be stored (default "parts")
-f string
    Provide a relative path to file you wish to split.
-l uint
    Line limit per part-file. (default 10)
-p string
    Provide a prefix for each part. (default "part_")
-s string
    Provide a suffix for each part.
-crlf bool
    Use CRLF instead of LF for line endings.
```

## Running the program (example)
Let's say you have a file ```bigfile.sql``` of, let's say, ```100gb``` in size. It contains queries you might want to execute, but there might be an error in one of the lines. 
You try to open it in Notepad to see what's up, but the thing won't open. It won't load such a big file.
Damn it! Happens every time, right? Well, how about splitting it into thousands (or more) (or less?) of files and examining each one of them?

In your particular case, this command will solve your 100th life problem, so you're left with just 99.

```
go run main.go -f bigfile.sql
```

Now you see that a new folder ```parts``` is created in the directory where your ```bigfile.sql``` is.
It starts getting populated with files such as ```part_1.sql```, ```part_2.sql```, etc...

Congratulations! Now you've got 99 problems but a ```bigfile.sql``` ain't one. 
