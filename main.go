package main

import (
	"github.com/champ-oss/file-sync/pkg/common"
	"github.com/champ-oss/file-sync/pkg/config"
	"github.com/champ-oss/file-sync/pkg/git/cli"
	"github.com/champ-oss/file-sync/pkg/github"
	log "github.com/sirupsen/logrus"
	"os"
)

const lockFile = ".terraform.lock.hcl"

func main() {
	log.SetLevel(log.DebugLevel)

	token := config.GetToken()
	workspace := config.GetWorkspace()
	terraformDir := os.Getenv("WORKING_DIRECTORY")
	repoName := config.GetRepoName()
	ownerName := config.GetOwnerName()
	targetBranch := config.GetTargetBranch()
	pullRequestBranch := config.GetPullRequestBranch()
	user := config.GetUser()
	email := config.GetEmail()
	commitMsg := config.GetCommitMessage()

	err := cli.SetAuthor(workspace, user, email)
	if err != nil {
		panic(err)
	}

	err = cli.Fetch(workspace)
	if err != nil {
		panic(err)
	}

	err = cli.Branch(workspace, pullRequestBranch)
	if err != nil {
		panic(err)
	}

	err = cli.Checkout(workspace, pullRequestBranch)
	if err != nil {
		panic(err)
	}

	err = cli.Reset(workspace, pullRequestBranch)
	if err != nil {
		panic(err)
	}

	_, err = common.RunCommand(terraformDir, "terraform", "init", "-upgrade", "-backend=false")
	if err != nil {
		panic(err)
	}

	if modified := cli.AnyModified(terraformDir, []string{lockFile}); !modified {
		log.Info("terraform lockfile is up to date")
	} else {
		err = cli.Add(terraformDir, lockFile)
		if err != nil {
			panic(err)
		}

		err = cli.Commit(workspace, commitMsg)
		if err != nil {
			log.Fatal(err)
		}

		err = cli.Push(workspace, pullRequestBranch)
		if err != nil {
			log.Fatal(err)
		}
	}

	client := github.GetClient(token)
	err = github.CreatePullRequest(client, ownerName, repoName, "Update Terraform Lockfile", pullRequestBranch, targetBranch)
	if err != nil {
		log.Fatal(err)
	}
}
