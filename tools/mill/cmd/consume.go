package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ThreeDotsLabs/watermill/message"
)

var consumer message.Subscriber

func addConsumeCmd(parent *cobra.Command, topicKey string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
