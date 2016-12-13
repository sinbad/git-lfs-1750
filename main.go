package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Utility to re-create a file structure for https://github.com/git-lfs/git-lfs/issues/1750
// Provide a template directory which contains:
//    filestructure.txt - a dump from "find ." reporting the entire file structure
//    .gitattributes  } Files at any dir level which are copied into destination
//    .gitignore      }
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Required: template dir")
		fmt.Println(" git-lfs-1750 <templatedir> [destdir]")
		os.Exit(1)
	}

	templatedir := os.Args[1]
	stat, err := os.Stat(templatedir)
	checkError(err)
	if !stat.IsDir() {
		fmt.Println("Error:", templatedir, "is not a directory")
		os.Exit(3)
	}
	outputdir := "issue1750"
	if len(os.Args) > 2 {
		outputdir = os.Args[2]
	}
	fmt.Println("Target directory:", outputdir)

	checkError(os.MkdirAll(outputdir, 0755))

	filestructure := filepath.Join(templatedir, "filestructure.txt")
	f, err := os.OpenFile(filestructure, os.O_RDONLY, 0664)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	randReader := rand.New(rand.NewSource(77))
	numFiles := 0
	lastLine := ""
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore git repo and root
		if line == "." || strings.HasPrefix(line, "./.git") {
			continue
		}

		// To ignore dirs, write files 1 element behind, and skip writing any
		// entries which have sub-entries in the next line (because that's a dir)
		if len(lastLine) > 0 && !strings.HasPrefix(line, lastLine) {
			checkError(writeFile(filepath.Join(outputdir, filepath.Clean(lastLine)), randReader))

			numFiles++
		}

		lastLine = line

	}

	// Last entry is always a file
	checkError(writeFile(filepath.Join(outputdir, filepath.Clean(lastLine)), randReader))
	numFiles++

	fmt.Println("Created", numFiles, "files")

	// Now walk the template dir and copy .gitattributes and .gitignore files
	filepath.Walk(templatedir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if base == ".gitattributes" || base == ".gitignore" {
			relpath, err := filepath.Rel(templatedir, path)
			checkError(err)
			dest := filepath.Join(outputdir, relpath)

			copyFile(path, dest)
		}
		return nil
	})

	// Git init
	cmd := exec.Command("git", "init", outputdir)
	checkError(cmd.Run())
	fmt.Println("Created git repo")

	fmt.Println("Done")

}

func copyFile(src, dst string) {
	// copy file
	r, err := os.Open(src)
	checkError(err)
	defer r.Close()
	w, err := os.Create(dst)
	checkError(err)
	defer w.Close()
	_, err = io.Copy(w, r)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
		os.Exit(3)
	}
}

func writeFile(destfile string, source io.Reader) error {
	// Create file with a small amount of data, doesn't matter what
	err := os.MkdirAll(filepath.Dir(destfile), 0755)
	if err != nil {
		return err
	}
	of, err := os.OpenFile(destfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer of.Close()
	// 1-100 bytes
	_, err = io.CopyN(of, source, rand.Int63n(100)+1)
	return err
}
