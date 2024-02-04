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
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/photowey/nemo-examples/configs"
	"github.com/photowey/nemo-examples/internal/api/v1/controller"
	"github.com/photowey/nemo-examples/internal/app"
	"github.com/photowey/nemo-examples/pkg/expression"
	"github.com/photowey/nemo-examples/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	SystemExit                = 1
	ErrChannelSize            = 1
	HealthCheckTimeout uint32 = 15
)

var start = &cobra.Command{
	Use:   "start",
	Short: "Start app",
	Long:  `Start the nemo example project application`,
	Run: func(cmd *cobra.Command, args []string) {
		quit := make(chan os.Signal, SystemExit)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		crush := make(chan error, ErrChannelSize)

		timeout := expression.TrinaryOperationUInt32(configs.Timeout() > 0, configs.Timeout(), HealthCheckTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second) // ping
		defer cancel()

		ctrl := controller.NewController(configs.Host())
		ctrl.Run(ctx, crush)

		select {
		case <-quit:
			clear()
			logger.Info("nemoexp: shutting down server...")
			return
		case err := <-crush:
			logger.Errorf("nemoexp: server crush, and quit: %v", err)
			clear()
			return
		}
	},
}

func clear() {
	app.Close()
	logger.Sync()
}
