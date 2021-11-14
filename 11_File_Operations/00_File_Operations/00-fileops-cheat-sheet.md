# File Operations

## Basic Operations

    create empty file
                        newFile, err := os.Create("test.txt")
    Truncate a file
                        err := os.Truncate("test.txt", 100)
    Get file info
                        fileInfo, err := os.State("test.txt")
    Rename a file
                        err := os.Rename(oldPath, newPath)
    Delete a file
                        err := os.Remove("test.txt")
    Open a file for reading file,
                        err := os.Open("test.txt")
    close a file
                        err := file.Close()
    Change file permission
                        err := os.Chmod("test.txt", 0777)
    Change file ownership
                        err := os.Chown("test.txt", os.Getuid(), os.Getgid())
    Change file timestamps
                        err := os.Chtimes("test.txt", lastAccessTime, lastModifyTime)

## File open flag

    when opening file with os.OpenFile, flags control how the file behaves.

    os.O_RDONLY             open the file read only
    os.O_WRONGLY            open the file write only
    os.O_RDWR               open the file read write
    os.O_APPEND             append data to the file when writing
    os.O_CREATE             create a new file if none exists
    os.O_EXCL               used with O_CREATE, file must not exit
    os.O_SYNC               open for synchronous I/O
    O_TRUNC                 if possible, truncate file when opened

## Hard Link & Symbol Link

    create a hard link
                        err := os.Link("test.txt", "test_copy.txt")
    create a symbol link
                        err:= os.Symlink("test.txt", "test_sym.txt")
    Get link file info
                        fileInfo, err := os.Lstat("test_sym.txt")
    Change link file owner
                        err := os.Lchown("test_sym.txt", uid, gid)
    Read a link
                        dest, err := os.ReadLink("link_file.txt")

    ### hard link
        - creates a new pointer to the same place.
        - A file will only be deleted from disk after all links are removed.
        - Hard links only work on the same file system.
        - A hard link is what you might consider a 'normal' link.
    ## symbolic link, or soft link
        - only reference other files by name.
        - They can point to files on different filesystems.
        - Not all systems support symlinks.

## Read and Write

    Write bytes to file
                    n, err := file.Write([]byte("hello, world!\n"))
    Write string to file
                    n, err := file.WriteString("Hello, world!\n")
    Write at offset
                    n, err := file.WriteAt([]byte("Hello), 10)
    Read to byte
                    n, err := file.Read(byteSlice)
    Read exactly n bytes
                    n, err := io.ReadFull(file, byteSlice)
    Read at least n bytes
                    n, err := io.ReadAtLeast(file, byteSlice, minBytes)
    Read all bytes of a file
                    byteSlice, err := ioutil.ReadAll(file)
    Read from offset
                    n, err := file.ReadAt(byteSlice, 10)
