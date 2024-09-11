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

package engine

import (
	"errors"
	"os"
	"strings"

	"github.com/4ra1n/poc-runner/xerr"
)

func getTargetList() ([]string, error) {
	var targetList []string
	// 如果结尾是 txt 且开头没有 http 认为输入是文件
	if strings.HasSuffix(target, ".txt") &&
		!strings.HasPrefix(target, "http") {
		data, err := os.ReadFile(target)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		splits := strings.Split(string(data), "\n")
		if len(splits) < 1 {
			return nil, xerr.Wrap(errors.New("invalid target.txt file"))
		}
		for _, split := range splits {
			split = strings.TrimSpace(split)
			split = strings.Trim(split, "\r")
			targetList = append(targetList, split)
		}
	} else {
		targetList = append(targetList, target)
	}
	return targetList, nil
}
