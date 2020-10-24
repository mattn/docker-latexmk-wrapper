package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := strings.Join([]string{
		`C:\Program Files\Docker\Docker\resources\bin`,
		os.Getenv("PATH"),
	}, string(os.PathListSeparator))

	args := []string{"latexmk"}
	for _, arg := range os.Args[1:] {
		args = append(args, fmt.Sprintf("%s", filepath.ToSlash(arg)))
	}

	cmd := exec.Command(
		`C:\Program Files\Docker\Docker\resources\docker.exe`,
		"run",
		"--rm",
		"-t",
		"--workdir=/root",
		"--volume",
		fmt.Sprintf("%s:/root", filepath.ToSlash(cwd)),
		"eisoku9618/latex:latex-japanese",
		"/bin/bash",
		"-c",
		strings.Join(args, " "),
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PATH=%s", path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
