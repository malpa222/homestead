package cmd

import (
	"homestead/lib/blogfsys"
	"homestead/lib/server"

	"github.com/spf13/cobra"
)

type serveFlags struct {
	rootName string
	rootVal  string

	portName string
	portVal  string

	httpsName string
	httpsVal  bool
}

var sFlags = serveFlags{
	rootName: "root",
	rootVal:  ".",

	portName: "port",
	portVal:  ":8080",

	httpsName: "https",
	httpsVal:  false,
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		root, _ := cmd.Flags().GetString(sFlags.rootName)
		https, _ := cmd.Flags().GetBool(sFlags.httpsName)
		port, _ := cmd.Flags().GetString(sFlags.portName)

		cfg := server.ServerCFG{
			Port:  port,
			HTTPS: https,
		}

		fsys := blogfsys.NewBlogFsys(root)
		server.Serve(fsys, cfg)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(
		&sFlags.rootVal,
		sFlags.rootName,
		"r",
		sFlags.rootVal,
		"Sets the port of the server")
	serveCmd.MarkFlagDirname(gFlags.rootName)
	serveCmd.MarkFlagRequired(gFlags.rootName)

	serveCmd.Flags().StringVarP(
		&sFlags.portVal,
		sFlags.portName,
		"p",
		sFlags.portVal,
		"The path to the website's root")
	serveCmd.MarkFlagRequired(sFlags.portName)

	serveCmd.Flags().BoolVar(
		&sFlags.httpsVal,
		sFlags.httpsName,
		sFlags.httpsVal,
		"Enables HTTPS. Disabled by default")
}
