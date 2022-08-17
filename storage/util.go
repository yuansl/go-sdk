package storage

import (
	"context"
	api "github.com/qiniu/go-sdk/v7"
	"github.com/qiniu/go-sdk/v7/internal/hostprovider"
	"strings"
	"time"
)

// ParsePutTime 提供了将PutTime转换为 time.Time 的功能
func ParsePutTime(putTime int64) (t time.Time) {
	t = time.Unix(0, putTime*100)
	return
}

// IsContextExpired 检查分片上传的ctx是否过期，提前一天让它过期
// 因为我们认为如果断点继续上传的话，最长需要1天时间
func IsContextExpired(blkPut BlkputRet) bool {
	if blkPut.Ctx == "" {
		return false
	}
	target := time.Unix(blkPut.ExpiredAt, 0).AddDate(0, 0, -1)
	now := time.Now()
	return now.After(target)
}

func shouldUploadRetry(err error) bool {
	if err == nil {
		return false
	}

	errInfo, ok := err.(*ErrorInfo)
	if !ok {
		return true
	}

	return errInfo.Code > 499 && errInfo.Code < 600 && errInfo.Code != 573 && errInfo.Code != 579
}

func shouldFreezeHost(err error) bool {
	if err == nil {
		return false
	}

	errInfo, ok := err.(*ErrorInfo)
	if !ok {
		// 转化不成功，非服务问题
		return false
	}

	return errInfo.Code > 499 && errInfo.Code < 600 && errInfo.Code != 573 && errInfo.Code != 579
}

func doUploadAction(hostProvider hostprovider.HostProvider, retryMax int, freezeDuration int, action func(host string) error) error {
	for {
		host, err := hostProvider.Provider()
		if err != nil {
			return api.NewError(ErrMaxUpRetry, err.Error())
		}

		for i := 0; ; i++ {
			err = action(host)

			// 请求成功
			if err == nil {
				return nil
			}

			// 取消
			if isCancelErr(err) {
				return err
			}

			// 如果需要则冻结当前 host
			if shouldFreezeHost(err) {
				_ = hostProvider.Freeze(host, err, freezeDuration)
			}

			// 不可重试错误
			if !shouldUploadRetry(err) {
				return err
			}

			// 单个 host 重试达到最大值
			if i >= retryMax {
				break
			}
		}
	}
}

func isCancelErr(err error) bool {
	if err == context.Canceled {
		return true
	}

	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "context canceled")
}
