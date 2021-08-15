## Why build tags

When releasing application binaries to customers, if we need to control features
based on our plans like "Free, Pro, and Enterprise levels", build tags can help.

## go build

    go build filename.go
        It will create filename.exe
    go build
        It will create folder.exe, with all go files in that folder, build within it.

## build Tag Syntax

In first line of the go program, write in below synatx.

// +build tag_name

NOTE: ensure to leave an empty line after the build tag line.

when building the binary, pass the corresponding tag(s) name(s).

go build -tags TAGNAME

## build Tag boolean logic

We can have more than one build tags.

| Build Tag Syntax           | Build Tag Sample         | Boolean Statement  |
| -------------------------- | ------------------------ | ------------------ |
| Space-separated elements   | // +build pro enterprise | pro OR enterprise  |
| Comma-separated elements   | // +build pro,enterprise | pro AND enterprise |
| Exclamation point elements | // +build !pro           | NOT pro            |

Ref: https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags


project Usage
------------
```
~go build

~ls
a-buildtag-features-app.exe  enterprise.go  main.go  pro.go  README.md

~a-buildtag-features-app.exe
> Free Feature #1
> Free Feature #2

~go build -tags pro

~a-buildtag-features-app.exe
> Free Feature #1
> Free Feature #2
> Pro Feature #1 
> Pro Feature #2 

~go build -tags enterprise

~a-buildtag-features-app.exe
> Free Feature #1
> Free Feature #2

~go build -tags "pro enterprise"

~a-buildtag-features-app.exe     
> Free Feature #1
> Free Feature #2
> Enterprise Feature #1
> Enterprise Feature #2
> Pro Feature #1
> Pro Feature #2
```
