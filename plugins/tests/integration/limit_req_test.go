// Copyright The HTNN Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"mosn.io/moe/pkg/filtermanager"
	"mosn.io/moe/plugins/tests/integration/control_plane"
	"mosn.io/moe/plugins/tests/integration/data_plane"
)

func TestLimitReq(t *testing.T) {
	dp, err := data_plane.StartDataPlane(t, nil)
	if err != nil {
		t.Fatalf("failed to start data plane: %v", err)
		return
	}
	defer dp.Stop()

	tests := []struct {
		name   string
		config *filtermanager.FilterManagerConfig
		run    func(t *testing.T)
	}{
		{
			name: "rps > 1",
			config: control_plane.NewSinglePluinConfig("limit_req", map[string]interface{}{
				"average": 1,
				"period":  "0.1s",
			}),
			run: func(t *testing.T) {
				resp, _ := dp.Head("/echo", nil)
				assert.Equal(t, 200, resp.StatusCode)

				time.Sleep(50 * time.Millisecond) // trigger delay
				now := time.Now()
				resp, _ = dp.Head("/echo", nil)
				pass := time.Since(now)
				assert.Equal(t, 200, resp.StatusCode)
				// delay time plus the req time
				assert.True(t, pass < 55*time.Millisecond, pass)

				time.Sleep(100 * time.Millisecond) // forget the limit
				resp, _ = dp.Head("/echo", nil)
				assert.Equal(t, 200, resp.StatusCode)
				resp, _ = dp.Head("/echo", nil)
				assert.Equal(t, 429, resp.StatusCode)
			},
		},
		{
			name: "rps <= 1",
			config: control_plane.NewSinglePluinConfig("limit_req", map[string]interface{}{
				"average": 1,
				"period":  "60s",
			}),
			run: func(t *testing.T) {
				resp, _ := dp.Head("/echo", nil)
				assert.Equal(t, 200, resp.StatusCode)
				resp, _ = dp.Head("/echo", nil)
				assert.Equal(t, 429, resp.StatusCode)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controlPlane.UseGoPluginConfig(tt.config, dp)
			tt.run(t)
		})
	}
}