package main

import (
	"fmt"
	"time"
	"io"
	"net/http"
	"os"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	file := "hosts"
	f, err := os.Create(file)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return file, n, err
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func main() {
	// Download file
	url := "http://someonewhocares.org/hosts/zero/hosts"
	filename, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
		return
	}
	fmt.Fprintf(os.Stdout, "%s => %s (%d bytes).\n", url, filename, n)

	// Backup old file
	backup_file := fmt.Sprintf("/etc/hosts.bak.%s", time.Now().Format("2006-01-02_15-04-05"))
	err = copyFileContents("/etc/hosts", backup_file)
	if err != nil {
		fmt.Println("Failed", err)
		return
	}
	fmt.Println("Backup written to", backup_file)

	// Copy file to /etc/hosts
	err = copyFileContents("hosts", "/etc/hosts")
	if err != nil {
		fmt.Println("Failed", err)
		return
	}

	fmt.Println("Success!")
}
