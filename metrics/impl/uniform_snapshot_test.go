/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package impl

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUniformSnapshot(t *testing.T) {
	snapshot := NewUniformSnapshot(nil)
	assert.NotNil(t, snapshot)
	size, err := snapshot.Size()
	assert.Equal(t, 0, size)
	assert.Nil(t, err)

	values := []int64{4, 1, 2, 3, }

	snapshot = NewUniformSnapshot(values)
	copied, err := snapshot.GetValues()
	assert.Equal(t, 4, len(copied))
	assert.Equal(t, int64(1), copied[0])
	assert.Equal(t, int64(2), copied[1])
	assert.Equal(t, int64(3), copied[2])
	assert.Equal(t, int64(4), copied[3])
}

func TestUniformSnapshot_GetValue(t *testing.T) {

	snapshot := NewUniformSnapshot([]int64{5, 1, 2, 3, 4})

	// small quantile
	value, err := snapshot.GetValue(0.0)
	assert.True(t, equals(1, value, 0.1))
	assert.Nil(t, err)

	// big quantile
	value, err = snapshot.GetValue(1.0)
	assert.True(t, equals(5, value, 0.1))
	assert.Nil(t, err)

	// invalid quantile
	value, err = snapshot.GetValue(math.NaN())
	assert.NotNil(t, err)
	value, err = snapshot.GetValue(-0.00001)
	assert.NotNil(t, err)
	value, err = snapshot.GetValue(1.0000001)
	assert.NotNil(t, err)
}

// compare two float numbers
func equals(expected float64, actual float64, delta float64) bool {
	return math.Abs(actual-expected) < delta
}
