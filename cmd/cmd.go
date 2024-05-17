package cmd

import (
	_ "github.com/ondrejsika-demo/slr-actions-demo/cmd/redis_set_large_data"
	"github.com/ondrejsika-demo/slr-actions-demo/cmd/root"
	_ "github.com/ondrejsika-demo/slr-actions-demo/cmd/test_clisso_cli"
	_ "github.com/ondrejsika-demo/slr-actions-demo/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
