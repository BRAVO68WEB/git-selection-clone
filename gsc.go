// Author: Jyotirmoy Bandyopadhayaya

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Clone a repository
func cloneRepo(gitURL string, branch string, output_folder string) {
	cmd := exec.Command("git", "clone", "-b", branch, gitURL, output_folder)
	cmd.Run()
}

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 3 {
		fmt.Println("Usage: gsc <git-url> <output-folder>")
		os.Exit(1)
	}

	// Arguments
	gitURL := os.Args[1]
	output_folder := os.Args[2]

	// Get current ts
	ts := time.Now().Unix()
	tempRepo := fmt.Sprintf("%d", ts)

	// Check if the output folder exists, throw an error if it does
	if _, err := os.Stat(output_folder); !os.IsNotExist(err) {
		fmt.Println("The output folder already exists")
		os.Exit(1)
	}

	// Create a temporary folder
	os.Mkdir(tempRepo, 0755)

	// Parse the git URL
	initialSplit := strings.Split(gitURL, "https://")[1]
	splitURL := strings.Split(initialSplit, "/")
	gitServerURL := "https://" + splitURL[0]
	username := splitURL[1]
	repoName := splitURL[2]
	branch := splitURL[4]
	folderPathElements := splitURL[5:]
	folderPath := strings.Join(folderPathElements, "/")

	// Clone the repository using the git command
	cloneUrl := gitServerURL + "/" + username + "/" + repoName + ".git"
	cloneRepo(cloneUrl, branch, tempRepo)

	// Copy the folder to the output folder
	cmd := exec.Command("cp", "-r", tempRepo+"/"+folderPath, output_folder)
	cmd.Run()

	// Remove the temporary folder
	cmd = exec.Command("rm", "-rf", tempRepo)
	cmd.Run()

	// Print the success message
	fmt.Println("Happy Hacking! 🙌")
	os.Exit(0)
}
