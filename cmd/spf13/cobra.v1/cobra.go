package cobra

import (
	"fmt"
	"net"
	"time"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/spf13/cobra"
)

// NewCommand returns Command.
func NewCommand(cmd *cobra.Command, cmds ...*cobra.Command) *cobra.Command {
	if len(cmds) > 0 {
		cmd.AddCommand(cmds...)
	}
	return cmd
}

// Run uses the args and run through the command tree finding
// appropriate matches for commands and then corresponding flags.
func Run(rootCmd *cobra.Command, cmds ...*cobra.Command) error {

	rootCmd.AddCommand(cmds...)

	rootCmd.DisableFlagParsing = true

	for _, entry := range config.Entries() {
		parseFlag(rootCmd, entry)
	}

	rootCmd.PersistentFlags().StringSlice(config.ConfArgument, nil, "path to one or more config files")

	return rootCmd.Execute()
}

func parseFlag(cmd *cobra.Command, c config.Config) { // nolint

	switch t := c.Value.(type) {

	case string:
		cmd.PersistentFlags().String(c.Key, t, c.Description)
	case []string:
		cmd.PersistentFlags().StringSlice(c.Key, t, c.Description)
	case map[string]string:
		var s string
		for key, val := range t {
			s = s + fmt.Sprintf("%s=\"%s\" ", key, val)
		}
		cmd.PersistentFlags().String(c.Key, s, c.Description)
	case bool:
		cmd.PersistentFlags().Bool(c.Key, t, c.Description)
	case []bool:
		cmd.PersistentFlags().BoolSlice(c.Key, t, c.Description)
	case int:
		cmd.PersistentFlags().Int(c.Key, t, c.Description)
	case []int:
		cmd.PersistentFlags().IntSlice(c.Key, t, c.Description)
	case int8:
		cmd.PersistentFlags().Int8(c.Key, t, c.Description)
	case int16:
		cmd.PersistentFlags().Int16(c.Key, t, c.Description)
	case int32:
		cmd.PersistentFlags().Int32(c.Key, t, c.Description)
	case int64:
		cmd.PersistentFlags().Int64(c.Key, t, c.Description)
	case uint:
		cmd.PersistentFlags().Uint(c.Key, t, c.Description)
	case uint64:
		cmd.PersistentFlags().Uint64(c.Key, t, c.Description)
	case time.Duration:
		cmd.PersistentFlags().Duration(c.Key, t, c.Description)
	case []byte:
		cmd.PersistentFlags().BytesBase64(c.Key, t, c.Description)
	case float64:
		cmd.PersistentFlags().Float64(c.Key, t, c.Description)
	case net.IPNet:
		cmd.PersistentFlags().IPNet(c.Key, t, c.Description)
	case net.IP:
		cmd.PersistentFlags().IP(c.Key, t, c.Description)
	case net.IPMask:
		cmd.PersistentFlags().IPMask(c.Key, t, c.Description)
	default:
		log.Warnf("key %s with unknown type %s", c.Key, t)
	}

}
