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

package base

import (
	"sort"
	"sync"
)

// Map
// golang 的 map 是无序的
// 有时候其实是需要一个有序的 map 结构
type Map[K comparable, V any] struct {
	lock   sync.Mutex
	keys   []K
	values map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		values: make(map[K]V),
	}
}

func (m *Map[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.keys = make([]K, 0)
	m.values = make(map[K]V)
}

func (m *Map[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *Map[K, V]) Get(key K) V {
	m.lock.Lock()
	defer m.lock.Unlock()
	// 自行确保这里不是空
	val, _ := m.values[key]
	return val
}

func (m *Map[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.values[key]; ok {
		delete(m.values, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

func (m *Map[K, V]) Keys() []K {
	m.lock.Lock()
	defer m.lock.Unlock()
	return append([]K(nil), m.keys...)
}

func (m *Map[K, V]) Length() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.keys)
}

func (m *Map[K, V]) Sort(compare func(a, b K) bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	sort.Slice(m.keys, func(i, j int) bool {
		return compare(m.keys[i], m.keys[j])
	})
}

func (m *Map[K, V]) ToMap() map[K]V {
	result := make(map[K]V)
	for _, key := range m.Keys() {
		v := m.Get(key)
		result[key] = v
	}
	return result
}
