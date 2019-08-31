# websocket-dynamic-reverse-proxy


## Reverse Proxy Tool with Dynamic Configuration at Runtime


 - [Getting started](#getting-started)
 - [Supported Web Frameworks](#supported-web-frameworks)
 - [How to use it with Gin](#how-to-use-it-with-gin)
 - [Implementation Status](#implementation-status)
 - [Declarative Comments Format](#declarative-comments-format)
	- [General API Info](#general-api-info)
	- [API Operation](#api-operation)
	- [Security](#security)
 - [Examples](#examples)
	- [Descriptions over multiple lines](#descriptions-over-multiple-lines)
	- [User defined structure with an array type](#user-defined-structure-with-an-array-type)
	- [Add a headers in response](#add-a-headers-in-response) 
	- [Use multiple path params](#use-multiple-path-params)
	- [Example value of struct](#example-value-of-struct)
	- [Description of struct](#description-of-struct)
	- [Use swaggertype tag to supported custom type](#use-swaggertype-tag-to-supported-custom-type)
	- [Add extension info to struct field](#add-extension-info-to-struct-field)
	- [How to using security annotations](#how-to-using-security-annotations)
- [About the Project](#about-the-project)

## Getting started



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
