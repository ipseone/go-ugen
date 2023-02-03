package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.ipse.one/ugen"
)

var mainCmd = &cobra.Command{
	Use:          "ugen",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		stdoutStat, err := os.Stdout.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat stdout: %w", err)
		}

		g := &ugen.Generator{
			AnyCharacter: viper.GetBool("any"),
			Base64:       viper.GetBool("base64"),
			Hex:          viper.GetBool("hex"),
			ULID:         viper.GetBool("ulid"),
			UUID:         viper.GetBool("uuid"),
			Lower:        viper.GetBool("lower"),
			Upper:        viper.GetBool("upper"),
			Prefix:       viper.GetString("prefix"),
			Suffix:       viper.GetString("suffix"),
			Separator:    viper.GetString("separator"),
			WithLineFeed: (stdoutStat.Mode() & os.ModeCharDevice) == os.ModeCharDevice,
		}

		if err := g.Gen(os.Stdout, viper.GetInt("count"), viper.GetInt("length")); err != nil {
			return fmt.Errorf("failed to generate: %w", err)
		}

		return nil
	},
}

func init() {
	mainCmd.PersistentFlags().BoolP("any", "a", false, "Use any character from the random generator.")
	if err := viper.BindPFlag("any", mainCmd.PersistentFlags().Lookup("any")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("base64", false, "Output base64 encoded strings. Sets --any.")
	if err := viper.BindPFlag("base64", mainCmd.PersistentFlags().Lookup("base64")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().IntP("count", "c", 1, "The number of unique strings to generate.")
	if err := viper.BindPFlag("count", mainCmd.PersistentFlags().Lookup("count")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("hex", false, "Output hex encoded strings. Sets --any.")
	if err := viper.BindPFlag("hex", mainCmd.PersistentFlags().Lookup("hex")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().IntP("length", "l", 12, "Length of the generated string.")
	if err := viper.BindPFlag("length", mainCmd.PersistentFlags().Lookup("length")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("lower", false, "Output characters in lower case.")
	if err := viper.BindPFlag("lower", mainCmd.PersistentFlags().Lookup("lower")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().String("prefix", "", "Write prefix before each generated string.")
	if err := viper.BindPFlag("prefix", mainCmd.PersistentFlags().Lookup("prefix")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().StringP("separator", "s", "\n", "Separator for generated strings.")
	if err := viper.BindPFlag("separator", mainCmd.PersistentFlags().Lookup("separator")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().String("suffix", "", "Write suffix after each generated string.")
	if err := viper.BindPFlag("suffix", mainCmd.PersistentFlags().Lookup("suffix")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("ulid", false, "Generate a ULID.")
	if err := viper.BindPFlag("ulid", mainCmd.PersistentFlags().Lookup("ulid")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("upper", false, "Output characters in upper case.")
	if err := viper.BindPFlag("upper", mainCmd.PersistentFlags().Lookup("upper")); err != nil {
		panic(err)
	}

	mainCmd.PersistentFlags().Bool("uuid", false, "Generate a random (v4) UUID.")
	if err := viper.BindPFlag("uuid", mainCmd.PersistentFlags().Lookup("uuid")); err != nil {
		panic(err)
	}
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
