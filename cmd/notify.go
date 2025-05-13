package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type NotifyOptions struct {
	ChatName string
	Message  string
}

var notifyOptions NotifyOptions

var notifyCmd = &cobra.Command{
	Use:   "notify [communication application] [flags]",
	Short: "Notify the user about the events",
	Long:  "Notify the user about the events. The communication application can be any communication application that is supported by this application.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		communicationApp := args[0]
		fmt.Println("Notifying user using communication application:", communicationApp)
		fmt.Println("Chat Name:", notifyOptions.ChatName)
		fmt.Println("Message:", notifyOptions.Message)
	},
}

func init() {
	// add the notify command to the root command
	rootCmd.AddCommand(notifyCmd)

	// add any flags required for the notify command here
	notifyCmd.Flags().StringVarP(&notifyOptions.ChatName, "chat-name", "c", "", "Name of the chat to notify")
	notifyCmd.MarkFlagRequired("chat-name")
	notifyCmd.Flags().StringVarP(&notifyOptions.Message, "message", "m", "", "Message to send in the notification")
	notifyCmd.MarkFlagRequired("message")
}
