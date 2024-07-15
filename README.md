![image](https://github.com/user-attachments/assets/21b99935-f24e-4b01-8f9e-c4d19d1d3f0b)# ip2domain

ip2domain is a CLI tool to lookup domains associated with an IP address using the FOFA API and save the results to a CSV file.
ip2domain 是一个命令行工具，用于使用 FOFA API 查询与 IP 地址相关的域名，并将结果保存到 CSV 文件中。

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
ip2domain --apikey 5**************2 --ip 1.1.1.1 --output results.csv

ip2domain --apikey 5**************2 --batch ips.txt --output results.csv

```

## Result

形成字段为域名、ip、端口的表格
![image](https://github.com/user-attachments/assets/de9b2026-6719-4a17-869b-904704fc170f)
