# OQ

目前功能
1. 極簡的 http server ( 實作 ServeHTTP Interface, 印出 request 內容 ) 
2. 極簡的 tcp echo server , 接入 http request, 並回應
 
### Getting Started
啟動 server 端
```
go run .
```

client 端
``
curl localhost:8887
``


註1：使用 main.go 裡面的註解, 已切換想執行哪種 server
註2：http server 只會 log request 資料
