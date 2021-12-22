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

package workCards

import (
	"fmt"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/components/cardlist"
	"github.com/erda-project/erda-infra/providers/component-protocol/components/cardlist/impl"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/workbench/common"
	"github.com/erda-project/erda/modules/admin/services/workbench"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

const (
	stateKeyIssueViewGroupValue         = "issueViewGroupValue"         // kanban
	stateKeyIssueViewGroupChildrenValue = "issueViewGroupChildrenValue" // kanban: status
	stateKeyIssueFilterConditions       = "filterConditions"            // apistructs.IssuePagingRequest
)

type CardList struct {
	base.DefaultProvider
	impl.DefaultCard
	filterReq apistructs.IssuePagingRequest

	bdl *bundle.Bundle
	wb  *workbench.Workbench
}

func (c *CardList) RegisterInitializeOp() (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) {
		c.LoadList(sdk)
	}
}

func (c *CardList) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return c.RegisterInitializeOp()
}

type TextMeta struct {
	MainText string
	SubText  string
	// 点击跳转
	Operations map[string]common.Operation
}

func getAppExtra(sdk *cptype.SDK, app apistructs.AppWorkBenchItem) cptype.Extra {
	extra := cptype.Extra{Extra: make(cptype.ExtraMap)}
	extra.Extra["textMeta"] = []TextMeta{
		{
			MainText: fmt.Sprintf("%d", app.AppOpenMrNum),
			SubText:  "MR " + sdk.I18n("Count"),
			Operations: map[string]common.Operation{"clickGoto": common.Operation{
				JumpOut: false,
				Target:  "",
				Query:   nil,
				Params:  nil,
			}},
		},
		{
			MainText: fmt.Sprintf("%d", app.AppRuntimeNum),
			SubText:  "Runtime " + sdk.I18n("Count"),
			Operations: map[string]common.Operation{"clickGoto": common.Operation{
				JumpOut: false,
				Target:  "",
				Query:   nil,
				Params:  nil,
			}},
		},
	}
	return extra
}

func getProjectExtra(sdk *cptype.SDK) cptype.Extra {
	return cptype.Extra{}
}

func (c *CardList) LoadList(sdk *cptype.SDK) {
	apiIdentity := apistructs.Identity{}
	err := common.Transfer(apiIdentity, sdk.Identity)
	if err != nil {
		return
	}
	data := cardlist.Data{}
	if tab, ok := (*sdk.GlobalState)[common.TabKey]; ok {
		switch tab {
		case common.TabApplication:
			apps, err := c.wb.ListSubAppWbData(apiIdentity, 0)
			if err != nil {
				return
			}
			for i, app := range apps.List {
				data.Cards = append(data.Cards, cardlist.Card{
					ID:     fmt.Sprintf("%d", i),
					ImgURL: "",
					Icon:   "",
					Title:  app.Name,
					Labels: nil,
					Star:   true,
					Extra:  getAppExtra(sdk, app),
				})
			}
		case common.TabProject:
			projects, err := c.wb.ListSubProjWbData(apiIdentity)
			if err != nil {
				return
			}
			for i, _ := range projects.List {
				data.Cards = append(data.Cards, cardlist.Card{
					ID:     fmt.Sprintf("%d", i),
					ImgURL: "",
					Icon:   "",
					Title:  "",
					Labels: nil,
					Star:   true,
					Extra:  getProjectExtra(sdk),
				})
			}
		}
	}
	c.StdDataPtr = &data
}

func init() {
	base.InitProviderWithCreator(common.ScenarioKey, "workCards", func() servicehub.Provider {
		return &CardList{}
	})
}
