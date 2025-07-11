// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package agenthealth

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/confmap/xconfmap"

	"github.com/aws/amazon-cloudwatch-agent/extension/agenthealth/handler/stats/agent"
)

func TestLoadConfig(t *testing.T) {
	testCases := []struct {
		id   component.ID
		want component.Config
	}{
		{
			id:   component.NewID(TypeStr),
			want: NewFactory().CreateDefaultConfig(),
		},
		{
			id:   component.NewIDWithName(TypeStr, "1"),
			want: &Config{IsUsageDataEnabled: false, Stats: nil},
		},
		{
			id:   component.NewIDWithName(TypeStr, "2"),
			want: &Config{IsUsageDataEnabled: true, Stats: &agent.StatsConfig{Operations: []string{"ListBuckets"}}},
		},
	}
	for _, testCase := range testCases {
		conf, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
		require.NoError(t, err)
		cfg := NewFactory().CreateDefaultConfig()
		sub, err := conf.Sub(testCase.id.String())
		require.NoError(t, err)
		require.NoError(t, sub.Unmarshal(cfg))

		assert.NoError(t, xconfmap.Validate(cfg))
		assert.Equal(t, testCase.want, cfg)
	}
}
