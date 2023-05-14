// Copyright 2020 The go-laniakea Authors
// This file is part of the go-laniakea library.
//
// The go-laniakea library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-laniakea library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-laniakea library. If not, see <http://www.gnu.org/licenses/>.

//go:build ios || js
// +build ios js

package metrics

// ReadCPUStats retrieves the current CPU stats. Internally this uses `gosigar`,
// which is not supported on the platforms in this file.
func ReadCPUStats(stats *CPUStats) {}
