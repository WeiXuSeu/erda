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

package scenesSetInfo

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/auto-test-scenes/common/gshelper"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "scenesSetInfo",
		func() servicehub.Provider { return &ScenesSetInfo{} })
}

func (s *ScenesSetInfo) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) (err error) {
	if err = s.initFromProtocol(ctx, c, event, gs); err != nil {
		return
	}

	setID := s.gsHelper.GetGlobalSelectedSetID()
	s.Props = map[string]interface{}{
		"visible": func() bool {
			return setID != 0 && s.gsHelper.GetGlobalActiveConfig() == gshelper.SceneSetConfigKey
		}(),
	}
	if setID == 0 {
		return nil
	}
	rsp, err := s.atTestPlan.GetSceneSet(setID)
	if err != nil {
		return err
	}

	creator, err := s.bdl.GetCurrentUser(rsp.CreatorID)
	if err == nil {
		rsp.CreatorID = creator.Nick
	}

	if rsp.UpdaterID != "" {
		updaterName, err := s.bdl.GetCurrentUser(rsp.UpdaterID)
		if err == nil {
			rsp.UpdaterID = updaterName.Nick
		}
	}

	s.Data.SceneSet = *rsp
	s.Data.CreateATString = rsp.CreatedAt.Format("2006-01-02 15:04:05")
	s.Data.UpdateATString = rsp.UpdatedAt.Format("2006-01-02 15:04:05")

	s.Props["fields"] = []PropColumn{
		{
			Label:      "名称",
			ValueKey:   "name",
			RenderType: "ellipsis",
		},
		{
			Label:      "描述",
			ValueKey:   "description",
			RenderType: "ellipsis",
		},
		{
			Label:    "创建人",
			ValueKey: "creatorID",
		},
		{
			Label:    "创建时间",
			ValueKey: "createAtString",
		},
		{
			Label:    "更新人",
			ValueKey: "updaterID",
		},
		{
			Label:    "更新时间",
			ValueKey: "updateAtString",
		},
	}

	s.RenderProtocol(c, gs)
	return nil
}

func (s *ScenesSetInfo) RenderProtocol(c *cptype.Component, g *cptype.GlobalStateData) {
	if c.Data == nil {
		d := make(cptype.ComponentData)
		c.Data = d
	}
	(*c).Data["data"] = s.Data
	c.Props = s.Props
}
