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

package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/photowey/nemo-examples/configs"
	"github.com/photowey/nemo-examples/pkg/expression"
	"github.com/photowey/nemo-examples/pkg/filez"
	"github.com/photowey/nemo-examples/pkg/logger"
	"github.com/photowey/nemo-examples/pkg/stringz"
	"github.com/photowey/perrors"
	"github.com/valyala/fasthttp"
)

const (
	PidName = "./.run"
)

type Controller struct {
	Host string
}

func NewController(host string) *Controller {
	return &Controller{
		Host: host,
	}
}

func (ctrl *Controller) Run(cancelAble context.Context, crush chan error) {
	go func() {
		crush <- fasthttp.ListenAndServe(ctrl.Host, ctrl.HandleFastHTTP)
	}()

	go func() {
		if err := health(cancelAble); err != nil {
			logger.Error(err.Error())
		} else {
			if logger.IsInfoEnabled() {
				pid := os.Getpid()
				logger.Infof("nemoexp: the server:[pid:%d] has been deployed on:[%s](HTTP) successfully.", pid, configs.Host())

				writePid(pid)
			}
		}
	}()
}

func writePid(pid int) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, PidName)
	_ = filez.Override(path, []byte(stringz.String(pid)))
}

func (ctrl *Controller) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
}

func health(cancelAble context.Context) error {
	host := configs.Host()
	url := fmt.Sprintf("http://%s/health", host)

	bind := strings.Split(host, ":")[0]
	if bind == "" || bind == "0.0.0.0" {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(host, ":")[1])
	}

	for {
		request, err := http.NewRequestWithContext(cancelAble, http.MethodGet, url, nil)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(request)
		if err == nil && resp.StatusCode == http.StatusOK {
			_ = resp.Body.Close()

			return nil
		}

		select {
		case <-cancelAble.Done():
			timeout := expression.TrinaryOperationUInt32(configs.Timeout() > 0, configs.Timeout(), 15)
			return perrors.Errorf("nemoexp: can't' ping http server within the specified time:[%d s] interval", timeout)
		default:
		}
	}
}
