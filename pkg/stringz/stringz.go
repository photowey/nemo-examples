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

package stringz

import (
	"fmt"
)

func String(source any) string {
	return fmt.Sprintf("%v", source)
}

func ReplaceTemplate(template string, args ...any) string {
	return fmt.Sprintf(template, args...)
}

func IsBlankString(str string) bool {
	return "" == str
}

func IsNotBlankString(str string) bool {
	return !IsBlankString(str)
}

func IsEmptyStringSlice(target []string) bool {
	return len(target) == 0
}

func IsNotEmptyStringSlice(target []string) bool {
	return !IsEmptyStringSlice(target)
}
