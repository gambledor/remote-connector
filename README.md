# Remote Connector
Remote connector opens an ssh interactive session chosen by a CLI menu'.
The menu' is generated looking for a hidden file named *.remote_connections* into logged user home directory.
The file content is an array of json object which syntax is, here below:
```json
{
  "machines": [
    {
        "user": "<ssh-username>",
        "name": "<my remote hostname alias>",
        "host": "<remote hostname/ip>",
        "protocol": "<ssh>"
    }
  ]
}
```

Each json object stays for a remote connection.

## Command line params

-c <id>, To connect to the id menù number machine
-X, The connection is X mode enabled (-X)
