package main

import (
    "flag"
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "strings"
    "syscall"
    "time"
)

func messageForDirAndFile(isDir bool) string {
    if isDir {
        return "File folder"
    }
    return "Simple file"
}

func fileandfolderCount(path string) (int, int, error) {
    fi := 0
    fo := 0
    err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            fo++
        } else {
            fi++
        }
        return nil
    })
    return fi, fo, err
}

func DirSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}

// stat is to return the info of file and directory
func stat(path string) (fs.FileInfo, error) {
    fileInfo, err := os.Stat(path)
    return fileInfo, err
}

func main() {

    loc := flag.String("f", "", "provide full path of file/folder")
    flag.Parse()
    if strings.TrimSpace(*loc) == "" {
        fmt.Println("flag is empty")
        flag.Usage()
        os.Exit(1)
    }

    info, err := stat(*loc)
    if err != nil {
        fmt.Println("failed to get stat of loc", err)
        os.Exit(2)
    }

    var size int64
    var file, folder int
    if info.IsDir() {
        size, err = DirSize(*loc)
        if err != nil {
            fmt.Println("failed to get stat of loc", err)
            os.Exit(2)
        }

        file, folder, err = fileandfolderCount(*loc)
        if err != nil {
            fmt.Println("failed to count folder and files into directory", err)
            os.Exit(2)
        }
    } else {
        size = info.Size()
        file = 1
        folder = 0
    }

    name := info.Name()
    updatedAt := info.ModTime().Format(time.UnixDate)

    sys := info.Sys().(*syscall.Win32FileAttributeData)
    cTime := time.Unix(0, sys.CreationTime.Nanoseconds())
    createdAt := cTime.Format(time.UnixDate)

    aTime := time.Unix(0, sys.LastAccessTime.Nanoseconds())
    lastAccess := aTime.Format(time.UnixDate)

    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %-55s |\n", "Types of file", messageForDirAndFile(info.IsDir()))
    fmt.Printf("|%s|\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %-55s |\n", "Name", name)
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %s %-46s |\n", "Size", fmt.Sprintf("%d", size), "bytes")
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    if folder != 0 && file != 1 {
        fmt.Printf("| %-20s | %s Files, %s %-39s |\n", "Contains", fmt.Sprintf("%d", file), fmt.Sprintf("%d", folder-1), "Folders")
        fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    }
    fmt.Printf("| %-20s | %-55s |\n", "Permissions", info.Mode())
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %-55s |\n", "Created", createdAt)
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %-55s |\n", "Modified", updatedAt)
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))
    fmt.Printf("| %-20s | %-55s |\n", "Recently Accessed", lastAccess)
    fmt.Printf("+%s+\n", strings.Repeat("-", 80))

}



// $ go run k_tool.go -f "full/path/of/file/or/dir"