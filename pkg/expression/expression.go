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

package expression

func TrinaryOperation(expression bool, expect, standBy any) any {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationString(expression bool, expect, standBy string) string {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationInt64(expression bool, expect, standBy int64) int64 {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationInt32(expression bool, expect, standBy int32) int32 {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationInt(expression bool, expect, standBy int) int {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationUInt(expression bool, expect, standBy uint) uint {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationUInt32(expression bool, expect, standBy uint32) uint32 {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationUInt64(expression bool, expect, standBy uint64) uint64 {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationFloat64(expression bool, expect, standBy float64) float64 {
	if expression {
		return expect
	}

	return standBy
}

func TrinaryOperationFloat32(expression bool, expect, standBy float32) float32 {
	if expression {
		return expect
	}

	return standBy
}

func GenericTrinaryOperation[T any](expression bool, expect, standBy T) T {
	if expression {
		return expect
	}

	return standBy
}
