package mci

import (
	"fmt"
	"os"
	"strings"
)

// CheckHAProxyServers 检查
func (s Settings) CheckHAProxyServers() {
	if s.ValidateAndSetDefault(SetDefault) != nil {
		os.Exit(1)
	}

	linesInFile, err := SearchPatternLinesInFile(s.HAProxyCfg,
		`(?is)`+s.GetAnchorStart()+`(.+)`+s.GetAnchorEnd(),
		`(?i)server\s+\S+\s(\d+(\.\d+){3}(:\d+)?)`)
	if err != nil {
		fmt.Printf("SearchPatternLinesInFile error %v\n", err)
		return
	}

	fmt.Println(strings.Join(linesInFile, "\n"))
}
