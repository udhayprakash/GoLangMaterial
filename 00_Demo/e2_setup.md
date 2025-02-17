# Online execution

https://play.golang.org/
https://www.onlinegdb.com/online_go_compiler
https://onecompiler.com/go
https://goplay.space/
https://goplay.tools/

# IDEs

- vim: vim-go plugin provides Go programming language support
- Visual Studio Code: Go extension provides support for the Go programming language
- GoLand: GoLand is distributed either as a standalone IDE or as a plugin for IntelliJ IDEA Ultimate
- Atom: Go-Plus is an Atom package that provides enhanced Go support
- LiteIDE

## Ubuntu/Debian/Kali Linux:

    1) To Install Go

        Download Go:
            wget https://dl.google.com/go/go1.23.4.linux-amd64.tar.gz
        
        Extract the Go tarball:
            sudo tar -C /usr/local -xvf go1.23.4.linux-amd64.tar.gz

        Open the profile
            vim ~/.profile

        Add below Go Paths in that profile file
            export GOPATH=$HOME/go
            export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

        After saving the changes, run below command to load these changes:
            source ~/.profile

        To check go version
            go version

        Finally, clean the downloaded binary
            rm go1.23.4.linux-amd64.tar.gz


    2) To update Go
        Check existing go version:
            $ go version
        
        Remove existing go installation:
            $ sudo rm -r /usr/local/go/

        Download latest version:
            $ wget https://dl.google.com/go/go1.23.4.linux-amd64.tar.gz

        Extract the downloaded Archive:
            sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz

        Verify the Golang version
            $ go version
        
        Remove the Go tarball:
            $ rm go1.11.linux-amd64.tar.gz

## Windows/MAC Installation Procedure

    - Go to Download Link (https://golang.org/dl) and download the corresponding executable

## Environment variables Setup

    - With latest versios of Go, GOPATH setup is not needed.
    - Developers do not have to care about GOPATH anymore.
    - It has a default value (depending on OS)
    - GOPATH is still used by the Go toolchain internally, for caching downloaded modules and compilation artifacts.

##  using goplay - to run in go playground

    go install github.com/haya14busa/goplay/cmd/goplay@latest
    goplay -version
    
    (if doesnt work, need environment variable setup)
    export PATH=$PATH:$(go env GOPATH)/bin

    (configure proxy settings, if needed)
    export HTTP_PROXY=http://your-proxy-url:port
    export HTTPS_PROXY=http://your-proxy-url:port


    (To increase timeout)
    export GOPLAY_TIMEOUT=30


    To export, using command line 

        goplay upload yourfile.go


    curl -s -X POST --data-binary @test.go https://play.golang.org/share | awk -F/ '{print "https://play.golang.org/p/"$NF}'

    (To debug upload issue)
    strace goplay upload yourfile.go
