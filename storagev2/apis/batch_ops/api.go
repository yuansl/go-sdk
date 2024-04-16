// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 批量操作意指在单一请求中执行多次（最大限制1000次） 查询元信息、修改元信息、移动、复制、删除、修改状态、修改存储类型、修改生命周期和解冻操作，极大提高对象管理效率。其中，解冻操作仅针对归档存储文件有效
package batch_ops

import (
	"encoding/json"
	credentials "github.com/qiniu/go-sdk/v7/storagev2/credentials"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Operations  []string                        // 单一对象管理指令
}

// 获取 API 所用的响应
type Response struct {
	OperationResponses OperationResponses // 所有管理指令的响应信息
}

// 响应数据
type Data struct {
	Error                       string // 管理指令的错误信息，仅在发生错误时才返回
	Size                        int64  // 对象大小，单位为字节，仅对 stat 指令才有效
	Hash                        string // 对象哈希值，仅对 stat 指令才有效
	MimeType                    string // 对象 MIME 类型，仅对 stat 指令才有效
	Type                        int64  // 对象存储类型，`0` 表示普通存储，`1` 表示低频存储，`2` 表示归档存储，仅对 stat 指令才有效
	PutTime                     int64  // 文件上传时间，UNIX 时间戳格式，单位为 100 纳秒，仅对 stat 指令才有效
	EndUser                     string // 资源内容的唯一属主标识
	RestoringStatus             int64  // 归档存储文件的解冻状态，`2` 表示解冻完成，`1` 表示解冻中；归档文件冻结时，不返回该字段，仅对 stat 指令才有效
	Status                      int64  // 文件状态。`1` 表示禁用；只有禁用状态的文件才会返回该字段，仅对 stat 指令才有效
	Md5                         string // 对象 MD5 值，只有通过直传文件和追加文件 API 上传的文件，服务端确保有该字段返回，仅对 stat 指令才有效
	ExpirationTime              int64  // 文件过期删除日期，UNIX 时间戳格式，文件在设置过期时间后才会返回该字段，仅对 stat 指令才有效
	TransitionToIaTime          int64  // 文件生命周期中转为低频存储的日期，UNIX 时间戳格式，文件在设置转低频后才会返回该字段，仅对 stat 指令才有效
	TransitionToArchiveTime     int64  // 文件生命周期中转为归档存储的日期，UNIX 时间戳格式，文件在设置转归档后才会返回该字段，仅对 stat 指令才有效
	TransitionToDeepArchiveTime int64  // 文件生命周期中转为深度归档存储的日期，UNIX 时间戳格式，文件在设置转归档后才会返回该字段，仅对 stat 指令才有效
	TransitionToArchiveIrTime   int64  // 文件生命周期中转为归档直读存储的日期，UNIX 时间戳格式，文件在设置转归档直读后才会返回该字段，仅对 stat 指令才有效
}

// 管理指令的响应数据
type OperationResponseData = Data
type jsonData struct {
	Error                       string `json:"error,omitempty"`                   // 管理指令的错误信息，仅在发生错误时才返回
	Size                        int64  `json:"fsize,omitempty"`                   // 对象大小，单位为字节，仅对 stat 指令才有效
	Hash                        string `json:"hash,omitempty"`                    // 对象哈希值，仅对 stat 指令才有效
	MimeType                    string `json:"mimeType,omitempty"`                // 对象 MIME 类型，仅对 stat 指令才有效
	Type                        int64  `json:"type,omitempty"`                    // 对象存储类型，`0` 表示普通存储，`1` 表示低频存储，`2` 表示归档存储，仅对 stat 指令才有效
	PutTime                     int64  `json:"putTime,omitempty"`                 // 文件上传时间，UNIX 时间戳格式，单位为 100 纳秒，仅对 stat 指令才有效
	EndUser                     string `json:"endUser,omitempty"`                 // 资源内容的唯一属主标识
	RestoringStatus             int64  `json:"restoreStatus,omitempty"`           // 归档存储文件的解冻状态，`2` 表示解冻完成，`1` 表示解冻中；归档文件冻结时，不返回该字段，仅对 stat 指令才有效
	Status                      int64  `json:"status,omitempty"`                  // 文件状态。`1` 表示禁用；只有禁用状态的文件才会返回该字段，仅对 stat 指令才有效
	Md5                         string `json:"md5,omitempty"`                     // 对象 MD5 值，只有通过直传文件和追加文件 API 上传的文件，服务端确保有该字段返回，仅对 stat 指令才有效
	ExpirationTime              int64  `json:"expiration,omitempty"`              // 文件过期删除日期，UNIX 时间戳格式，文件在设置过期时间后才会返回该字段，仅对 stat 指令才有效
	TransitionToIaTime          int64  `json:"transitionToIA,omitempty"`          // 文件生命周期中转为低频存储的日期，UNIX 时间戳格式，文件在设置转低频后才会返回该字段，仅对 stat 指令才有效
	TransitionToArchiveTime     int64  `json:"transitionToARCHIVE,omitempty"`     // 文件生命周期中转为归档存储的日期，UNIX 时间戳格式，文件在设置转归档后才会返回该字段，仅对 stat 指令才有效
	TransitionToDeepArchiveTime int64  `json:"transitionToDeepArchive,omitempty"` // 文件生命周期中转为深度归档存储的日期，UNIX 时间戳格式，文件在设置转归档后才会返回该字段，仅对 stat 指令才有效
	TransitionToArchiveIrTime   int64  `json:"transitionToArchiveIR,omitempty"`   // 文件生命周期中转为归档直读存储的日期，UNIX 时间戳格式，文件在设置转归档直读后才会返回该字段，仅对 stat 指令才有效
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{Error: j.Error, Size: j.Size, Hash: j.Hash, MimeType: j.MimeType, Type: j.Type, PutTime: j.PutTime, EndUser: j.EndUser, RestoringStatus: j.RestoringStatus, Status: j.Status, Md5: j.Md5, ExpirationTime: j.ExpirationTime, TransitionToIaTime: j.TransitionToIaTime, TransitionToArchiveTime: j.TransitionToArchiveTime, TransitionToDeepArchiveTime: j.TransitionToDeepArchiveTime, TransitionToArchiveIrTime: j.TransitionToArchiveIrTime})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Error = nj.Error
	j.Size = nj.Size
	j.Hash = nj.Hash
	j.MimeType = nj.MimeType
	j.Type = nj.Type
	j.PutTime = nj.PutTime
	j.EndUser = nj.EndUser
	j.RestoringStatus = nj.RestoringStatus
	j.Status = nj.Status
	j.Md5 = nj.Md5
	j.ExpirationTime = nj.ExpirationTime
	j.TransitionToIaTime = nj.TransitionToIaTime
	j.TransitionToArchiveTime = nj.TransitionToArchiveTime
	j.TransitionToDeepArchiveTime = nj.TransitionToDeepArchiveTime
	j.TransitionToArchiveIrTime = nj.TransitionToArchiveIrTime
	return nil
}
func (j *Data) validate() error {
	return nil
}

// 每个管理指令的响应信息
type OperationResponse struct {
	Code int64                 // 响应状态码
	Data OperationResponseData // 响应数据
}
type jsonOperationResponse struct {
	Code int64                 `json:"code"`           // 响应状态码
	Data OperationResponseData `json:"data,omitempty"` // 响应数据
}

func (j *OperationResponse) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonOperationResponse{Code: j.Code, Data: j.Data})
}
func (j *OperationResponse) UnmarshalJSON(data []byte) error {
	var nj jsonOperationResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Code = nj.Code
	j.Data = nj.Data
	return nil
}
func (j *OperationResponse) validate() error {
	if j.Code == 0 {
		return errors.MissingRequiredFieldError{Name: "Code"}
	}
	return nil
}

// 所有管理指令的响应信息
type OperationResponses []OperationResponse

func (j *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.OperationResponses)
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var array OperationResponses
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}
	j.OperationResponses = array
	return nil
}