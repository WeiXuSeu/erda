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

package workList

import (
	"strconv"

	"github.com/erda-project/erda-infra/providers/component-protocol/components/list"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/personal-workbench/i18n"
	"github.com/erda-project/erda/modules/admin/component-protocol/types"
)

// GenProjKvInfo show type: DevOps, MSP, DevOps(primary)/MSP
// TODO: operations
func (l *ProjAppList) GenProjKvInfo(proj apistructs.WorkbenchProjOverviewItem) (kvs []list.KvInfo) {

	switch proj.ProjectDTO.Type {
	case types.ProjTypeDevops:
		if proj.IssueInfo == nil {
			proj.IssueInfo = &apistructs.ProjectIssueInfo{}
		}
		kvs = []list.KvInfo{
			// issue expired
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueExpired),
				Value: strconv.FormatInt(int64(proj.IssueInfo.ExpiredIssueNum), 10),
			},
			// issue will expire today
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueExpiredToday),
				Value: strconv.FormatInt(int64(proj.IssueInfo.ExpiredOneDayNum), 10),
			},
			// issue undo
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueUndo),
				Value: strconv.FormatInt(int64(proj.IssueInfo.TotalIssueNum), 10),
			},
		}
		if proj.StatisticInfo != nil {
			// msp alert info
			altKv := list.KvInfo{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyMspLast24HAlertCount),
				Value: strconv.FormatInt(proj.StatisticInfo.Last24HAlertCount, 10),
			}
			kvs = append(kvs, altKv)
		}
	case types.ProjTypeMSP:
		if proj.StatisticInfo == nil {
			return
		}
		kvs = []list.KvInfo{
			// msp service info
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyMspServiceCount),
				Value: strconv.FormatInt(proj.StatisticInfo.ServiceCount, 10),
			},
			// msp alert info
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyMspLast24HAlertCount),
				Value: strconv.FormatInt(proj.StatisticInfo.Last24HAlertCount, 10),
			},
		}
	}
	return
}
