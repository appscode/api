package version

import "fmt"

func (m *Version) Print() {
	if m.Name != "" {
		fmt.Printf("Name = %v\n", m.Name)
	}
	if m.Version != "" {
		fmt.Printf("Version = %v\n", m.Version)
	}
	if m.VersionStrategy != "" {
		fmt.Printf("VersionStrategy = %v\n", m.VersionStrategy)
	}
	if m.Os != "" {
		fmt.Printf("Os = %v\n", m.Os)
	}
	if m.Arch != "" {
		fmt.Printf("Arch = %v\n", m.Arch)
	}

	if m.CommitHash != "" {
		fmt.Printf("CommitHash = %v\n", m.CommitHash)
	}
	if m.GitBranch != "" {
		fmt.Printf("GitBranch = %v\n", m.GitBranch)
	}
	if m.GitTag != "" {
		fmt.Printf("GitTag = %v\n", m.GitTag)
	}
	if m.CommitTimestamp != "" {
		fmt.Printf("CommitTimestamp = %v\n", m.CommitTimestamp)
	}

	if m.BuildTimestamp != "" {
		fmt.Printf("BuildTimestamp = %v\n", m.BuildTimestamp)
	}
	if m.BuildHost != "" {
		fmt.Printf("BuildHost = %v\n", m.BuildHost)
	}
	if m.BuildHostOs != "" {
		fmt.Printf("BuildHostOs = %v\n", m.BuildHostOs)
	}
	if m.BuildHostArch != "" {
		fmt.Printf("BuildHostArch = %v\n", m.BuildHostArch)
	}
}
