package file

import (
	"ForgeAdmin/pkg/utils/file/oss/aliyun"
	"ForgeAdmin/pkg/utils/file/oss/qiniu"
)

type UploadFileConf struct {
	MaxSize int64 `json:",default=100"` // 上传文件最大大小 单位MB
	//Local   LocalConfig `json:",optional"`    // 本地上传配置
	Oss OssConfig `json:",optional"` // oss上传配置
}

type LocalConfig struct {
	Path string // 本地路径
}

type OssConfig struct {
	Qiniu  *qiniu.OssQiniuConf   `json:",optional"` // 七牛云配置
	AliYun *aliyun.OssALiYunConf `json:",optional"` // 阿里云配置
}
