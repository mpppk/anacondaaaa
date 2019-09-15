# golang CLI Template
[![CircleCI](https://circleci.com/gh/mpppk/anacondaaaa.svg?style=svg)](https://circleci.com/gh/mpppk/anacondaaaa)
[![Build status](https://ci.appveyor.com/api/projects/status/qv1fyq6fm8ni4cne?svg=true)](https://ci.appveyor.com/project/mpppk/anacondaaaa)
[![codecov](https://codecov.io/gh/mpppk/anacondaaaa/branch/master/graph/badge.svg)](https://codecov.io/gh/mpppk/anacondaaaa)
[![GoDoc](https://godoc.org/github.com/mpppk/anacondaaaa?status.svg)](https://godoc.org/github.com/mpppk/anacondaaaa)

*golang project template for building CLI*

## Setup
### Setup by Command
1. `git clone https://github.com/mpppk/anacondaaaa your_awesome_tool`
1. Replace all strings `anacondaaaa` in this repository to `your_awesome_tool`

### Setup on GitHub
Click "Use this template" button on GitHub project page.

![](https://github.com/mpppk/anacondaaaa/wiki/images/template-button.png)

## Project structure
* `/cmd` includes golang files which implements command and sub commands.
* `/pkg` includes golang files which will be used by other projects.
* `/internal` includes golang files which will **not** be used by other projects.
* `/testdata` includes data files for tests. (see https://golang.org/cmd/go/#hdr-Test_packages)

For more detail, see [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## Add new sub command
If you want to create new sub command, Add new go file to `/cmd`.

For more details, see [spf13/cobra](https://github.com/spf13/cobra).

## Task runner
Use [Makefile](https://github.com/mpppk/anacondaaaa/blob/master/Makefile).

## Error handling
Use [golang.org/x/xerrors](https://godoc.org/golang.org/x/xerrors).

## Documentation
Write [godoc](https://blog.golang.org/godoc-documenting-go-code)([example code](https://github.com/mpppk/anacondaaaa/blob/master/pkg/sum/sum.go#L9))
 or *Example test*([example code](https://github.com/mpppk/anacondaaaa/blob/master/pkg/sum/sum_test.go#L13-L18https://github.com/mpppk/anacondaaaa/blob/master/pkg/sum/sum_test.go#L13-L18)).

## Testing
Don't write test in same package, instead put to `package-name_test` package.  
For example, test of [pkg/sum/sum.go](https://github.com/mpppk/anacondaaaa/blob/master/pkg/sum/sum_test.go) is in `sum_test` package, not `sum` package.  
To use unexported variables or functions in test, expose these by `export_test.go` file.  
(ex. [/internal/option/root_export_test.go](https://github.com/mpppk/anacondaaaa/blob/master/internal/option/root_export_test.go))

For more details, see [this article(Japanese)](https://tech.mercari.com/entry/2018/08/08/080000).

### cmd test
Recommended way is to wrap cobra.Command instance by func (unlike the code generated by cobra add).
For example, see [cmd/sum_test.go](https://github.com/mpppk/anacondaaaa/blob/master/cmd/sum_test.go).

### with file system
This template depends [spf13/afero](https://github.com/spf13/afero).  
`afero.OsFs` is used in packages and `afero.MemMapFs` is used in tests.
For example, see [cmd/sum_test.go#TestSumWithOutFile](https://github.com/mpppk/anacondaaaa/blob/master/cmd/sum_test.go)

## Auto release via Circle CI powered by [goreleaser](https://github.com/goreleaser/goreleaser)
Create version tag (e.g. v0.0.1) and push to GitHub.  
Then goreleaser will release the new version on Circle CI.  
(Before push tag, you must provide GitHub token to Circle CI as environment variable)

For more details, see [my article (Japanese)](https://qiita.com/mpppk/items/ab328356ca14938a1208).

## Build & Run Docker image

```bash
$ docker build -t anacondaaaa .
...
$ docker run anacondaaaa sum 1 2
3
```

## SaaS integration
### Circle CI
This template includes [.circleci/config.yml](https://github.com/mpppk/anacondaaaa/blob/master/.circleci/config.yml).

### AppVeyor
This template includes [appveyor.yml](https://github.com/mpppk/anacondaaaa/blob/master/appveyor.yml).

### CodeCov
Makefile includes [codecov task](https://github.com/mpppk/anacondaaaa/blob/master/Makefile) which send coverage to CodeCov.  
Circle CI also send coverage to CodeCov by its job.

### Renovate
This template includes [renovate.json](https://github.com/mpppk/anacondaaaa/blob/master/renovate.json).

## README template

--------

# anacondaaaa

## Installation

### MacOS

### Linux
Download from [GitHub Releases](https://github.com/mpppk/anacondaaaa/releases)

### Windows
Download from [GitHub Releases](https://github.com/mpppk/anacondaaaa/releases)

## Usage

*Write usage of your awesome tool here*

