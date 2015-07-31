/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package siesta

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func setStringConfig(where *string, what string) {
	if what != "" {
		*where = what
	}
}

func setStringSliceConfig(where *[]string, what string, delimiter string) {
	if what != "" {
		splitted := strings.Split(what, delimiter)
		if len(splitted) > 0 {
			*where = splitted
		}
	}
}

func setBoolConfig(where *bool, what string) {
	if what != "" {
		*where = what == "true"
	}
}

func setDurationConfig(where *time.Duration, what string) error {
	if what != "" {
		value, err := time.ParseDuration(what)
		if err == nil {
			*where = value
		}
		return err
	}
	return nil
}

func setIntConfig(where *int, what string) error {
	if what != "" {
		value, err := strconv.Atoi(what)
		if err == nil {
			*where = value
		}
		return err
	}
	return nil
}

func setInt32Config(where *int32, what string) error {
	if what != "" {
		value, err := strconv.Atoi(what)
		if err == nil {
			*where = int32(value)
		}
		return err
	}
	return nil
}

func GetSerializer(name string) (func(value interface{}) ([]byte, error), error) {
	switch name {
	case "":
		return ByteSerializer, nil
	case "ByteSerializer":
		return ByteSerializer, nil
	case "StringSerializer":
		return StringSerializer, nil
	default:
		return nil, fmt.Errorf("Unknown serializer: %s", name)
	}
}
