package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/monirz/postgo/pkg/fetcher"
	"github.com/spf13/cobra"
)

var (
	url    string
	method string
	header []string

	rootCmd = &cobra.Command{
		Use:   "postgo",
		Short: "postgo is a REST API testing tool.",
		Run:   fetch,
	}
)

func init() {

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the source file, ex: http://example.com/file.mp4")
	rootCmd.MarkFlagRequired("url")

	rootCmd.Flags().StringVarP(&method, "method", "X", "", "REST API method, ex: GET")
	// rootCmd.MarkFlagRequired("method")

	// rootC

	rootCmd.Flags().StringArrayVarP(&header, "header", "H", []string{}, "HTTP Header, ex: -H 'Content-Type: application/json' ")
	rootCmd.MarkFlagRequired("header")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func fetch(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		fmt.Println("Invalid arguments")
		os.Exit(-1)
	}

	if url == "" {
		fmt.Println("URL is required")
		os.Exit(-1)
	}

	if method == "" {
		method = "GET"
	}

	headers := make(map[string]string)

	if len(header) > 0 {
		for _, v := range header {
			splitted := strings.Split(v, ":")

			headers[strings.TrimSpace(splitted[0])] = strings.TrimSpace(splitted[1])

		}
	}

	//get it from arguments
	// contentType := make(map[string]string)

	resp, err := fetcher.Fetch(url, method, headers)

	if err != nil {
		log.Fatal(err)
	}

	if resp != nil {
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		}
	}

	resp.Body.Close()
}
