# websocket-dynamic-reverse-proxy


## Reverse Proxy Tool with Dynamic Configuration at Runtime



- Listener Websocket Addr : 0.0.0.0:8181

- Upstream Websocket Addr : [Custom]

- Custom SettingUp Proxy :  0.0.0.0:6661

AddProxy
```javascript
request({
    uri: "http://127.0.0.1:6161/add",
    method: "POST",
    body: {
        listenerPath:"/some-listener-addr",
        upstream:"ws://127.0.0.1:3333/"
    }
})
```


HealthCheck
```javascript

request({
    uri: "http://127.0.0.1:6161/",
    method: "GET",
})
```