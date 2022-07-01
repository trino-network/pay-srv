# pay-srv支付服务器

**pay-srv** 是Trino Network的支付服务器，提供数字货币支付接入API。接口文档地址：http://23.254.247.38:8081/

接口采用jwt认证，调用接口需要传人服务器提供的token，例如下面例子中的token为：f140004f-1b51-4abe-a0d4-58beef45d70c

## 支付接入流程如下：

### 1. 创建票据，获取票据ID

创建接口：
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/invoices/create' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer f140004f-1b51-4abe-a0d4-58beef45d70c' \
  -d '{
  "amount": 10000000,
  "jump_url": "https://trino.network/#/order/2022062214060769050",
  "metadata": "e29c1B698F98BdFe6Ca4012dEE6FB350D73E40AE",
  "network": "goerli",
  "notify_url": "https://trino.network/pay/notify",
  "recipient": "0x889fB20cA7D3B2Bbb7bc888F18d1fE68742087Af",
  "title": "支付订单NO：23435435646"
}'
```
返回：

```shell
{
  "invoice_id": "7503cd8a-c921-4368-a32d-6c1d01d86da9",
  "jump_url": "https://xxx.com/payment/page"
}

```
### 2. 跳转到支付页面，使用数字钱包进行支付

支付页面为静态页，目前还没开发，需要用vue实现。
测试时可以使用：https://remix.ethereum.org工具发起支付。

进入该页面首先要调用如下接口获取支付信息：
```shell
curl -X 'GET' \
  'http://23.254.247.38:8080/api/v1/invoices/payment/13af45c9-17ed-4682-b75b-8b8d252cca69' \
  -H 'accept: application/json'
```
返回内容：
```shell
{
  "amount": 100000000000,
  "jump_url": "https://trino.network/#/order/2022062214060769050",
  "network": "goerli",
  "recipient": "0x164B4c5F9ca1A1BC0Fe0Fa39A76a12aAb85683e1",
  "title": "支付订单NO：23435435646"
}
```

### 3. 服务器监听区块链网络，获取支付结果，并发起回调
回调内容：
http header
```shell
trino-pay-Sig 值为数据的hmac签名， key为认证用的token
```
数据：
```json
{"invoice_id": "7503cd8a-c921-4368-a32d-6c1d01d86da9", "status": "paid"}
```


### 4. 客户端收到回调，并获取支付结果
调用如下接口：
```shell

curl -X 'GET' \
  'http://23.254.247.38:8080/api/v1/invoices/13af45c9-17ed-4682-b75b-8b8d252cca69' \
    -H 'Authorization: Bearer f140004f-1b51-4abe-a0d4-58beef45d70c' \
  -H 'accept: application/json'
```
返回：
```json
{
  "metadata": "e29c1B698F98BdFe6Ca4012dEE6FB350D73E40AE",
  "status": "paid",
  "txn_hash": "0xa477aae4aa81df7c7b108bcc65450c7fa36b77f671b2768759e6485acd10f507"
}
```
其中支付信息可以解码metadata获取

PHP插件实现方式可以参考：https://github.com/trino-network/v2board/blob/master/app/Payments/BTCPay.php

#### 支付页面开发可以参考：

示例项目： https://github.com/remote-gildor/hardhat-vue-starter.git  

trino合约代码：https://github.com/trino-network/contracts  

etherjs使用教程：https://learnblockchain.cn/docs/ethers.js/index.html 

hardhat使用教程：https://hardhat.org/tutorial


