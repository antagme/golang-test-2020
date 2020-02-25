attaching outputs.


```

curl -i -X POST -F file=@71UxUnp-LCL._AC_SL1500_.jpg   localhost:8080/upload
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Date: Tue, 25 Feb 2020 13:31:23 GMT
Content-Length: 90
Content-Type: text/plain; charset=utf-8

File Uploaded localhost:8080/5577006791947779410_modified.jpg Stored in Redis with key: 1
```
