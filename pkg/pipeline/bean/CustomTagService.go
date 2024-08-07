/*
 * Copyright (c) 2024. Devtron Inc.
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

package bean

import (
	"fmt"
	"github.com/devtron-labs/devtron/internal/sql/repository"
)

const (
	EntityNull = iota
	EntityTypeCiPipelineId
	EntityTypePreCD
	EntityTypePostCD
)

const (
	ImagePathPattern                                              = "%s/%s:%s" // dockerReg/dockerRepo:Tag
	ImageTagUnavailableMessage                                    = "Desired image tag already exists"
	REGEX_PATTERN_FOR_ENSURING_ONLY_ONE_VARIABLE_BETWEEN_BRACKETS = `\{.{2,}\}`
	REGEX_PATTERN_FOR_CHARACTER_OTHER_THEN_X_OR_x                 = `\{[^xX]|{}\}`
	REGEX_PATTERN_FOR_IMAGE_TAG                                   = `^[a-zA-Z0-9]+[a-zA-Z0-9._-]*$`
)

var (
	ErrImagePathInUse = fmt.Errorf(ImageTagUnavailableMessage)
)

const (
	IMAGE_TAG_VARIABLE_NAME_X = "{X}"
	IMAGE_TAG_VARIABLE_NAME_x = "{x}"
)

type CustomTagArrayResponse map[int]map[string]*repository.CustomTag

func (resp CustomTagArrayResponse) GetCustomTagForEntityKey(entityKey int, entityValue string) *repository.CustomTag {
	if resp == nil {
		return nil
	} else if resp[entityKey] == nil {
		return nil
	} else {
		return resp[entityKey][entityValue]
	}
}
