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

package filez

import (
	"bufio"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func Override(path string, contents []byte) error {
	return Write(path, contents, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o666)
}

func Append(path string, contents []byte) error {
	return Write(path, contents, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o766)
}

func Write(path string, contents []byte, flag int, perm os.FileMode) error {
	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	writer := bufio.NewWriter(file)
	_, err = writer.Write(contents)
	if err != nil {
		return err
	}

	err = writer.Flush()

	return err
}
