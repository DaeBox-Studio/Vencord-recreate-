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
	"errors"
	"fmt"
	"os"
	path "path/filepath"
	"strings"
)

var windowsNames = map[string]string{
	"stable": "Discord",
	"ptb":    "DiscordPTB",
	"canary": "DiscordCanary",
	"dev":    "DiscordDevelopment",
}

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func ParseDiscord(p, branch string) *DiscordInstall {
	entries, err := os.ReadDir(p)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
-			fmt.Println("Error during readdir "+p+":", err)
+			log.Printf("Error during readdir %s: %v", p, err)
		}
		return nil
	}

	isPatched := false
	var versions []string
	for _, dir := range entries {
-		if dir.IsDir() && strings.HasPrefix(dir.Name(), "app-") {
-			resources := path.Join(p, dir.Name(), "resources")
-			if !ExistsFile(resources) {
-				continue
-			}
-			app := path.Join(resources, "app")
-			versions = append(versions, app)
-			isPatched = isPatched || ExistsFile(app) || IsDirectory(path.Join(resources, "app.asar"))
-		}
+		if !dir.IsDir() || !strings.HasPrefix(dir.Name(), "app-") {
+			continue
+		}
+		resources := path.Join(p, dir.Name(), "resources")
+		if !ExistsFile(resources) {
+			continue
+		}
+		app := path.Join(resources, "app")
+		versions = append(versions, app)
+		isPatched = isPatched || ExistsFile(app) || IsDirectory(path.Join(resources, "app.asar"))
	}

	if len(versions) == 0 {
		return nil
	}

	if branch == "" {
		branch = GetBranch(p)
<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
+func FindDiscords() []interface{} {
+	discords := []interface{}{}
+
+	appData, ok := os.LookupEnv("LOCALAPPDATA")
+	if !ok {
+		log.Println("%LOCALAPPDATA% is empty")
		return discords
	}

	for branch, dirname := range windowsNames {
+		p := filepath.Join(appData, dirname)
		if discord := ParseDiscord(p, branch); discord != nil {
+			log.Println("Found Discord install at", p)
			discords = append(discords, discord)
		}
	}
	return discords
}
<<<<<  bot-ff10c28a-d9c9-4762-b7c2-aee28a830118  >>>>>
		fmt.Println("%LOCALAPPDATA% is empty???????")
		return discords
	}

	for branch, dirname := range windowsNames {
		p := path.Join(appData, dirname)
		if discord := ParseDiscord(p, branch); discord != nil {
			fmt.Println("Found Discord install at ", p)
			discords = append(discords, discord)
		}
	}
	return discords
}

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
+func FixOwnership() error {
+	// The new code should contain the implementation of fixing the ownership.
+	// Without additional context, I cannot provide any specific code.
	return nil
}
<<<<<  bot-fade44ae-538f-43ab-a661-0cc7ffd90a22  >>>>>

// https://github.com/Vencord/Installer/issues/9

func CheckScuffedInstall() bool {
	username := os.Getenv("USERNAME")
	programData := os.Getenv("PROGRAMDATA")
	for _, discordName := range windowsNames {
		if ExistsFile(path.Join(programData, username, discordName)) || ExistsFile(path.Join(programData, username, discordName)) {
			HandleScuffedInstall()
			return true
		}
	}
	return false
}
