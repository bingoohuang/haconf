package mci

import (
	"github.com/bingoohuang/goreflect"
	"github.com/creasty/defaults"
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

// Settings 表示初始化化MySQL集群所需要的参数结构
type Settings struct {
	ListenName          string   `validate:"empty=false"`              // HAProxy配置文件中listen名称
	Port                int      `validate:"gte=0"`                    // HAProxy前端端口号
	BackAddrs           []string `validate:"empty=false"`              // 后端地址
	HAProxyCfg          string   `default:"/etc/haproxy.cfg"`          // HAProxy配置文件地址
	HAProxyRestartShell string   `default:"systemctl restart haproxy"` // HAProxy重启命令

	Debug bool // 测试模式，只打印HAProxy配置, 不实际执行
}

func (s Settings) GetAnchorStart() string { return s.ListenName + "Start" }
func (s Settings) GetAnchorEnd() string   { return s.ListenName + "End" }

// Result 表示初始化结果
type Result struct {
	HAProxy string
}

type SettingsOption int

const (
	Validate SettingsOption = iota
	SetDefault
)

func (s *Settings) ValidateAndSetDefault(options ...SettingsOption) error {
	if goreflect.SliceContains(options, Validate) {
		if err := validate.Validate(s); err != nil {
			logrus.Errorf("error %v", err)
			return err
		}
	}

	if goreflect.SliceContains(options, SetDefault) {
		if err := defaults.Set(s); err != nil {
			logrus.Errorf("defaults set %v", err)
			return err
		}
	}

	return nil
}
