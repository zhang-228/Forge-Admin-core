package config

import (
	"Forge-Admin-core/pkg/utils/email"
	"Forge-Admin-core/pkg/utils/file"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"os"
	"strings"
)

type SystemConf struct {
	JwtAuth struct {
		AccessSecret string // jwt密钥
		AccessExpire int64  // jwt过期时间 有效期单位: 秒
	}
	Mysql struct {
		Datasource string // mysql datasource
	}
	RDB    redis.RedisConf     `json:",optional"` // redis配置
	Email  email.Email         `json:",optional"` // 邮箱配置(用于接口错误推送、硬件监控)
	Upload file.UploadFileConf `json:",optional"` // 上传配置

	//MultipleLogins string              `json:",optional"` // 多登录拦截
	//Collection collection.CollectorConf `json:",optional"` // 数据采集
}

func (c *SystemConf) InitMysql() {
	db, err := sql.Open("mysql", c.Mysql.Datasource)
	if err != nil {
		panic("init mysql database failed," + err.Error())
	}
	if err = db.Ping(); err != nil {
		db.Close()
		panic("ping mysql database failed," + err.Error())
	}
	defer db.Close()

	sqlFile, err := os.ReadFile("config/zero-admin.sql")
	if err != nil {
		panic("read mysql sql file failed," + err.Error())
	}

	sqlArr := strings.Split(string(sqlFile), ";")
	for _, v := range sqlArr {
		if len(v) > 0 {
			_, err = db.Exec(v)
			if err != nil {
				panic("init mysql database failed," + err.Error())
			}
		}
	}
	fmt.Println("init mysql database success")
}
