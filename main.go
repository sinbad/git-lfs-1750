package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const gitattribs = `*.zip filter=lfs diff=lfs merge=lfs -text
*.xlsm filter=lfs diff=lfs merge=lfs -text
*.kml filter=lfs diff=lfs merge=lfs -text
`

// Utility to re-create a file structure for https://github.com/git-lfs/git-lfs/issues/1750
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Required: input file")
		fmt.Println(" git-lfs-1750 <contentsfile> [dirname]")
		os.Exit(1)
	}

	filename := os.Args[1]
	dir := "issue1750"
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}
	fmt.Println("Target directory:", dir)

	checkError(os.MkdirAll(dir, 0755))
	f, err := os.OpenFile(filename, os.O_RDONLY, 0664)
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
			checkError(writeFile(filepath.Join(dir, filepath.Clean(lastLine)), randReader))

			numFiles++
		}

		lastLine = line

	}

	// Last entry is always a file
	checkError(writeFile(filepath.Join(dir, filepath.Clean(lastLine)), randReader))
	numFiles++

	fmt.Println("Created", numFiles, "files")

	// Create gitattributes
	checkError(ioutil.WriteFile(filepath.Join(dir, ".gitattributes"), []byte(gitattribs), 0644))
	fmt.Println("Created .gitattributes")

	// Git init
	cmd := exec.Command("git", "init", dir)
	checkError(cmd.Run())
	fmt.Println("Created git repo")

	fmt.Println("Done")

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
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
