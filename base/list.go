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
	"fmt"
	"sync"

	"github.com/4ra1n/poc-runner/xerr"
)

// List
// 这是一个 list 的包装
type List[T any] struct {
	lock  sync.Mutex
	items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		lock:  sync.Mutex{},
		items: []T{},
	}
}

func (l *List[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = []T{}
}

func (l *List[T]) Add(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) (T, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var zeroValue T
	if index < 0 || index >= len(l.items) {
		return zeroValue, xerr.Wrap(fmt.Errorf("index out of range"))
	}
	return l.items[index], nil
}

func (l *List[T]) Remove(index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index >= len(l.items) {
		return xerr.Wrap(fmt.Errorf("index out of range"))
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil
}

func (l *List[T]) Length() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return len(l.items)
}

func (l *List[T]) Items() []T {
	l.lock.Lock()
	defer l.lock.Unlock()
	copiedItems := make([]T, len(l.items))
	copy(copiedItems, l.items)
	return copiedItems
}
