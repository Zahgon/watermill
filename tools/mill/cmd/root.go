package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"

	"github.com/ThreeDotsLabs/watermill"
)

var cfgFile string
var logger watermill.LoggerAdapter

var rootCmd = &cobra.Command{
	Use:   "mill",
	Short: "A CLI for Watermill.",
	Long: `A CLI for Watermill.

Use console-based producer or consumer for various pub/sub providers.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log := viper.GetBool("log")
		debug := viper.GetBool("debug")
		trace := viper.GetBool("trace")
		if log || debug || trace {
			logger = watermill.NewStdLogger(debug, trace)
		} else {
			logger = watermill.NopLogger{}
		}

		if err := checkRequiredFlags(cmd.Flags()); err != nil {
			return err
		}

		writeConfig := viper.GetString("writeConfig")
		if writeConfig != "" {
			settings := viper.AllSettings()
			delete(settings, "writeconfig")
			b, err := yaml.Marshal(settings)
			if err != nil {
				return errors.Wrap(err, "could not marshal config to yaml")
			}

			f, err := os.Create(writeConfig)
			if err != nil {
				return errors.Wrap(err, "could not create file for write")
			}
			_, err = fmt.Fprintf(f, "%s", b)
			if err != nil {
				return errors.Wrap(err, "could not write to file")
			}
		}

		return nil
	},
}

func Execute() { _ = "STUB: not implemented"; return }

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().SortFlags = false

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mill.yaml)")

	outputFlags := pflag.NewFlagSet("output", pflag.ExitOnError)
	outputFlags.BoolP("log", "l", false, "If true, the logger output is sent to stderr. No logger output otherwise.")
	ensure(viper.BindPFlag("log", outputFlags.Lookup("log")))

	outputFlags.BoolP("debug", "d", false, "If true, debug output is enabled from the logger")
	ensure(viper.BindPFlag("debug", outputFlags.Lookup("debug")))

	outputFlags.Bool("trace", false, "If true, trace output is enabled from the logger")
	ensure(viper.BindPFlag("trace", outputFlags.Lookup("trace")))

	outputFlags.String("write-config", "", "Write the config of the current command as yaml to the specified path")
	ensure(viper.BindPFlag("writeConfig", outputFlags.Lookup("write-config")))

	rootCmd.PersistentFlags().AddFlagSet(outputFlags)
}

func initConfig() { _ = "STUB: not implemented"; return }

func ensure(err error) { _ = "STUB: not implemented"; return }

func checkRequiredFlags(flags *pflag.FlagSet) error { _ = "STUB: not implemented"; return nil }
