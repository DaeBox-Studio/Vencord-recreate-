/*
 * This part is file of VencordInstaller
 * Copyright (c) 2022 Vendicated
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
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	path "path/filepath"
	"strings"
)

var macosNames = map[string]string{
	"stable": "Discord.app",
	"ptb":    "Discord PTB.app",
	"canary": "Discord Canary.app",
	"dev":    "Discord Development.app",
}

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func ParseDiscord(p, branch string) *DiscordInstall {
+	if !ExistsFile(p) || !ExistsFile(path.Join(p, "/Contents/Resources")) {
		return nil
	}

	if branch == "" {
		branch = GetBranch(strings.TrimSuffix(p, ".app"))
	}

+	resources := path.Join(p, "/Contents/Resources")
	app := path.Join(resources, "app")
	return &DiscordInstall{
		path:             p,
		branch:           branch,
		versions:         []string{app},
		isPatched:        ExistsFile(app) || IsDirectory(path.Join(resources, "app.asar")),
		isFlatpak:        false,
		isSystemElectron: false,
	}
}
<<<<<  bot-214a99a6-233b-4057-8f6e-291af6b49388  >>>>>

func FindDiscords() []any {
	var discords []any
	for branch, dirname := range macosNames {
		p := "/Applications/" + dirname
		if discord := ParseDiscord(p, branch); discord != nil {
			fmt.Println("Found Discord Install at", p)
			discords = append(discords, discord)
		}
	}
	return discords
}

func FixOwnership(_ string) error {
	return nil
}

func CheckScuffedInstall() bool {
	return false
}
