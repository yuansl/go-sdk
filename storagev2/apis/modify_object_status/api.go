// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 修改文件的存储状态，即禁用状态和启用状态间的的互相转换
package modify_object_status

import credentials "github.com/qiniu/go-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Entry       string                          // 指定目标对象空间与目标对象名称，格式为 <目标对象空间>:<目标对象名称>
	Status      int64                           // `0` 表示启用；`1` 表示禁用
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
