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

package uc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apierrors"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/auth"
	"github.com/erda-project/erda/modules/pkg/user"
	"github.com/erda-project/erda/pkg/discover"
	"github.com/erda-project/erda/pkg/http/httpclient"
	"github.com/erda-project/erda/pkg/http/httpserver"
	"github.com/erda-project/erda/pkg/http/httpserver/errorresp"
	"github.com/erda-project/erda/pkg/strutil"
	"github.com/erda-project/erda/pkg/ucauth"
)

var UC_USER_FREEZE = apis.ApiSpec{
	Path:         "/api/users/<userID>/actions/freeze",
	Scheme:       "http",
	Method:       "PUT",
	Custom:       freezeUser,
	CheckLogin:   true,
	CheckToken:   true,
	RequestType:  apistructs.UserFreezeRequest{},
	ResponseType: apistructs.UserFreezeResponse{},
	IsOpenAPI:    true,
	Doc:          "summary: 用户冻结",
}

func freezeUser(w http.ResponseWriter, r *http.Request) {
	operatorID, err := user.GetUserID(r)
	if err != nil {
		apierrors.ErrAdminUser.NotLogin().Write(w)
		return
	}

	if err := checkPermission(operatorID, apistructs.UpdateAction); err != nil {
		errorresp.ErrWrite(err, w)
		return
	}
	token, err := auth.GetDiceClientToken()
	if err != nil {
		logrus.Errorf("failed to get token: %v", err)
		apierrors.ErrFreezeUser.InternalError(err).
			Write(w)
		return
	}

	// check login & permission
	_, err = mustManageUsersPerm(r, apierrors.ErrFreezeUser)
	if err != nil {
		errorresp.ErrWrite(err, w)
		return
	}

	// get req
	userID := strutil.Split(r.URL.Path, "/", true)[2]
	logrus.Debugf("to freeze userID: %v", userID)

	// handle
	if err := handleFreezeUser(userID, operatorID.String(), token); err != nil {
		errorresp.ErrWrite(err, w)
		return
	}
	httpserver.WriteData(w, nil)
}

func handleFreezeUser(userID, operatorID string, token ucauth.OAuthToken) error {
	if token.TokenType == ucauth.OryCompatibleClientId {
		return ucauth.ChangeUserState(token.AccessToken, userID, ucauth.UserInActive)
	}

	var resp struct {
		Success bool   `json:"success"`
		Result  bool   `json:"result"`
		Error   string `json:"error"`
	}
	r, err := httpclient.New().Put(discover.UC()).
		Path("/api/user/admin/freeze/"+userID).
		Header("Authorization", strutil.Concat("Bearer ", token.AccessToken)).
		Header("operatorId", operatorID).
		Do().JSON(&resp)
	if err != nil {
		return apierrors.ErrFreezeUser.InternalError(err)
	}
	if !r.IsOK() {
		return apierrors.ErrFreezeUser.InternalError(fmt.Errorf("internal status code: %v", r.StatusCode()))
	}
	if !resp.Success {
		return apierrors.ErrFreezeUser.InternalError(errors.New(resp.Error))
	}
	return nil
}
