package main

import (
	"code.google.com/p/go.net/context"
	"github.com/m0sth8/vane/pkg/target"
	"github.com/spf13/cobra"
)

var (
	url *string
)

var VaneCmd = &cobra.Command{
	Use:   "vane",
	Short: "Vane is a WordPress vulnerability scanner",
	Long:  `Vane is a GPL fork of the now non-free popular WordPress vulnerability scanner WPScan.`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	if *url == "" {
		cmd.Println("The URL is mandatory, please supply it with --url= or -u")
		cmd.Usage()
		return
	}
	wpTarget := target.NewTarget(*url)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if !wpTarget.WPSite.IsOnline(ctx) {
		cmd.Printf("The WordPress URL supplied '%s' seems to be down.\n", wpTarget.Url)
		return
	}

}

func main() {
	url = VaneCmd.Flags().StringP("url", "u", "", "The WordPress URL/domain to scan")
	//	VaneCmd.AddCommand()
	//	VaneCmd.DebugFlags()
	VaneCmd.Execute()
}
