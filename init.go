package mci

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Exec 执行HAProxy配置更新
func (s Settings) Exec() (r Result, err error) {
	if s.ValidateAndSetDefault(Validate, SetDefault) != nil {
		os.Exit(1)
	}

	r.HAProxy = s.createHAProxyConfig()

	if s.Debug {
		logrus.Infof("HAProxy:%s", r.HAProxy)

		return r, err
	}

	if err := s.overwriteHAProxyCnf(&r); err != nil {
		return r, err
	}

	if err := s.restartHAProxy(); err != nil {
		return r, err
	}

	return r, err
}
