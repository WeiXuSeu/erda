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
)

// GenAppKvInfo show: mr num, runtime num
func (l *WorkList) GenAppKvInfo(app apistructs.AppWorkBenchItem) (kvs []list.KvInfo) {
	kvs = []list.KvInfo{
		// mr count
		{
			ID:    strconv.FormatUint(app.ID, 10),
			Key:   l.sdk.I18n(i18n.I18nKeyMrCount),
			Value: strconv.FormatInt(int64(app.AppOpenMrNum), 10),
		},
		// service count
		{
			ID:    strconv.FormatUint(app.ID, 10),
			Key:   l.sdk.I18n(i18n.I18nKeyRuntimeCount),
			Value: strconv.FormatInt(int64(app.AppOpenMrNum), 10),
		},
	}
	return
}
