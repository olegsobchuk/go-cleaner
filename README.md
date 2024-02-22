# GO-CLEANER

## To create executable file

``` go
env GOOS=windows GOARCH=amd64 go build -o cleaner.exe
```

## How to use

First run it creates `cleaner_config.yml` file with default parameters.
You can change any, but service wont go futher without this config.

### Config attrebutes

|Value|Type|Description|
|-|-|-|
|`path`|string|path to folder for checking|
|`real`|boolean|if `true` it removes files, if `false` - just analize and builds `dump_file.txt` with results|
|`ready`|boolean|if `false` it doesn't start, if `true` goes futher|
