package cmd

import (
	"bufio"
	"fmt"
	"ip2domain/fofa"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	apiKey    string
	ip        string
	output    string
	batchFile string
)

var rootCmd = &cobra.Command{
	Use:   "ip2domain",
	Short: "IP to Domain lookup tool using FOFA API",
	Run: func(cmd *cobra.Command, args []string) {
		if apiKey == "" {
			fmt.Println("API key is required")
			return
		}

		if batchFile != "" {
			file, err := os.Open(batchFile)
			if err != nil {
				fmt.Printf("Error opening batch file: %v\n", err)
				return
			}
			defer file.Close()

			var allResults [][]string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				ip = strings.TrimSpace(scanner.Text())
				if ip == "" {
					continue
				}
				results, err := fofa.FetchDomains(apiKey, ip)
				if err != nil {
					fmt.Printf("Error fetching domains for IP %s: %v\n", ip, err)
					continue
				}
				allResults = append(allResults, results...)
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading batch file: %v\n", err)
			}

			err = fofa.SaveToCSV(allResults, output)
			if err != nil {
				fmt.Printf("Error saving to CSV: %v\n", err)
			}
		} else if ip != "" {
			results, err := fofa.FetchDomains(apiKey, ip)
			if err != nil {
				fmt.Printf("Error fetching domains: %v\n", err)
				return
			}
			err = fofa.SaveToCSV(results, output)
			if err != nil {
				fmt.Printf("Error saving to CSV: %v\n", err)
			}
		} else {
			fmt.Println("IP address or batch file is required")
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "FOFA API key")
	rootCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP address to lookup")
	rootCmd.Flags().StringVarP(&output, "output", "o", "results.csv", "Output CSV file")
	rootCmd.Flags().StringVarP(&batchFile, "batch", "b", "", "File containing multiple IP addresses to lookup")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
