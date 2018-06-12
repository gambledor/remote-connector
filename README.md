# Remote Connector
Remote connector opens an ssh interactive session to chosen remote computer.
It looks for a hidden file named *.remote_connections* into logged user home directory.
The file content is an array of json object which syntax is, here below:
```json
[
    {
        "user": "<ssh-username>",
        "name": "<my remote hostname alias>",
        "host": "<remote hostname/ip>"
        "protocol": "<ssh>"
    }
]
```
