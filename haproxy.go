package mci

import (
	"errors"
	"fmt"
	"time"

	"github.com/gobars/cmd"
	"github.com/sirupsen/logrus"
)

func (s Settings) createHAProxyConfig() string {
	rConfig := fmt.Sprintf(`
listen %s
  bind 127.0.0.1:%d
  mode tcp
  option tcpka
`, s.ListenName, s.Port)

	for seq, addr := range s.BackAddrs {
		rConfig += fmt.Sprintf("  server s-%d %s check inter 1s\n", seq+1, addr)
	}

	return rConfig
}

func (s Settings) restartHAProxy() error {
	if s.HAProxyRestartShell == "" {
		logrus.Warnf("HAProxyRestartShell is empty")
		return nil
	}

	_, status := cmd.Bash(s.HAProxyRestartShell, cmd.Timeout(5*time.Second), cmd.Buffered(false))
	if status.Error != nil {
		return status.Error
	}

	if status.Exit != 0 {
		return fmt.Errorf("error exiting code %d, stdout:%s, stderr:%s",
			status.Exit, status.Stdout, status.Stderr)
	}

	return nil
}

func (s Settings) overwriteHAProxyCnf(r *Result) error {
	if s.HAProxyCfg == "" {
		return errors.New("HAProxyCfg required")
	}

	if err := FileExists(s.HAProxyCfg); err != nil {
		return err
	}

	logrus.Infof("prepare to overwriteHAProxyCnf %s", r.HAProxy)

	if err := ReplaceFileContent(s.HAProxyCfg,
		`(?is)#\s*`+s.GetAnchorStart()+`(.+)#\s*`+s.GetAnchorEnd(), r.HAProxy); err != nil {
		logrus.Warnf("overwriteHAProxyCnf error: %v", err)
		return err
	}

	logrus.Infof("overwriteHAProxyCnf completed")

	return nil
}
