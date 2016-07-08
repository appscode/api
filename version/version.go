package version

import (
	"bytes"
	"fmt"
)

var (
	Name            string
	Version         string
	Os              string
	Arch            string
	CommitHash      string
	GitBranch       string
	GitTag          string
	CommitTimestamp string
	BuildTimestamp  string
	BuildHost       string
	BuildHostOs     string
	BuildHostArch   string
)

func Print() {
	fmt.Printf("%v\n", Name)
	fmt.Printf("Version = %v\n", Version)
	fmt.Printf("Os = %v\n", Os)
	fmt.Printf("Arch = %v\n", Arch)

	fmt.Printf("CommitHash = %v\n", CommitHash)
	fmt.Printf("GitBranch = %v\n", GitBranch)
	fmt.Printf("GitTag = %v\n", GitTag)
	fmt.Printf("CommitTimestamp = %v\n", CommitTimestamp)

	fmt.Printf("BuildTimestamp = %v\n", BuildTimestamp)
	fmt.Printf("BuildHost = %v\n", BuildHost)
	fmt.Printf("BuildHostOs = %v\n", BuildHostOs)
	fmt.Printf("BuildHostArch = %v\n", BuildHostArch)
}

func String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Version = %v\n", Version))
	buf.WriteString(fmt.Sprintf("Os = %v\n", Os))
	buf.WriteString(fmt.Sprintf("Arch = %v\n", Arch))

	buf.WriteString(fmt.Sprintf("CommitHash = %v\n", CommitHash))
	buf.WriteString(fmt.Sprintf("GitBranch = %v\n", GitBranch))
	buf.WriteString(fmt.Sprintf("GitTag = %v\n", GitTag))
	buf.WriteString(fmt.Sprintf("CommitTimestamp = %v\n", CommitTimestamp))

	buf.WriteString(fmt.Sprintf("BuildTimestamp = %v\n", BuildTimestamp))
	buf.WriteString(fmt.Sprintf("BuildHost = %v\n", BuildHost))
	buf.WriteString(fmt.Sprintf("BuildHostOs = %v\n", BuildHostOs))
	buf.WriteString(fmt.Sprintf("BuildHostArch = %v\n", BuildHostArch))
	return buf.String()
}
