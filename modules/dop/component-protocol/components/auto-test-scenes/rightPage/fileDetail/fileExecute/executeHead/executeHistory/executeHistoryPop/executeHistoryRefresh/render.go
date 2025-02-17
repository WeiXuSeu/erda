// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package executeHistoryRefresh

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type ComponentAction struct {
	base.DefaultProvider
}

func (ca *ComponentAction) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	c.Operations = map[string]interface{}{
		"click": map[string]interface{}{
			"key":    "toNewRecord",
			"reload": true,
		},
	}
	c.Props = map[string]interface{}{
		"text":       "返回最新记录",
		"type":       "text",
		"prefixIcon": "shuaxin",
		"style":      map[string]interface{}{"width": 140},
	}
	c.State = map[string]interface{}{"pageNo": 1}
	return nil
}

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "executeHistoryRefresh",
		func() servicehub.Provider { return &ComponentAction{} })
}
