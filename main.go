package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/champ-oss/file-sync/pkg/common"
	"github.com/champ-oss/file-sync/pkg/config"
	"github.com/champ-oss/file-sync/pkg/git/cli"
	"github.com/champ-oss/file-sync/pkg/github"

	log "github.com/sirupsen/logrus"
)

const lockFile = "**/.terraform.lock.hcl"

func main() {
	log.SetLevel(log.DebugLevel)

	token := config.GetToken()
	workspace := config.GetWorkspace()
	terraformDir := os.Getenv("WORKING_DIRECTORY")
	terraCmd := "terragrunt"
	repoName := config.GetRepoName()
	ownerName := config.GetOwnerName()
	targetBranch := config.GetTargetBranch()
	pullRequestBranch := config.GetPullRequestBranch()
	user := config.GetUser()
	email := config.GetEmail()
	commitMsg := config.GetCommitMessage()

	err := cli.SetAuthor(workspace, user, email)
	if err != nil {
		log.Panic(err)
	}

	err = cli.Fetch(workspace)
	if err != nil {
		log.Panic(err)
	}

	err = cli.Branch(workspace, pullRequestBranch)
	if err != nil {
		log.Panic(err)
	}

	err = cli.Checkout(workspace, pullRequestBranch)
	if err != nil {
		log.Panic(err)
	}

	err = cli.Reset(workspace, pullRequestBranch)
	if err != nil {
		log.Panic(err)
	}

	err = filepath.WalkDir(terraformDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && !strings.Contains(path, ".terragrunt-cache") {
			log.Info("Checking for terragrunt.hcl in ", path)
			if _, err = os.Stat(path + "/terragrunt.hcl"); err == nil {
				log.Info("terragrunt.hcl found in ", path)
				data, err := os.ReadFile(path + "/terragrunt.hcl")

				if err != nil {
					log.Panic(err)
				}

				if strings.Contains(string(data), "source = ") {
					log.Info("Updating terraform providers in ", path)
					_, err = common.RunCommand(path, terraCmd, "init", "-upgrade", "-backend=false")
					if err != nil {
						log.Panic(err)
					}
				} else {
					log.Info("No source module configured, moving on.")
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}

	if modified := cli.AnyModified(terraformDir, []string{lockFile}); !modified {
		log.Info("terraform lockfile is up to date")
	} else {
		err = cli.Add(terraformDir, lockFile)
		if err != nil {
			log.Panic(err)
		}

		err = cli.Commit(workspace, commitMsg)
		if err != nil {
			log.Fatal(err)
		}

		err = cli.Push(workspace, pullRequestBranch)
		if err != nil {
			log.Fatal(err)
		}

		client := github.GetClient(token)
		err = github.CreatePullRequest(client, ownerName, repoName, "Update Terraform Lockfile", pullRequestBranch, targetBranch)
		if err != nil {
			log.Fatal(err)
		}
	}

}
