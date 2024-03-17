# GO-CLEANER

## To create executable file

``` go
env GOOS=windows GOARCH=amd64 go build -o cleaner_ver.exe
```

## How to use

On first run it creates `cleaner_config.yml` file with default parameters.
You can change any, but service wont go futher without this config.

### Config attributes

|Value|Type|Description|
|-|-|-|
|`ready`|boolean|if `false` it doesn't start to avoid unconfigurable run|
|`real`|boolean|on `false` it just analizes and builds `dump_file.txt` with results, if `true` it removes files|
|`path`|string|path to folder for checking|
