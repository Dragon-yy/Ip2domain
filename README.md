# ip2domain

ip2domain is a CLI tool to lookup domains associated with an IP address using the FOFA API and save the results to a CSV file.

## Usage

```sh
ip2domain --apikey <FOFA_API_KEY> --ip <IP_ADDRESS> --output <OUTPUT_CSV_FILE>
ip2domain --apikey <FOFA_API_KEY> --batch <BATCH_FILE> --output <OUTPUT_CSV_FILE>
--apikey or -k: FOFA API key (required)
--ip or -i: IP address to lookup
--batch or -b: File containing multiple IP addresses to lookup
--output or -o: Output CSV file (default: results.csv)
````

## Example

```sh
ip2domain --apikey 5**************2 --ip 154.210.40.225 --output results.csv

ip2domain --apikey 5**************2 --batch ips.txt --output results.csv

```
