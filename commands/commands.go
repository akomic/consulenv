package commands

import (
	"fmt"
	"os"
	// "path/filepath"
	// "strings"

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
		Run:   fetch,
	}
)

var (
	consulpath string
)

func init() {
	cobra.OnInitialize(initConfig)

	MyCmd.PersistentFlags().StringP("config", "c", "", "config file")

	MyCmd.PersistentFlags().StringP("addr", "", "127.0.0.1:8500", "Consul server address")
	MyCmd.PersistentFlags().StringP("token", "", "", "Consul token")
	MyCmd.PersistentFlags().StringP("auth", "", "", "Consul server API user:pass")
	MyCmd.PersistentFlags().StringP("ssl", "", "false", "Consul server HTTPS")

	MyCmd.PersistentFlags().MarkHidden("addr")
	MyCmd.PersistentFlags().MarkHidden("token")
	MyCmd.PersistentFlags().MarkHidden("auth")
	MyCmd.PersistentFlags().MarkHidden("ssl")

	MyCmd.PersistentFlags().StringSliceP("path", "p", nil, "Path")
	MyCmd.PersistentFlags().BoolP("export", "e", false, "Export bash format")
	MyCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbosity")

	viper.BindPFlag("config", MyCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("addr", MyCmd.PersistentFlags().Lookup("addr"))
	viper.BindPFlag("token", MyCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("auth", MyCmd.PersistentFlags().Lookup("auth"))
	viper.BindPFlag("ssl", MyCmd.PersistentFlags().Lookup("ssl"))

	viper.BindPFlag("path", MyCmd.PersistentFlags().Lookup("path"))
	viper.BindPFlag("export", MyCmd.PersistentFlags().Lookup("export"))
	viper.BindPFlag("verbose", MyCmd.PersistentFlags().Lookup("verbose"))

	viper.BindEnv("addr", "CONSUL_HTTP_ADDR")
	viper.BindEnv("token", "CONSUL_HTTP_TOKEN")
	viper.BindEnv("auth", "CONSUL_HTTP_AUTH")
	viper.BindEnv("ssl", "CONSUL_HTTP_SSL")
}

func initConfig() {
	cfgFile := viper.GetString("config")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read config file: ", err.Error())
			os.Exit(1)
		}
	}

	addr := viper.GetString("addr")

	if addr == "" {
		fmt.Println("You need to configure access to Consul server through: config file/env/flags")
		os.Exit(1)
	}
}

func fetch(ccmd *cobra.Command, args []string) {
	paths := viper.GetStringSlice("path")

	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "At least one -p required.")
		ccmd.HelpFunc()(ccmd, args)
		os.Exit(1)
	}

	consul.Get()
}
