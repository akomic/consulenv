package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"consulenv/consul"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config  string //
	daemon  bool   //
	version bool   //

	// MyCmd ...
	MyCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,

		PersistentPreRun: func(ccmd *cobra.Command, args []string) {
			// get the filepath
			abs, err := filepath.Abs(filepath.Join(os.Getenv("HOME"), ".consulenv/config.yml"))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading filepath: ", err.Error())
			}

			// get the config name
			base := filepath.Base(abs)

			// get the path
			path := filepath.Dir(abs)

			//
			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			// Find and read the config file; Handle errors reading the config file
			if err := viper.ReadInConfig(); err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read config file: ", err.Error())
				os.Exit(1)
			}
		},

		Run: fetch,
	}
)

var (
	consulPrefix string
)

func init() {
	MyCmd.PersistentFlags().StringP("server", "", "", "Consul server address")
	MyCmd.PersistentFlags().StringP("scheme", "", "http", "Consul server scheme")
	MyCmd.PersistentFlags().StringP("user", "", "", "Consul server API user")
	MyCmd.PersistentFlags().StringP("pass", "", "", "Consul server API pass")
	MyCmd.PersistentFlags().StringP("token", "", "", "Consul token")

	MyCmd.PersistentFlags().MarkHidden("server")
	MyCmd.PersistentFlags().MarkHidden("scheme")
	MyCmd.PersistentFlags().MarkHidden("user")
	MyCmd.PersistentFlags().MarkHidden("pass")
	MyCmd.PersistentFlags().MarkHidden("token")

	MyCmd.Flags().StringSliceP("prefix", "p", nil, "Prefix")

	viper.BindPFlag("server", MyCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("user", MyCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pass", MyCmd.PersistentFlags().Lookup("pass"))
	viper.BindPFlag("scheme", MyCmd.PersistentFlags().Lookup("scheme"))
	viper.BindPFlag("token", MyCmd.PersistentFlags().Lookup("token"))

	viper.BindPFlag("prefix", MyCmd.Flags().Lookup("prefix"))
}

func fetch(ccmd *cobra.Command, args []string) {
	prefixes := viper.GetStringSlice("prefix")

	if len(prefixes) == 0 {
		fmt.Println("At least one -p required.")
		ccmd.HelpFunc()(ccmd, args)
	}

	consul.Get()
}
