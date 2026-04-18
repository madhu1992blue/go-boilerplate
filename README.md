# go-boilerplate

# Initialize new Go Module

```bash
go mod init github.com/madhu1992blue/go-boilerplate
```

This created the file "go.mod" in the current directory.
This contains:

- module github.com/madhu1992blue/go-boilerplate
    - The name of the module (github.com/madhu1992blue/go-boilerplate)
- go 1.26.2
    - The version of Go used (go 1.26.2)

# Setup Project Structure for Application Entrypoint : /cmd

- Create a directory called "cmd"
    - Each subdirectory in "cmd" will be a separate application entrypoint.
    - For example, if we have applications app1 and app2, we will have directories /cmd/app1 and /cmd/app2.
    - Each subdirectory in "cmd" will have a "main" package with a "main" function that will be the entrypoint for that application.
    - We added a demo application in /cmd/demo

# Let's ensure we can build

- We added a Makefile for building the binary
- make build-demo
- Also, we will exclude bin from version control.

