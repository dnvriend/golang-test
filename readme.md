# golang-test
A small study project on the [Go Programming Language](https://golang.org/) 

## Introduction
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. The latest version of 
the programming language and runtime as of Feb 2018 is v1.10. Go is available for Linux, Mac and Windows and compiles to native executables.
Go is a compiled, statically typed language in the tradition C, with garbage collection, limited structural typing, memory safety features, Garbage Collection, 
and CSP-style concurrent programming features added. The compiler and other language tools originally developed by Google are all free and open source.
The programming language uses the imperative, structural programming paradigm and has first class support for implicit typing and functions. The language 
is very concise and clean and provides low level primitives for generic data processing. The standard library provides support for all kinds of IO 
and provides abstractions like http and filesystem IO. There are a lot of third party libraries to process JSON, AVRO, Protobuf. There are also 
frameworks available like [Revel](https://revel.github.io/) - web framework, [Iris](https://iris-go.com/) - web framework, [GORM](http://gorm.io/) - an ORM Framework
and many more.

Go focuses on developer productivity as a team, so the programming language supports all basic construct many developers know from C or Java, but 
lacks support for higher abstractions like we know from Scala or Python, which makes the language a bit more verbose, but keeps the language understandable
for the whole team. Another pro is that Go solves problems in an idiomatic way, most of the time there is a single solution to a problem that allows for 
an understandable code base.

Although Go is a very new language, it is fast, has a small memory footprint and can compete with mature runtimes in speed and features, although sometimes
the JVM wins and sometimes Go, most features are on par with the JVM. 

In my opinion, where Go would really shine is in a serverless environment, because of the small memory footprint, fast start times, low GC overhead and very
small distribution file, it would make sense having a runtime that doesn't have a problem to stop and start every 5 minutes or so, launches very fast thus
leaving more time for processing.
  
## Workspace
The workspace contains the following directories

- src -> contains your source code, but also other people's source code
- pkg -> contains packages (libraries) that are compiled, ready to be linked to your executable
- bin -> will contain your executables

## Source directory structure
Most often, code resides in a repository, and github is a great place to store your code. Source code is organized using the following structure:

```
├── bin
│   └── map
├── environment-setup.sh
├── install.sh
├── pkg
├── readme.md
└── src
    └── github.com
        └── dnvriend
            ├── cards
            │   ├── deck.go
            │   ├── deck_test.go
            │   └── main.go
            ├── location
            │   └── main.go
            ├── map
            │   └── main.go
            ├── red
            │   └── main.go
            └── structs
                └── main.go
```

Create a directory of your source code repo eg. `github.com`, then create a dir for your repo username eg. `dnvriend` then place all your projects there. 
You can use either `git` to clone your project or use `go get` to get a repository that will be placed in this structure.

## Compiling a project
Binaries - executables - or packages - libraries - must be build per project basis, so when you need to build your executable, you must navigate to the 
directory where the `.go` file resides of your library or executable and type `go install`. A binary - executable - will be placed in the `bin` directory and 
a package - library - will be placed in the `pkg` directory. Multiple binaries can be build and all will be put in the `bin` directory. Also multiple packages 
can be build and will be put in the `pkg` directory. Please note, only after a package - library - is build, it can be imported in your project.

## Function visibility
By default, all functions in the same package have full visibility. Between different packages though, only functions starting with a capital letter are 
visible - exported - so `func toString` is not visible - not exported - but `func ToString` is visible.

## Resources
- [Golang Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language))

## Videos
- [The state of Go v1.10](https://www.youtube.com/watch?v=iR7LPAXWfmw) 
- [The state of Go v1.9](https://www.youtube.com/watch?v=vFJkH4qDjJ0)

 