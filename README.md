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

# Understand Build Related Environment Variables

| Variable | Default (2026) | Value (Container / CI/CD) | Value (Non-Container) | Importance |
|---|---|---|---|---|
| CGO_ENABLED | 1 (usually) | 0 | 1 (if using C libs) | High. 0 creates a static binary that runs on any Linux distro without external C library dependencies. |
| GOPROXY | proxy.golang.org | proxy.internal,direct | proxy.golang.org,direct | High. Directs where modules are fetched from. Use internal proxies in CI to avoid rate limits. |
| GOPATH | $HOME/go | /go | User Default | Medium. Now primarily stores the pkg/mod cache and installed bin files. |
| GOMODCACHE | $GOPATH/pkg/mod | Persisted Path | Default | Medium. In CI/CD, caching this folder is the single biggest factor in reducing build times. |
| GOCACHE | OS-dependent | Persisted Path | Default | Medium. Caches intermediate build results. Crucial for incremental builds in CI/CD. |

In our case, though we are not customizing GOPROXY, GOPATH, GOMODCACHE, GOCACHE - it uses the values from the Github Actions which does set the GOPATH, GOMODCACHE, GOCACHE for us.

uses: actions/setup-go@v6 has cache enabled by default. So, we are not customizing above. Otherwise, its important.

# Runtime Customization via Environment

These variables tune the Green Tea GC and the Go scheduler while your application is actually running.

| Variable | Meaning | Container Value | Bare Metal Value | Importance |
|---|---|---|---|---|
| GOMEMLIMIT | A "soft" cap on total runtime memory usage. | 90% of RAM Limit | Unset | Critical. Prevents "OOM-Kills" in Kubernetes/Docker by forcing GC to trigger before a crash. |
| GOGC | % of heap growth before GC triggers. | 100 (Default) | 100 | Medium. Green Tea GC is efficient enough that most apps no longer need to tweak this. |
| GOMAXPROCS | Number of threads for Go code. | Match CPU Limit | Unset (Auto-detect) | High. Ensures the Go scheduler doesn't thrash when running in restricted CPU environments. |
