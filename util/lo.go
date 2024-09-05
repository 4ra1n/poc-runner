/*
 * poc-runner project
 * Copyright (C) 2024 4ra1n
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package util

func Keys[K comparable, V any](in ...map[K]V) []K {
	size := 0
	for i := range in {
		size += len(in[i])
	}
	result := make([]K, 0, size)
	for i := range in {
		for k := range in[i] {
			result = append(result, k)
		}
	}
	return result
}

func Contains[T comparable](collection []T, element T) bool {
	for i := range collection {
		if collection[i] == element {
			return true
		}
	}
	return false
}

func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i := range collection {
		result[i] = iteratee(collection[i], i)
	}
	return result
}
