package main

import (
	"runtime"
)

var (
	/*
		These variables store data about the git repo used to generate this build.
		They are filled in on compile time using (https://github.com/JoshuaDoes/govvv).
	*/

	//GitBranch stores the repo branch
	GitBranch string
	//GitCommit stores the repo commit
	GitCommit string
	//GitCommitFull stores the full repo commit
	GitCommitFull string
	//GitCommitMsg stores the changes made for the current repo commit
	GitCommitMsg string
	//GitState stores whether this build is clean from the repo or has unstaged work done
	GitState string
	//BuildDate stores the build date
	BuildDate string

	/*
		These variables store human-readable build identifiers.
	*/

	//BuildID stores a unique build ID inspired by the Android Open Source Project
	BuildID = "bedrocksrv-" + GitState + " " + GitBranch + "-" + GitCommit

	//GitHubCommitURL stores the URL to the current commit
	GitHubCommitURL = "https://github.com/JoshuaDoes/bedrocksrv/commit/" + GitCommitFull

	//GolangVersion stores the version of Go used to build this release
	GolangVersion = runtime.Version()
)
