// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/qiniu/go-sdk/v7/auth"
	getbucketinfos "github.com/qiniu/go-sdk/v7/storagev2/apis/get_bucket_infos"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
	httpclient "github.com/qiniu/go-sdk/v7/storagev2/http_client"
	uplog "github.com/qiniu/go-sdk/v7/storagev2/internal/uplog"
	region "github.com/qiniu/go-sdk/v7/storagev2/region"
	uptoken "github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type innerGetBucketInfosRequest getbucketinfos.Request

func (query *innerGetBucketInfosRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Region != "" {
		allQuery.Set("region", query.Region)
	}
	if query.Statistics {
		allQuery.Set("fs", strconv.FormatBool(query.Statistics))
	}
	return allQuery, nil
}
func (request *innerGetBucketInfosRequest) getAccessKey(ctx context.Context) (string, error) {
	if request.Credentials != nil {
		if credentials, err := request.Credentials.Get(ctx); err != nil {
			return "", err
		} else {
			return credentials.AccessKey, nil
		}
	}
	return "", nil
}

type GetBucketInfosRequest = getbucketinfos.Request
type GetBucketInfosResponse = getbucketinfos.Response

// 获取用户所有存储空间信息
func (storage *Storage) GetBucketInfos(ctx context.Context, request *GetBucketInfosRequest, options *Options) (*GetBucketInfosResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetBucketInfosRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	var pathSegments []string
	pathSegments = append(pathSegments, "v2", "bucketInfos")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	bucketName := options.OverwrittenBucketName
	var objectName string
	uplogInterceptor, err := uplog.NewRequestUplog("getBucketInfos", bucketName, objectName, func() (string, error) {
		credentials := innerRequest.Credentials
		if credentials == nil {
			credentials = storage.client.GetCredentials()
		}
		putPolicy, err := uptoken.NewPutPolicy(bucketName, time.Now().Add(time.Hour))
		if err != nil {
			return "", err
		}
		return uptoken.NewSigner(putPolicy, credentials).GetUpToken(ctx)
	})
	if err != nil {
		return nil, err
	}
	req := httpclient.Request{Method: "POST", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, BufferResponse: true}
	if options.OverwrittenEndpoints == nil && options.OverwrittenRegion == nil && storage.client.GetRegions() == nil {
		query := storage.client.GetBucketQuery()
		if query == nil {
			bucketHosts := httpclient.DefaultBucketHosts()
			if options.OverwrittenBucketHosts != nil {
				req.Endpoints = options.OverwrittenBucketHosts
			} else {
				req.Endpoints = bucketHosts
			}
		}
		if query != nil {
			var accessKey string
			var err error
			if accessKey, err = innerRequest.getAccessKey(ctx); err != nil {
				return nil, err
			}
			if accessKey == "" {
				if credentialsProvider := storage.client.GetCredentials(); credentialsProvider != nil {
					if creds, err := credentialsProvider.Get(ctx); err != nil {
						return nil, err
					} else if creds != nil {
						accessKey = creds.AccessKey
					}
				}
			}
			if accessKey != "" && bucketName != "" {
				req.Region = query.Query(accessKey, bucketName)
			}
		}
	}
	var respBody GetBucketInfosResponse
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}