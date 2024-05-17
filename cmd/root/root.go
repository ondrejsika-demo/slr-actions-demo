package root

import (
	"github.com/ondrejsika-demo/slr-actions-demo/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "slr-actions-demo",
	Short: "slr-actions-demo, " + version.Version,
}
