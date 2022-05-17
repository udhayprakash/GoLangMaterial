package main

import (
	"flag"
	"github.com/golang/glog" //glog
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	defer glog.Flush()
	flag.Parse()

	if strings.TrimSpace(flag.Lookup("log_dir").Value.String()) == "" {
		exe, _ := os.Executable()
		flag.Set("log_dir", path.Join(filepath.Dir(exe), "logs"))
	}
	flag.Set("alsologtostderr", "true")
	glog.Infof("log_dir: %s", flag.Lookup("log_dir").Value)

	glog.Infof("Info 日志: %s", "some info")
	glog.Warningf("Warning 日志: %s", "some warning")
	glog.Errorf("Error 日志: %s", "some error")
	// glog.Fatalf("Fatal 日志: %s", "some fatal")  // terminates the program

	// $ go build . && ./hello -log_dir="./logs" -v 4
	glog.V(1).Infof("V1 日志: %s", "some info")
	glog.V(2).Infof("V2 日志: %s", "some info")
	glog.V(3).Infof("V3 日志: %s", "some info")
	glog.V(4).Infof("V4 日志: %s", "some info")
}

/*
$ ls logs
hello.ERROR
hello.INFO
hello.Mac-mini.ljh.log.ERROR.20220402-231029.9073
hello.Mac-mini.ljh.log.INFO.20220402-231029.9073
hello.Mac-mini.ljh.log.WARNING.20220402-231029.9073
hello.WARNING
$
$ cat logs/hello.INFO
Log file created at: 2022/04/02 23:10:29
Running on machine: Mac-mini
Binary: Built with gc go1.18 for darwin/amd64
Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
I0402 23:10:29.729131    9073 hello.go:22] log_dir: /Users/ljh/Documents/helloGo/logs
I0402 23:10:29.729662    9073 hello.go:24] Info 日志: some info
W0402 23:10:29.729674    9073 hello.go:25] Warning 日志: some warning
E0402 23:10:29.729936    9073 hello.go:26] Error 日志: some error
I0402 23:10:29.730170    9073 hello.go:30] V1 日志: some info
I0402 23:10:29.730180    9073 hello.go:31] V2 日志: some info
I0402 23:10:29.730185    9073 hello.go:32] V3 日志: some info
I0402 23:10:29.730190    9073 hello.go:33] V4 日志: some info
$
*/