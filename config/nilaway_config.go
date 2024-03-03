// This must be package main
package main

import (
	"errors"
	"fmt"
	"go.uber.org/nilaway/config"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	if _, ok := conf.(map[string]any); !ok {
		return nil, errors.New("expected map[string]any")
	}

	for k, v := range conf.(map[string]any) {
		err := config.Analyzer.Flags.Set(k, fmt.Sprintf("%s", v))
		if err != nil {
			return nil, err
		}
	}

	return []*analysis.Analyzer{config.Analyzer}, nil
}
