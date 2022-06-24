package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func GetCmd() *cobra.Command {
	cmdReplace := &cobra.Command{
		Use:   "get [path] [jsonfilename]",
		Short: "Get JSON path value",
		Long:  `Get JSON path value using dot notation`,
		Args:  cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			dotPath := args[0]
			filePath := args[1]

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}

			value := gjson.GetBytes(content, dotPath)
			fmt.Printf("%s \n", value.String())
		},
	}
	return cmdReplace
}
