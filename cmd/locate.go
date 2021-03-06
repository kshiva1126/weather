// Copyright © 2019 kshiva1126 <kshiva1126@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/kshiva1126/weather/common"

	"github.com/spf13/cobra"
)

const (
	ENDPOINT = "http://api.openweathermap.org/data/2.5/forecast"
)

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "notice the weather of locate",
	Long:  `notice the weather of the arg locate(For example Tokyo, Osaka, Fukuoka, and so on)`,
	Run: func(cmd *cobra.Command, args []string) {
		common.EnvLoad()
		val := url.Values{}

		place := args[0]
		val.Add("city", place)
		key := os.Getenv("API_KEY")

		resp, err := http.Get(ENDPOINT + "?q=" + val.Get("city") + "&units=metric&APPID=" + key)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		common.ParseJsonReceivedAndExecute(resp)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(locateCmd)
}
