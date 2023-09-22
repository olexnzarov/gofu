package update

import (
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

type updateBuilder struct {
	cmd    *cobra.Command
	config *pb.ProcessConfiguration
	mask   []string
}

func newUpdateBuilder(cmd *cobra.Command) *updateBuilder {
	return &updateBuilder{
		cmd:  cmd,
		mask: []string{},
		config: &pb.ProcessConfiguration{
			RestartPolicy: &pb.ProcessConfiguration_RestartPolicy{},
		},
	}
}

func (u *updateBuilder) includeFlag(flagName string, updateFunc func(config *pb.ProcessConfiguration) (field string)) {
	flag := u.cmd.Flags().Lookup(flagName)
	if flag == nil || !flag.Changed {
		return
	}
	field := updateFunc(u.config)
	u.mask = append(u.mask, field)
}
