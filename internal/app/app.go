/*
 * Copyright © 2024 the original author or authors.
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

package app

import (
	"github.com/photowey/nemo-examples/configs"
	"github.com/photowey/nemo-examples/pkg/logger"
	"github.com/photowey/perrors"
)

func Start(conf string) error {
	err := loadConfig(conf)
	if err != nil {
		return err
	}

	err = dummyInitLogger()
	if err != nil {
		return err
	}

	return nil
}

func loadConfig(conf string) error {
	// TODO
	// load config by nemo.

	return nil
}

func dummyInitLogger() error {
	if err := logger.Init(configs.Logger()); err != nil {
		return perrors.Errorf("nemoexp: init logger failed: %v", err)
	}

	return nil
}

func Close() {
}
