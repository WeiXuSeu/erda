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

package PodStatus

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/cmp"
	cmpcputil "github.com/erda-project/erda/modules/cmp/component-protocol/cputil"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

var steveServer cmp.SteveServer

func (podStatus *PodStatus) Init(ctx servicehub.Context) error {
	server, ok := ctx.Service("cmp").(cmp.SteveServer)
	if !ok {
		return errors.New("failed to init component, cmp service in ctx is not a steveServer")
	}
	steveServer = server
	return podStatus.DefaultProvider.Init(ctx)
}

func (podStatus *PodStatus) Render(ctx context.Context, c *cptype.Component, s cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	if err := podStatus.GenComponentState(c); err != nil {
		return err
	}
	sdk := cputil.SDK(ctx)

	userID := sdk.Identity.UserID
	orgID := sdk.Identity.OrgID

	splits := strings.Split(podStatus.State.PodID, "_")
	if len(splits) != 2 {
		return fmt.Errorf("invalid pod id: %s", podStatus.State.PodID)
	}

	namespace, name := splits[0], splits[1]
	req := &apistructs.SteveRequest{
		UserID:      userID,
		OrgID:       orgID,
		Type:        apistructs.K8SPod,
		ClusterName: podStatus.State.ClusterName,
		Name:        name,
		Namespace:   namespace,
	}

	resp, err := steveServer.GetSteveResource(ctx, req)
	if err != nil {
		return err
	}
	obj := resp.Data()

	fields := obj.StringSlice("metadata", "fields")
	if len(fields) != 9 {
		return fmt.Errorf("pod %s/%s has invalid fields length", namespace, name)
	}
	status := fields[2]
	color, breathing := cmpcputil.ParsePodStatus(status)
	podStatus.Props.Status = color
	podStatus.Props.Text = cputil.I18n(ctx, status)
	podStatus.Props.Breathing = breathing
	podStatus.Transfer(c)
	return nil
}

func (podStatus *PodStatus) GenComponentState(c *cptype.Component) error {
	if c == nil || c.State == nil {
		return nil
	}

	jsonData, err := json.Marshal(c.State)
	if err != nil {
		logrus.Errorf("failed to marshal for eventTable state, %v", err)
		return err
	}
	var state State
	err = json.Unmarshal(jsonData, &state)
	if err != nil {
		logrus.Errorf("failed to unmarshal for eventTable state, %v", err)
		return err
	}
	podStatus.State = state
	return nil
}

func (podStatus *PodStatus) Transfer(c *cptype.Component) {
	c.Props = cputil.MustConvertProps(podStatus.Props)
	c.State = map[string]interface{}{
		"clusterName": podStatus.State.ClusterName,
		"podId":       podStatus.State.PodID,
	}
}

func init() {
	base.InitProviderWithCreator("cmp-dashboard-podDetail", "podStatus", func() servicehub.Provider {
		return &PodStatus{
			Type: "Text",
		}
	})
}
