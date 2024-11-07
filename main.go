package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func scriptFileToLastCommit(repoPath string, targetCommit *object.Commit, filePath string) string {

	cmd := exec.Command("git", "log", targetCommit.Hash.String(), "--follow", "--", filePath)
	cmd.Dir = repoPath

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalf("Failed to execute %s", cmd.String())
	}

	//fmt.Println(string(stdout))

	hashs := GetCommitHashes(string(stdout))

	var hashToUse string

	for _, hash := range hashs {
		if hash != targetCommit.Hash.String() {
			hashToUse = hash
			break
		}
	}

	if len(hashToUse) == 0 {
		log.Fatalf("Cannot find commit id from %s", string(stdout))
	}

	return fmt.Sprintf("git restore --source=%s -- %s", hashToUse, filePath)
}

func listFiles(repo *git.Repository, commit *object.Commit) []string {

	parents := commit.ParentHashes
	numParents := len(parents)
	fmt.Printf("Number of Parent Commits: %d\n", numParents)

	var parentCommits []*object.Commit

	for _, parentHash := range parents {
		parentCommit, err := repo.CommitObject(parentHash)
		if err != nil {
			log.Fatalf("Failed to get parent commit %s: %s", parentHash, err)
		}
		parentCommits = append(parentCommits, parentCommit)
		fmt.Printf("Parent Commit: %s\n", parentCommit.Hash)
	}

	filePaths := make([]string, 0)

	if numParents > 0 {
		fmt.Println("Comparing first parent with the merge commit:")
		tree1, err := parentCommits[0].Tree()
		if err != nil {
			log.Fatalf("Failed to get tree for first parent: %s", err)
		}

		tree2, err := commit.Tree()
		if err != nil {
			log.Fatalf("Failed to get tree for the merge commit: %s", err)
		}

		diff, err := tree1.Diff(tree2)
		if err != nil {
			log.Fatalf("Failed to get diff between first parent and merge commit: %s", err)
		}

		fmt.Println("Changed files between first parent and merge commit:")

		for _, change := range diff {
			fmt.Println(change.To.Name)
			filePaths = append(filePaths, change.To.Name)
		}
	}

	return filePaths
}

func regret(repoPath string, commitId string) {

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Failed to open repository: %s", err)
	}

	commit, err := repo.CommitObject(plumbing.NewHash(commitId))
	if err != nil {
		log.Fatalf("Failed to get commit: %s", err)
	}

	fmt.Printf("Commit: %s\n", commit.Hash)
	fmt.Printf("Author: %s\n", commit.Author)
	fmt.Printf("Date: %s\n", commit.Committer.When)
	fmt.Println("Message:")
	fmt.Println(commit.Message)

	filePaths := listFiles(repo, commit)

	allScripts := make([]string, 0)

	for _, filePath := range filePaths {

		aScript := scriptFileToLastCommit(repoPath, commit, filePath)
		allScripts = append(allScripts, aScript)
	}

	fmt.Println("====================================")
	fmt.Println("Scripts to restore:")
	for _, script := range allScripts {
		fmt.Println(script)
	}

}

func main() {
	//Change here
	gitPath := "/YOUR_LOCAL_REPO_PATH"

	//Change here, the commit id you want to revert
	commitId := "8fb011a7dba0651f6cfa5882febb3376cc46eed6"

	regret(gitPath, commitId)
}
