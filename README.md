# OQ

目前功能
1. 極簡的 http server ( 實作 ServeHTTP Interface, 印出 request 內容 ) 
2. 極簡的 tcp echo server , 接入 http request, 並回應
3. 切換 http/tcp 的方式: 使用 main.go 裡面的註解 

### Getting Started
啟動 server 端
```bash
go run .
```

client 端
```bash
curl localhost:8887
```

