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

package main

import (
	"math/rand"
	"os"
	"time"

	App "github.com/photowey/nemo-examples/cmd/app"
	"github.com/photowey/nemo-examples/internal/version"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	_ = os.Setenv("app.version", version.Version())
}

func main() {
	App.Run()
}
