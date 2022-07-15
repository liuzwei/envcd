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

package envcd

import (
	"errors"

	"github.com/acmestack/envcd/internal/core/exchanger/etcd"
	"github.com/acmestack/envcd/internal/pkg/config"
	"github.com/acmestack/envcd/internal/pkg/exchanger"
)

// EnvcdConfig the envcd global config
var EnvcdConfig *config.Config

type Envcd struct {
	exchanger   exchanger.Exchanger
	envcdConfig *config.Config
}

// Start envcd by envcd config
//  @param envcdConfig the config for envcd
func Start(envcdConfig *config.Config) *Envcd {
	// show start information & parser config
	envcdConfig.StartInformation()
	EnvcdConfig = envcdConfig
	instance := &Envcd{exchanger: etcd.New()}
	instance.envcdConfig = envcdConfig
	return instance
}

// Put new data to Exchanger by key and value
func (envcd *Envcd) Put(key interface{}, value interface{}) error {
	if envcd == nil || envcd.exchanger == nil {
		return errors.New("IIllegal state for exchanger")
	}
	return envcd.exchanger.Put(key, value)
}

// Get the data from Exchanger by key
func (envcd *Envcd) Get(key interface{}) (interface{}, error) {
	if envcd == nil || envcd.exchanger == nil {
		return nil, errors.New("IIllegal state for exchanger")
	}
	return envcd.exchanger.Get(key)
}

// Find delete the data from Exchanger by key
func (envcd *Envcd) Find(key interface{}) (interface{}, error) {
	if envcd == nil || envcd.exchanger == nil {
		return nil, errors.New("IIllegal state for exchanger")
	}
	return envcd.exchanger.Find(key)
}

// Remove delete the data from Exchanger by key
func (envcd *Envcd) Remove(key interface{}) error {
	if envcd == nil || envcd.exchanger == nil {
		return errors.New("IIllegal state for exchanger")
	}
	return envcd.exchanger.Remove(key)
}
