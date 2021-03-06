package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	goversion "go.hein.dev/go-version"
)

var (
	buildVersion = "dev"
	buildCommit  = "none"
	buildDate    = "unknown"
	versionCmd   = &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			versionOutput := goversion.New(buildVersion, buildCommit, buildDate)

			fmt.Printf("%+v", versionOutput.ToJSON())
		},
	}
)
