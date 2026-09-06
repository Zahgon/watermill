package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ThreeDotsLabs/watermill/message"
)

var producer message.Publisher

func addProduceCmd(parent *cobra.Command, topicKey string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
