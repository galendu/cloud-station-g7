# cloud-station-g7

## Contents  
- Introduction  
- Quick Start  
- 
## Introduction  

cloud-station-g7 is a cloud object storage client transfer station, providing storage systems for Alibaba Cloud obs, Tencent Cloud obs, Huawei Cloud obs and minio

## Quick Start

```text
$ go run main.go  --help
cloud-station-cli 云中转站

Usage:
  cloud-station-cli [flags]
  cloud-station-cli [command]

Examples:
cloud-station-cli cmds

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  upload      upload 文件上传

Flags:
  -h, --help     help for cloud-station-cli
  -v, --versin   cloud station 版本信息

Use "cloud-station-cli [command] --help" for more information about a command.

upload 文件上传

Usage:
  cloud-station-cli upload [flags]

Examples:
upload -f filename

Flags:
  -k, --access_key string      oss storage provier ak
  -s, --access_secret string   oss storage provier sk
  -b, --bucket_name string     oss storage provier bucket name (default "galendu")
  -e, --endpoint string        oss storage provier bucket name (default "oss-cn-beijing.aliyuncs.com")
  -h, --help                   help for upload
  -p, --provider string        oss storage provier [aliyun/tx/aws] (default "aliyun")
  -f, --upload_file string     upload file name

Global Flags:
  -v, --versin   cloud station 版本信息
```