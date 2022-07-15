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

package plugin

import "github.com/acmestack/envcd/internal/pkg/executor"

// Chain the executor chain
// this is openapi chain, when http or client request into openapi, construct this chain
type Chain struct {
	executors []executor.Executor
	index     int
}

func New(executors ...executor.Executor) *Chain {
	return &Chain{executors: executors, index: 0}
}

// Execute chain executor
//  @param context chain context
func (chain *Chain) Execute(context interface{}) (ret interface{}, err error) {
	if chain == nil || chain.executors == nil || len(chain.executors) == 0 {
		return
	}
	if chain.index < len(chain.executors) {
		current := chain.executors[chain.index]
		chain.index++
		if current.Skip(context) {
			return chain.Execute(context)
		}
		// todo log
		// todo data
		return current.Execute(context, nil, chain)
	}
	return nil, nil
}