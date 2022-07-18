/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package context

import (
	"net/http"

	"github.com/acmestack/envcd/pkg/entity/data"
)

// Context for peer request
type Context struct {
	Uri         string
	Method      string
	Headers     map[string]interface{}
	ContentType string
	Parameters  map[string]interface{}
	Cookies     map[string]interface{}
	Body        interface{}
	Action      func() (*data.EnvcdResult, error)
	HttpRequest *http.Request
}
