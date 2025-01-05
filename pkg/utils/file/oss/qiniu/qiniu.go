package qiniu

import (
	"ForgeAdmin/pkg/xerr"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type OssQiniuConf struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

func NewQiniuOss(conf *OssQiniuConf) *OssQiniuConf {
	return &OssQiniuConf{
		AccessKey: conf.AccessKey,
		SecretKey: conf.SecretKey,
		Bucket:    conf.Bucket,
		Domain:    conf.Domain,
	}
}

// QiniuByForm 上传表单文件到七牛云
// ctx 上下文
// file 文件
// key 文件名
// fileSize 文件大小
// 返回文件地址
func (o *OssQiniuConf) QiniuByForm(ctx context.Context, formFile io.Reader, key string, fileSize int64) (string, error) {

	if o.AccessKey == "" || o.SecretKey == "" {
		return "", xerr.NewSystemConfError("未配置oss信息,如要使用该功能,请配置后重试")
	}

	putPolicy := storage.PutPolicy{
		Scope: o.Bucket,
	}

	mac := auth.New(o.AccessKey, o.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuabei, // 空间对应的机房
		UseHTTPS:      false,               // 是否使用https域名
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(ctx, &ret, upToken, key, formFile, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return storage.MakePublicURL(o.Domain, key), nil
}
