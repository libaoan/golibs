# README

## 自签名证书（可选）
- 1. 生成私钥
openssl genrsa -out key.pem 2048
- 2. 生成证书
openssl req -new -x509 -key key.pem -out cert.pem -days 1095
- 3. 复制证书文件到webserver/ssl路径下

## 运行测试套
```shell
go test webserver
```