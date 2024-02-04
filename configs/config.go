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

package configs

import (
	"github.com/photowey/nemo-examples/pkg/logger"
)

// Config global config instance.
type Config struct {
	Server Server
	App    App
}

// Server `server`(`application-dev.yml`) config node
type Server struct {
	Port int64
}

// App `nemo`(`application.yml`) config node.
type App struct {
	Application Application
	Profiles    Profiles
}

// Application `nemo.application`(`application.yml`) config node.
type Application struct {
	Name    string
	Timeout int64 // seconds
}

// Profiles `nemo.profiles`(`application.yml`) config node.
type Profiles struct {
	Active []string
}

// Health `nemo.health`(`application.yml`) config node.
type Health struct {
	Api string
}

func Timeout() uint32 {
	return uint32(10)
}

func Host() string {
	return "127.0.0.1:9527"
}

func Logger() logger.Config {
	return logger.Config{}
}
