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

package logger

import (
	"fmt"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

var LevelMap = map[Level]int{
	LevelDebug: 1 << 0,
	LevelInfo:  1 << 1,
	LevelWarn:  1 << 2,
	LevelError: 1 << 3,
}

type Config struct {
	Level Level  `yml:"level"`
	Path  string `yml:"path"`
}

func Debug(args ...any) {
}

func Debugf(template string, args ...any) {
}

func Info(args ...any) {
}

func Infof(template string, args ...any) {
	fmt.Printf(template+"\n", args...)
}

func Warn(args ...any) {
}

func Warnf(template string, args ...any) {
}

func Error(args ...any) {
}

func Errorf(template string, args ...any) {
}

func Sync() {
}

func Init(cfg Config) error {
	return nil
}

func LoggerEnabled() bool {
	return true
}

func LoggerNotEnabled() bool {
	return !LoggerEnabled()
}

func IsDebugEnabled() bool {
	return true
}

func IsInfoEnabled() bool {
	return true
}

func IsWarnEnabled() bool {
	return true
}

func IsErrorEnabled() bool {
	return true
}
