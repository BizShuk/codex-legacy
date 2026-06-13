# VScode React Note

### Debug

1. run React app first by `npm run start`
2. attach debugger on it by chrome debugger on the port

```json
{
    "type": "chrome",
    "request": "launch",
    "name": "Launch Chrome against localhost",
    "url": "http://localhost",
    "webRoot": "${workspaceFolder}",
    "port": 3000,
}
```
