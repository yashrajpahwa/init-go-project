package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Enter the project name: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	project := scanner.Text()

	if project == "" {
		fmt.Println("Project name cannot be empty")
		return
	}

	projectWordsCharactersArray := strings.Split(project, " ")
	if len(projectWordsCharactersArray) > 1 {
		fmt.Println("Project name cannot contain spaces")
		return
	}

	mkdirCmd := exec.Command("mkdir", project)

	gopath := os.Getenv("GOPATH")

	if gopath == "" {
		fmt.Println("GOPATH not set")
		return
	}

	githubUsername := os.Getenv("GITHUB_USERNAME")

	if githubUsername == "" {
		fmt.Println("GITHUB_USERNAME not set")
		return
	}

	workingDir := os.Getenv("GOPATH") + "/src/github.com/" + githubUsername

	mkdirCmd.Dir = workingDir

	mkDirErr := mkdirCmd.Run()

	if mkDirErr != nil {
		fmt.Println(mkDirErr)
		fmt.Println("Error creating project (maybe it already exists)")
		return
	}

	goModInitCmd := exec.Command("go", "mod", "init")
	goModInitCmd.Dir = workingDir + "/" + project

	goModInitErr := goModInitCmd.Run()

	if goModInitErr != nil {
		fmt.Println(goModInitErr)
		fmt.Println("Error with go mod init")
		return
	}

	fmt.Println("Project created successfully")
}
