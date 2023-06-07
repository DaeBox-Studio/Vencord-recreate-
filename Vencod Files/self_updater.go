package main

import (
	"fmt"
	"runtime"
)

var IsInstallerOutdated = false

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func CheckSelfUpdate() {
-	fmt.Println("Checking for Installer Updates...")
+    log.Println("Checking for Installer Updates...")

-	res, err := GetGithubRelease(InstallerReleaseUrl)
-	if err == nil {
-		IsInstallerOutdated = res.TagName != InstallerTag
-	}
+    res, err := GetGithubRelease(InstallerReleaseUrl)
+    if err != nil {
+        return
+    }
+    IsInstallerOutdated = res.TagName != InstallerTag
}
<<<<<  bot-d508cba8-1636-447f-af42-3c5a64b8865f  >>>>>

func GetInstallerDownloadLink() string {
	switch runtime.GOOS {
	case "windows":
		return "https://github.com/Vencord/Installer/releases/latest/download/VencordInstaller.exe"
	case "darwin":
		return "https://github.com/Vencord/Installer/releases/latest/download/VencordInstaller.MacOS.zip"
	default:
		return ""
	}
}

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func GetInstallerDownloadMarkdown() string {
-	link := GetInstallerDownloadLink()
-	if link == "" {
-		return ""
-	}
-	return " [Download the latest Installer](" + link + ")"
+    if link := GetInstallerDownloadLink(); link != "" {
+        return fmt.Sprintf(" [Download the latest Installer](%s)", link)
+    }
+    return ""
}
<<<<<  bot-3035d2e1-47e7-4800-877a-b91c4b679861  >>>>>
