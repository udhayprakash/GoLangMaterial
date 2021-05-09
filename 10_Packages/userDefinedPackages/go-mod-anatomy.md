# Dependency Management

    go get -d github.com/path/to/module # add or upgrade dep
    go get -d github.com/dep/two/v2@v2.1.0 # use specific version
    go get -d github.com/dep/commit@branch # use specific branch
    go get -d -u ./... # upgrade all modules used in subdirs

    go get -d github.com/dep/legacy@none # remove dep

# Useful commands

    go mod tidy # organize and clean up go.mod and go.sum
    go mod download # download deps into module cache
    go mod init github.com/path/to/module # initialize new module
    go mod why -m github.com/path/to/module # why is the module a dependency?
    go install github.com/path/to/bin@latest # build and install a binary

ref: https://encore.dev/guide/go.mod
