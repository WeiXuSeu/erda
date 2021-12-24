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
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/personal-workbench/common"
	"github.com/erda-project/erda/modules/admin/component-protocol/components/personal-workbench/i18n"
	"github.com/erda-project/erda/modules/admin/component-protocol/types"
	wb "github.com/erda-project/erda/modules/admin/services/workbench"
)

// GenProjKvColumnInfo show type: DevOps, MSP, DevOps(primary)/MSP
// TODO: operations
func (l *WorkList) GenProjKvColumnInfo(proj apistructs.WorkbenchProjOverviewItem, q wb.IssueUrlQueries) (kvs []list.KvInfo, columns map[string]interface{}) {
	var hovers []list.KvInfo
	columns = make(map[string]interface{})

	switch proj.ProjectDTO.Type {
	case types.ProjTypeDevops:
		// kv issue info
		if proj.IssueInfo == nil {
			proj.IssueInfo = &apistructs.ProjectIssueInfo{}
		}
		kvs = []list.KvInfo{
			// issue expired
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueExpired),
				Value: strconv.FormatInt(int64(proj.IssueInfo.ExpiredIssueNum), 10),
				Operations: map[cptype.OperationKey]cptype.Operation{
					list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
						WithServerDataPtr(list.OpItemClickGotoServerData{
							OpItemBasicServerData: list.OpItemBasicServerData{
								Query: map[string]interface{}{
									common.OpKeyIssueFilterUrlQuery: q.ExpiredQuery,
								},
								Params: map[string]interface{}{
									common.OpKeyProjectID: proj.ProjectDTO.ID,
								},
								Target: common.OpValTargetProjAllIssue,
							},
						}).
						Build(),
				},
			},
			// issue will expire today
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueExpiredToday),
				Value: strconv.FormatInt(int64(proj.IssueInfo.ExpiredOneDayNum), 10),
				Operations: map[cptype.OperationKey]cptype.Operation{
					list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
						WithServerDataPtr(list.OpItemClickGotoServerData{
							OpItemBasicServerData: list.OpItemBasicServerData{
								Query: map[string]interface{}{
									common.OpKeyIssueFilterUrlQuery: q.TodayExpireQuery,
								},
								Params: map[string]interface{}{
									common.OpKeyProjectID: proj.ProjectDTO.ID,
								},
								Target: common.OpValTargetProjAllIssue,
							},
						}).
						Build(),
				},
			},
			// issue undo
			{
				ID:    strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Key:   l.sdk.I18n(i18n.I18nKeyIssueUndo),
				Value: strconv.FormatInt(int64(proj.IssueInfo.TotalIssueNum), 10),
				Operations: map[cptype.OperationKey]cptype.Operation{
					list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
						WithServerDataPtr(list.OpItemClickGotoServerData{
							OpItemBasicServerData: list.OpItemBasicServerData{
								Query: map[string]interface{}{
									common.OpKeyIssueFilterUrlQuery: q.UndoQuery,
								},
								Params: map[string]interface{}{
									common.OpKeyProjectID: proj.ProjectDTO.ID,
								},
								Target: common.OpValTargetProjAllIssue,
							},
						}).
						Build(),
				},
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

		// hover info
		hovers = []list.KvInfo{
			// 项目管理
			{
				ID:   strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Icon: "xiangmuguanli",
				Tip:  "项目管理",
				Operations: map[cptype.OperationKey]cptype.Operation{
					list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
						WithServerDataPtr(list.OpItemClickGotoServerData{
							OpItemBasicServerData: list.OpItemBasicServerData{
								Params: map[string]interface{}{
									common.OpKeyProjectID: proj.ProjectDTO.ID,
								},
								Target: common.OpValTargetProjAllIssue,
							},
						}).
						Build(),
				},
			},
			// 应用开发
			{
				ID:   strconv.FormatUint(proj.ProjectDTO.ID, 10),
				Icon: "yingyongkaifa",
				Tip:  "应用开发",
				Operations: map[cptype.OperationKey]cptype.Operation{
					list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
						WithServerDataPtr(list.OpItemClickGotoServerData{
							OpItemBasicServerData: list.OpItemBasicServerData{
								Params: map[string]interface{}{
									common.OpKeyProjectID: proj.ProjectDTO.ID,
								},
								Target: common.OpValTargetProjApps,
							},
						}).
						Build(),
				},
			},
		}
		columns["hoverIcons"] = hovers

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
