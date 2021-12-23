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
	"github.com/erda-project/erda/modules/admin/component-protocol/components/personal-workbench/common"
	"github.com/erda-project/erda/modules/admin/services/workbench"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type WorkCards struct {
	base.DefaultProvider
	impl.DefaultCard
	filterReq apistructs.IssuePagingRequest
	State     State `json:"state"`
	bdl       *bundle.Bundle
	wb        *workbench.Workbench
}

func (wc *WorkCards) RegisterCardListGotoOp(opData cardlist.OpCardListGoto) (opFunc cptype.OperationFunc) {
	//TODO implement me
	return func(sdk *cptype.SDK) {}
}

func (wc *WorkCards) RegisterCardIconGotoOp(opData cardlist.OpCardListIconGoto) (opFunc cptype.OperationFunc) {
	//TODO implement me
	return func(sdk *cptype.SDK) {}
}

func (wc *WorkCards) RegisterCardListStarOp(opData cardlist.OpCardListStar) (opFunc cptype.OperationFunc) {
	//TODO implement me
	return func(sdk *cptype.SDK) {}
}

type State struct {
	TabName string `json:"tabName"`
}

func (wc *WorkCards) RegisterInitializeOp() (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) {
		wc.LoadList(sdk)
	}
}

func (wc *WorkCards) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return wc.RegisterInitializeOp()
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

func (wc *WorkCards) LoadList(sdk *cptype.SDK) {
	tabStr := ""
	apiIdentity := apistructs.Identity{}
	err := common.Transfer(sdk.Identity, apiIdentity)
	if err != nil {
		return
	}
	data := cardlist.Data{}
	if tab, ok := (*sdk.GlobalState)[common.WorkTabKey]; !ok {
		tabStr = wc.State.TabName
	} else {
		tabStr = tab.(string)
	}
	switch tabStr {
	case apistructs.WorkbenchItemApp.String():
		data.Title = sdk.I18n("star application")
		apps, err := wc.wb.ListSubAppWbData(apiIdentity, 0)
		if err != nil {
			return
		}
		data.TitleSummary = fmt.Sprintf("%d", len(apps.List))
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
	case apistructs.WorkbenchItemProj.String():
		data.Title = sdk.I18n("star project")
		projects, err := wc.wb.ListSubProjWbData(apiIdentity)
		if err != nil {
			return
		}
		data.TitleSummary = fmt.Sprintf("%d", len(projects.List))
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
	wc.StdDataPtr = &data
}

func init() {
	base.InitProviderWithCreator(common.ScenarioKey, "workCards", func() servicehub.Provider {
		return &WorkCards{}
	})
}
