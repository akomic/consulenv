package commands

import (
	"fmt"
	"os"

	"consulenv/consul"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config  string //
	daemon  bool   //
	version bool   //

	// Cmd ...
	Cmd = &cobra.Command{
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

	Cmd.PersistentFlags().StringP("config", "c", "", "config file")

	Cmd.PersistentFlags().StringP("addr", "", "127.0.0.1:8500", "Consul server address")
	Cmd.PersistentFlags().StringP("token", "", "", "Consul token")
	Cmd.PersistentFlags().StringP("auth", "", "", "Consul server API user:pass")
	Cmd.PersistentFlags().StringP("ssl", "", "false", "Consul server HTTPS")

	Cmd.PersistentFlags().MarkHidden("addr")
	Cmd.PersistentFlags().MarkHidden("token")
	Cmd.PersistentFlags().MarkHidden("auth")
	Cmd.PersistentFlags().MarkHidden("ssl")

	Cmd.PersistentFlags().StringSliceP("path", "p", nil, "Path")
	Cmd.PersistentFlags().BoolP("export", "e", false, "Export bash format")
	Cmd.PersistentFlags().BoolP("json", "j", false, "Return in JSON format")
	Cmd.PersistentFlags().BoolP("verbose", "v", false, "Verbosity")
	Cmd.PersistentFlags().BoolP("keys", "k", false, "List keys under prefix")

	viper.BindPFlag("config", Cmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("addr", Cmd.PersistentFlags().Lookup("addr"))
	viper.BindPFlag("token", Cmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("auth", Cmd.PersistentFlags().Lookup("auth"))
	viper.BindPFlag("ssl", Cmd.PersistentFlags().Lookup("ssl"))

	viper.BindPFlag("path", Cmd.PersistentFlags().Lookup("path"))
	viper.BindPFlag("export", Cmd.PersistentFlags().Lookup("export"))
	viper.BindPFlag("json", Cmd.PersistentFlags().Lookup("json"))
	viper.BindPFlag("verbose", Cmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("keys", Cmd.PersistentFlags().Lookup("keys"))

	viper.BindEnv("addr", "CONSUL_HTTP_ADDR")
	viper.BindEnv("token", "CONSUL_HTTP_TOKEN")
	viper.BindEnv("auth", "CONSUL_HTTP_AUTH")
	viper.BindEnv("ssl", "CONSUL_HTTP_SSL")
}

func initConfig() {
	cfgFile := viper.GetString("config")

	if cfgFile == "" {
		abs, err := filepath.Abs(filepath.Join(os.Getenv("HOME"), ".consulenv/config.yml"))
		if err == nil {
			cfgFile = abs
		}
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.ReadInConfig()
	}

	addr := viper.GetString("addr")
	token := viper.GetString("token")

	if addr == "" || token == "" {
		fmt.Println("You need to configure access to Consul server through: config file/env/flags")
		os.Exit(1)
	}
}

func fetch(ccmd *cobra.Command, args []string) {
	paths := viper.GetStringSlice("path")
	keys := viper.GetBool("keys")

	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "At least one -p required.")
		ccmd.HelpFunc()(ccmd, args)
		os.Exit(1)
	}

	if keys {
		consul.Keys()
	} else {
		consul.Get()
	}
}
