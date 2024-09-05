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

package expression

// EFunction
// 函数
type EFunction interface {
	Call(args []EValue) (EValue, error)
	ToString() EString
}

// EValue
// 基本类型
type EValue interface {
	ToString() EString
}

// EObject
// 对象
type EObject interface {
	EValue
	Get(name string) (EValue, error)
	Keys() []string
}
