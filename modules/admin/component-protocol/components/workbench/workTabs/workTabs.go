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

package workTabs

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/workbench/common"
	"github.com/erda-project/erda/modules/admin/services/workbench"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Data struct {
	Options []Option `json:"options"`
}

type Operation struct {
	ClientData
}

type ClientData struct {
	Value string `json:"value"`
}

type State struct {
	Value string `json:"value"`
}

type WorkTabs struct {
	base.DefaultProvider
	SDK        *cptype.SDK
	Data       Data `json:"data"`
	wb         *workbench.Workbench
	Operations map[string]Operation `json:"operations"`
	State      State                `json:"state"`
}

func (wt *WorkTabs) GetComponentValue(c cptype.Component) {
	wt.State = wt.GetState()
	wt.Operations = wt.GetOperation()
	return
}

func (wt *WorkTabs) GetState() State {
	return State{Value: common.TabProject}
}

func (wt *WorkTabs) SetState(state cptype.ComponentState) {
	err := common.Transfer(state, wt.State)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func (wt *WorkTabs) GetOperation() map[string]Operation {
	return map[string]Operation{"onChange": {ClientData{Value: ""}}}
}

// SetComponentValue mapping properties to Component
func (wt *WorkTabs) SetComponentValue(c *cptype.Component) error {
	var err error
	if err = common.Transfer(wt.State, &c.State); err != nil {
		return err
	}
	if err = common.Transfer(wt.Data, &c.Data); err != nil {
		return err
	}
	//if err = common.Transfer(l.Props, &c.Props); err != nil {
	//	return err
	//}
	//if err = common.Transfer(l.Operations, &c.Operations); err != nil {
	//	return err
	//}
	return nil
}

func (wt *WorkTabs) GetData(gs *cptype.GlobalStateData) (Data, error) {
	var (
		proData *apistructs.WorkbenchProjOverviewRespData
		appData *apistructs.AppWorkbenchResponseData
		err     error
	)
	wtData := Data{Options: []Option{
		{Value: common.TabProject, Label: wt.SDK.I18n("project")},
		{Value: common.TabApplication, Label: wt.SDK.I18n("app")},
	}}
	apiIdentity := apistructs.Identity{}
	err = common.Transfer(wt.SDK.Identity, apiIdentity)
	if err != nil {
		return wtData, err
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		pageReq := apistructs.PageRequest{}
		proData, err = wt.wb.ListQueryProjWbData(apiIdentity, pageReq, "")
		if err != nil {
			logrus.Errorf("tabs get project list err %v", err)
		}
		wg.Done()
	}()
	go func() {
		appReq := apistructs.ApplicationListRequest{}
		appData, err = wt.wb.ListAppWbData(apiIdentity, appReq, 0)
		if err != nil {
			logrus.Errorf("tabs get app list err %v", err)
		}
		wg.Done()
	}()
	wg.Wait()
	switch wt.State.Value {
	case common.TabProject:
		(*gs)[common.TabData] = proData
	case common.TabApplication:
		(*gs)[common.TabData] = appData
	}
	wtData.Options[0].Label += fmt.Sprintf("(%d)", proData.Total)
	wtData.Options[1].Label += fmt.Sprintf("(%d)", appData.TotalApps)
	return wtData, nil
}

// Render is empty implement.
func (wt *WorkTabs) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	wt.State.Value = common.TabProject
	wt.SDK = cputil.SDK(ctx)
	switch event.Operation {
	case cptype.InitializeOperation:
	case common.EventChangeEventTab:
		err := common.Transfer(wt.Operations, c.Operations)
		if err != nil {
			return err
		}
		wt.State.Value = wt.Operations[common.EventChangeEventTab].Value
	default:
		logrus.Errorf("scenario %v component WorkTabs does not support event %v", scenario, event)
		return nil
	}
	wtData, err := wt.GetData(gs)
	if err != nil {
		return err
	}
	wt.Data = wtData
	err = wt.SetComponentValue(c)
	if err != nil {
		return err
	}
	return nil
}
func init() {
	base.InitProviderWithCreator(common.ScenarioKey, "workTabs", func() servicehub.Provider {
		return &WorkTabs{}
	})
}
