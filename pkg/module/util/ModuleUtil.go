/*
 * Copyright (c) 2020-2024. Devtron Inc.
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

package moduleUtil

import (
	"fmt"
	"strings"
)

func BuildAllModuleEnableKeys(moduleName string) []string {
	var keys []string
	keys = append(keys, BuildModuleEnableKey(moduleName))
	if strings.Contains(moduleName, ".") {
		parent := strings.Split(moduleName, ".")[0]
		keys = append(keys, BuildModuleEnableKey(parent))
	}
	return keys
}

func BuildModuleEnableKey(moduleName string) string {
	return fmt.Sprintf("%s.%s", moduleName, "enabled")
}
