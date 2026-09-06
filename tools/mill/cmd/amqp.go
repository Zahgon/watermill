package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
)

var amqpCmd = &cobra.Command{
	Use:   "amqp",
	Short: "Commands for the AMQP Pub/Sub provider",
	Long: `Consume or produce messages from the AMQP Pub/Sub provider.

For the configuration of consuming/producing of the messages, check the help of the relevant command.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := rootCmd.PersistentPreRunE(cmd, args)
		if err != nil {
			return err
		}

		logger.Debug("Using AMQP Pub/Sub", nil)

		if cmd.Use == "consume" {
			consumer, err = amqp.NewSubscriber(amqpConsumerConfig(), logger)
			if err != nil {
				return err
			}
		}

		if cmd.Use == "produce" {
			producer, err = amqp.NewPublisher(amqpProducerConfig(), logger)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func amqpConsumerConfig() amqp.Config { _ = "STUB: not implemented"; return *new(amqp.Config) }

func amqpProducerConfig() amqp.Config { _ = "STUB: not implemented"; return *new(amqp.Config) }

func init() {
	rootCmd.AddCommand(amqpCmd)
	configureAmqpCmd()
	consumeCmd := addConsumeCmd(amqpCmd, "amqp.consume.queue")
	configureConsumeCmd(consumeCmd)
	produceCmd := addProduceCmd(amqpCmd, "amqp.produce.exchange")
	configureProduceCmd(produceCmd)
}

func configureAmqpCmd() { _ = "STUB: not implemented"; return }

func configureConsumeCmd(consumeCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func configureProduceCmd(produceCmd *cobra.Command) { _ = "STUB: not implemented"; return }
