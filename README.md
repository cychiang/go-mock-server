# Go Mock Server (Working in Progress)
Mock a server with a simple yaml configuration. 

## How to use it?
Start the server as following

```shell
go run cmd/main -config ${CONFIG_FILE_WITH_PATH}
```

## Configuration format
The configuration file is in `yaml` format as following:

```yaml
name: string       # Name of the server
routers:
  - path: string   # Path of the request 
    method: string # Method of the request
    body: string   # Designed response of the request
```
