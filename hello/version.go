package main

import "fmt"

var gitCommit = "N/A"
var buildDate = "N/A"

type versionInfo struct {
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
}

var ver = versionInfo{
	GitCommit: gitCommit,
	BuildDate: buildDate,
}

func (v versionInfo) String() string {
	return fmt.Sprintf("Commit: %s Built at: %s", v.GitCommit, v.BuildDate)
}
