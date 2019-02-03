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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const (
	ENDPOINT = "http://api.openweathermap.org/data/2.5/forecast"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "notice the weather of locate",
	Long:  `notice the weather of the arg locate(For example Tokyo, Osaka, Fukuoka, and so on)`,
	Run: func(cmd *cobra.Command, args []string) {
		EnvLoad()
		val := url.Values{}
		val.Add("city", "Tokyo")
		key := os.Getenv("API_KEY")

		resp, err := http.Get(ENDPOINT + "?q=" + val.Get("city") + ",jp&APPID=" + key)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		execute(resp)
	},
}

func init() {
	rootCmd.AddCommand(locateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// locateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// locateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
