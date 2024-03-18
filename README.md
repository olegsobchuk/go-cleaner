# GO-CLEANER

## To create executable file

### Windows

``` go
env GOOS=windows GOARCH=amd64 go build -o cleaner_ver.exe
```

### Linux

``` go
env GOOS=linux GOARCH=amd64 go build -o cleaner_ver
```

## How to use

On first run it creates `cleaner_config.yml` file with default parameters.
You can change any, but service wont go futher without this config.

### Config attributes

|Value|Type|Description|
|-|-|-|
|`path`|string|path to folder for checking|
|`real`|boolean|on `false` it just analizes and builds `dump_file.txt` with results, if `true` it removes files|
|`ready`|boolean|if `false` it doesn't start to avoid unconfigurable run|
|`size`|int64|file size threshold, when it should be ignorred. Files with size more than this value will be skipped for futher checking|
|`extensions`|object|check dedicated description|
|`files`|object|check dedicated description|
|`content`|[]string|the list of suspicious words in files|

`extensions` description

|Value|Type|Description|
|-|-|-|
|`whitelist`|[]string|the list of extensions. Files with these extensions will be ignorred for futher checking|
|`blacklist`|[]string|the list of extensions. Files with these extensions will be marked as suspicious without futher checking|

`files` description

|Value|Type|Description|
|-|-|-|
|`whitelist`|[]string|the list of filenames (with extension). Files with these names will be ignorred for futher checking|
|`blacklist`|[]string|the list of filenames (with extension). Files with these names will be marked as suspicious without futher checking|
