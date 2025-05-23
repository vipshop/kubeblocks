/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package util

import (
	"github.com/StudioSol/set"
)

// Set type Reference c++ set interface to implement stl set.

type Sets = set.LinkedHashSetString

func NewSet(strings ...string) *Sets {
	return set.NewLinkedHashSetString(strings...)
}

func Difference(left, right *Sets) *Sets {
	diff := NewSet()
	for e := range left.Iter() {
		if !right.InArray(e) {
			diff.Add(e)
		}
	}
	return diff
}

func ToSet[T interface{}](v map[string]T) *Sets {
	r := NewSet()
	for k := range v {
		r.Add(k)
	}
	return r
}

func EqSet(left, right *Sets) bool {
	if left.Length() != right.Length() {
		return false
	}
	for e := range left.Iter() {
		if !right.InArray(e) {
			return false
		}
	}
	return true
}

func MapKeyDifference[T interface{}](left, right map[string]T) *Sets {
	lSet := ToSet(left)
	rSet := ToSet(right)
	return Difference(lSet, rSet)
}

func Union(left, right *Sets) *Sets {
	deleteSet := Difference(left, right)
	return Difference(left, deleteSet)
}
