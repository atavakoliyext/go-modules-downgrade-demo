# Go Modules - Downgrade Demo

This repo demonstrates that order of `go get` calls for a series of modules will
result in different module graphs because of the behavior of [MVS Algorithm 4] --
i.e. the "Downgrade One Module" use-case.

This code uses the following 3 modules, 2 of which are not yet in `go.mod`:

- [github.com/russross/blackfriday/v2](https://github.com/russross/blackfriday)
- <github.com/Depado/bfchroma>
- <github.com/alecthomas/chroma>

A maintainer would like to add the latter two to `go.mod`. Given the features
used in this code, this maintainer would determine the minimum version
requirements for each are:

| Module @ Version                      | Reason                                        |
|---------------------------------------|-----------------------------------------------|
| [github.com/Depado/bfchroma@v1.3.0]   | New `EnableCSS()` function                    |
| [github.com/alecthomas/chroma@v0.6.1] | Go Module support + fixes to the VB.NET lexer |

[MVS Algorithm 4]: https://research.swtch.com/vgo-mvs#algorithm_4
[github.com/Depado/bfchroma@v1.3.0]: https://github.com/Depado/bfchroma/releases/tag/v1.3.0
[github.com/alecthomas/chroma@v0.6.1]: https://github.com/alecthomas/chroma/releases/tag/v0.6.1

Presented are 3 typical ways to get the desired versions. Of the 3, only one
(Scenario 2) results in the desired outcome.

## Scenario 1: Get both in one `go get` call

```shell
go get github.com/Depado/bfchroma@v1.3.0 github.com/alecthomas/chroma@v0.6.1
```

Fails with:

```
go get: github.com/Depado/bfchroma@v1.3.0 requires github.com/alecthomas/chroma@v0.7.3, not github.com/alecthomas/chroma@v0.6.1
```

## Scenario 2: Get chroma first, then bfchroma, then tidy


```shell
go get github.com/alecthomas/chroma@v0.6.1
go get github.com/Depado/bfchroma@v1.3.0
go mod tidy
```

Output:

```
go get: added github.com/alecthomas/chroma v0.6.1
go get: added github.com/Depado/bfchroma v1.3.0
go get: upgraded github.com/alecthomas/chroma v0.6.1 => v0.7.3
go get: added gopkg.in/russross/blackfriday.v2 v2.0.0
```

Diff of `go.mod`:

```diff
diff --git a/go.mod b/go.mod
index b5ceead..14b9b23 100644
--- a/go.mod
+++ b/go.mod
@@ -3,5 +3,7 @@ module github.com/atavakoliyext/go-modules-downgrade-demo
 go 1.16
 
 require (
+       github.com/Depado/bfchroma v1.3.0
+       github.com/alecthomas/chroma v0.7.3
        github.com/russross/blackfriday/v2 v2.0.1
 )
```

Tests pass:

```
$go test ./renderer
ok  	github.com/atavakoliyext/go-modules-downgrade-demo/renderer	1.132s
```

## Scenario 3: Get bfchroma first, then chroma, then tidy


```shell
go get github.com/Depado/bfchroma@v1.3.0
go get github.com/alecthomas/chroma@v0.6.1
go mod tidy
```

Output:

```
go get: added github.com/Depado/bfchroma v1.3.0
go get: downgraded github.com/Depado/bfchroma v1.3.0 => v1.2.0
go get: downgraded github.com/alecthomas/chroma v0.7.3 => v0.6.1
```

Diff of `go.mod`:

```diff
diff --git a/go.mod b/go.mod
index b5ceead..9e9bf64 100644
--- a/go.mod
+++ b/go.mod
@@ -3,5 +3,12 @@ module github.com/atavakoliyext/go-modules-downgrade-demo
 go 1.16
 
 require (
+       github.com/Depado/bfchroma v1.2.0
+       github.com/alecthomas/chroma v0.6.1
+       github.com/alecthomas/repr v0.0.0-20200325044227-4184120f674c // indirect
+       github.com/dlclark/regexp2 v1.2.0 // indirect
+       github.com/mattn/go-isatty v0.0.12 // indirect
        github.com/russross/blackfriday/v2 v2.0.1
+       github.com/stretchr/testify v1.6.1 // indirect
+       golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
 )
```

Tests fail:

```
$go test ./renderer
# github.com/atavakoliyext/go-modules-downgrade-demo/renderer [github.com/atavakoliyext/go-modules-downgrade-demo/renderer.test]
renderer/renderer.go:11:3: undefined: bfchroma.EmbedCSS
FAIL	github.com/atavakoliyext/go-modules-downgrade-demo/renderer [build failed]
FAIL
```
