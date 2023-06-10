package main

var gitCommit = "N/A"
var buildDate = "N/A"

var versionInfo = struct {
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
}{
	GitCommit: gitCommit,
	BuildDate: buildDate,
}
