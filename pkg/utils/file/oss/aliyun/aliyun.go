package aliyun

import (
	"Forge-Admin-core/pkg/xerr"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

type (
	OssALiYunConf struct {
		Bucket          string
		Endpoint        string
		AccessKeyId     string
		AccessKeySecret string
		Domain          string
	}

	OssHandler struct {
		Bucket *oss.Bucket
		domain string
	}
)

func MustNewALiYunOss(conf *OssALiYunConf) *OssHandler {

	return &OssHandler{
		Bucket: auth(conf),
		domain: conf.Domain,
	}
}

func auth(conf *OssALiYunConf) *oss.Bucket {
	client, err := oss.New(conf.Endpoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		panic("连接阿里云OSS服务器失败,err:" + err.Error())
	}
	bucket, err := client.Bucket(conf.Bucket)
	if err != nil {
		panic("指定的bucket错误或不存在,err:" + err.Error())
	}
	return bucket
}

func (h *OssHandler) UploadByForm(formFile io.Reader, key string) (string, error) {

	if err := h.Bucket.PutObject(key, formFile); err != nil {
		logx.Error(errors.Wrapf(err, ""))
		return "", xerr.NewCustomError(xerr.RequestParamsErr.Code(), "上传文件失败,可在控制台查看详细堆栈信息")
	}

	return makePublicURL(h.domain, key), nil

}

func makePublicURL(domain, key string) string {
	return domain + key
}
