/*
Copyright © 2020 Hotaka Saito <teengenerate@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hbagdi/go-kong/kong"
)

// consumersListCmd represents the consumersList command
var consumersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List All Consumers",
	Long:  `https://docs.konghq.com/1.4.x/admin-api/#list-all-consumers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("admin.url")
		if &url == nil {
			return errors.New("required admin url")
		}

		client, err := kong.NewClient(&url, nil)
		if err != nil {
			return err
		}
		consumers, err := client.Consumers.ListAll(context.Background())
		for _, consumer := range consumers {
			fmt.Printf("Username: %v\n", *consumer.Username)
		}

		return nil
	},
}

func init() {
	consumersCmd.AddCommand(consumersListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumersListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumersListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
