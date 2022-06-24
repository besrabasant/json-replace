package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func ReplaceCmd() *cobra.Command {
	cmdReplace := &cobra.Command{
		Use:   "replace [path] [value] [jsonfilename]",
		Short: "Replace JSON path value",
		Long:  `Replace JSON path value using got notation`,
		Args:  cobra.ExactValidArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			dotPath := args[0]
			replaceValue := args[1]
			filePath := args[2]

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}

			var objmap map[string]interface{}

			if err := json.Unmarshal(content, &objmap); err != nil {
				log.Fatal("Error parsing json data: ", err)
			}

			dotPathArray := strings.Split(dotPath, ".")

			newMap := make(map[string]interface{}, 1)

			for i := len(dotPathArray); i > 0; i-- {

				tempJson := make(map[string]interface{}, 1)
				part := dotPathArray[i-1]

				if i == len(dotPathArray) {
					tempJson[part] = replaceValue
					newMap = tempJson
				} else {
					tempJson[part] = newMap 
				}
				
				newMap = tempJson
			}

			newJson := mergeMaps(objmap, newMap)

			mJson, err := json.MarshalIndent(newJson, "", "    ")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ioutil.WriteFile(filePath, mJson, 0777)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmdReplace
}


func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
    out := make(map[string]interface{}, len(a))
    for k, v := range a {
        out[k] = v
    }
    for k, v := range b {
        // If you use map[string]interface{}, ok is always false here.
        // Because yaml.Unmarshal will give you map[string]interface{}.
        if v, ok := v.(map[string]interface{}); ok {
            if bv, ok := out[k]; ok {
                if bv, ok := bv.(map[string]interface{}); ok {
                    out[k] = mergeMaps(bv, v)
                    continue
                }
            }
        }
        out[k] = v
    }
    return out
}