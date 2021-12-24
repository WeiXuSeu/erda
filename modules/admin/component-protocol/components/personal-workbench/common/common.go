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

package common

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2/json"

	"github.com/erda-project/erda-infra/providers/component-protocol/components/list"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/personal-workbench/i18n"
)

const (
	ScenarioKey = "personal-workbench"

	// FilterNameKey list query key
	FilterNameKey = "FilterName"
	// WorkTabKey work tab switch, e.g: project, app
	WorkTabKey = "workTabKey"

	// TabData load data from tabs transfer with global state.
	TabData = "tabData"

	AppService     = "appService"
	ProjectService = "projectService"

	EventChangeEventTab = "onChange"

	// DevOpsStatus titleState status; error(red), success(green), processing(blue), warning(yellow), default(gray)
	DevOpsStatus = "processing"
	MspStatus    = "warning"

	MspProject    = "MSP"
	DevOpsProject = "DevOps"

	// OpKeyProjectID  operation related keys
	OpKeyProjectID = "projectId"

	// OpValTargetProjAllIssue target related keys
	OpValTargetProjAllIssue = "projectAllIssue"
)

var (
	PtrRequiredErr     = errors.New("b must be a pointer")
	NothingToBeDoneErr = errors.New("nothing to be done")
)

type Operation struct {
	JumpOut bool                   `json:"jumpOut"`
	Target  string                 `json:"target"`
	Query   map[string]interface{} `json:"query"`
	Params  map[string]interface{} `json:"params"`
}

// Transfer transfer a to b with json, kind of b must be pointer
func Transfer(a, b interface{}) error {
	if reflect.ValueOf(b).Kind() != reflect.Ptr {
		return PtrRequiredErr
	}
	if a == nil {
		return NothingToBeDoneErr
	}
	aBytes, err := json.Marshal(a)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(aBytes, b); err != nil {
		return err
	}
	return nil
}

func ToProjTitleState(tp string) (list.StateInfo, bool) {
	switch tp {
	case MspProject:
		return list.StateInfo{Text: i18n.I18nKeyMspProject, Status: MspStatus}, true
	case DevOpsProject:
		return list.StateInfo{Text: i18n.I18nKeyDevOpsProject, Status: DevOpsStatus}, true
	default:
		logrus.Warnf("wrong project type: %v", tp)
		return list.StateInfo{}, false
	}
}
