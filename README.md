# 微信支付v3接口

## 电商收付通
### 商户进件 sub_merchant
| 方法名 | 备注 |
| --- | --- |
ApplymentSubmit | 二级商户进件 
ApplymentQueryByID | 通过申请单ID查询申请状态 
ApplymentQueryByOutRequestNo | 通过业务申请编号查询申请状态 
SettlementModify | 修改结算帐号API 
SettlementQuery | 查询结算账户 
### 普通支付 transaction
| 方法名 | 备注 |
| --- | --- |
JsApiPrepay | JSAPI下单 
PayResultQueryByOutRequestNo | 商户订单号查询交易结果
PayResultQueryByTransactionID | 微信支付订单号查询交易结果
### 退款
| 方法名 | 备注 |
| --- | --- |
RefundApply | 申请退款
QueryRefundByOutRefundNo | 通过商户退款单号查询退款
QueryRefundByID | 通过微信支付退款单号查询退款
### 分账
| 方法名 | 备注 |
| --- | --- |
ProfitShareApply | 请求分账
ProfitShareQuery | 查询分账结果
ProfitReturnApply | 请求分账回退
ProfitReturnQuery | 查询分账回退结果
ProfitShareFinish | 完结分账
ProfitShareUnSplitAmountQuery | 查询订单剩余待分金额

## 公共api
| 方法名 | 备注 |
| --- | --- |
GetCertificates | 获取平台证书列表
MediaUpload | 上传图片

## 示例
```
// 创建微信支付服务商客户端
client := NewMerchantApiClient("xxxx", "xxxx", "apiClientCert", "https://api.mch.weixin.qq.com", 5*time.Second, certMap, "xxx")
file, _ := os.Open("image.png")
resp, err := client.MediaUpload(MediaUploadRequest{file: file})
```
