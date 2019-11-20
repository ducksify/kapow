/*
 * Copyright 2019 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/BBVA/kapow/internal/server/srverrors"
)

// GetReason returns the reason phrase part of an HTTP response
func GetReason(r *http.Response) string {
	if i := strings.IndexByte(r.Status, ' '); i != -1 {
		return r.Status[i+1:]
	}
	return ""
}

// GetReasonFromBody returns the reason phrase embedded within the JSON error
// body, or an error if no reason can be extracted
func GetReasonFromBody(r *http.Response) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("error reading response's body")
	}

	reason := &srverrors.ServerErrMessage{}
	err = json.Unmarshal(body, reason)
	if err != nil {
		return "", errors.New("error unmarshaling JSON")
	}

	if reason.Reason == "" {
		return "", errors.New("no reason")
	}

	return reason.Reason, nil
}
