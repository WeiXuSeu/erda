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

package elasticsearch

//
//import (
//	"context"
//	"fmt"
//	"github.com/erda-project/erda/modules/core/monitor/storekit/elasticsearch/index/loader"
//	"github.com/erda-project/erda/modules/msp/apm/exception/erda-error/storage"
//)
//
//func (p *provider) Count(ctx context.Context, sel *storage.Selector) int64 {
//	indices := p.Loader.Indices(ctx, sel.StartTime, sel.EndTime, loader.KeyPath{
//		Recursive: true,
//	})
//	fmt.Println(indices)
//
//	if len(indices) <= 0 {
//		return 0
//	}
//
//	// do query
//	ctx, cancel := context.WithTimeout(ctx, p.Cfg.QueryTimeout)
//	defer cancel()
//
//	count, err := p.client.Count(indices...).
//		IgnoreUnavailable(true).AllowNoIndices(true).Q("timestamp:[" + string(sel.StartTime) + " TO " + string(sel.EndTime) + "] AND error_id:" + sel.ErrorId).Do(ctx)
//	if err != nil {
//		return 0
//	}
//
//	return count
//}
