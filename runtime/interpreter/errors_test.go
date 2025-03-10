/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package interpreter

import (
	"testing"

	"github.com/onflow/cadence/runtime/common"
	"github.com/stretchr/testify/require"
)

func TestOverwriteError_Error(t *testing.T) {

	require.EqualError(t,
		OverwriteError{
			Address: NewAddressValueFromBytes([]byte{0x1}),
			Path: PathValue{
				Domain:     common.PathDomainStorage,
				Identifier: "test",
			},
		},
		"failed to save object: path /storage/test in account 0x1 already stores an object",
	)
}
