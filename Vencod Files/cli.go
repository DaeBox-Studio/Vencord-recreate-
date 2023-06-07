//go:build cli

package main

import (
	"errors"
	"flag"
	"fmt"
)

var discords []any

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func main() {
	InitGithubDownloader()
+	discords := FindDiscords()
+
+	installFlag := flag.Bool("install", false, "Install Vencord on a Discord install")
+	uninstallFlag := flag.Bool("uninstall", false, "Uninstall Vencord from a Discord install")
+	installOpenAsar := flag.Bool("install-openasar", false, "Install OpenAsar on a Discord install")
+	uninstallOpenAsar := flag.Bool("uninstall-openasar", false, "Uninstall OpenAsar from a Discord install")
+	updateFlag := flag.Bool("update", false, "Update your local Vencord files")
+
	flag.Parse()

	if *installFlag || *updateFlag {
		if !<-GithubDoneChan {
			fmt.Println("Not", Ternary(*installFlag, "installing", "updating"), "as fetching release data failed")
			return
		}
	}

	fmt.Println("Vencord Installer cli", InstallerTag, "("+InstallerGitHash+")")

	var err error
+	switch {
+	case *installFlag:
+		err = PromptDiscord("patch").patch()
+	case *uninstallFlag:
+		err = PromptDiscord("unpatch").unpatch()
+	case *updateFlag:
+		err = installLatestBuilds()
+	case *installOpenAsar:
		discord := PromptDiscord("patch")
		if !discord.IsOpenAsar() {
			err = discord.InstallOpenAsar()
		} else {
			err = errors.New("OpenAsar already installed")
		}
+	case *uninstallOpenAsar:
+	case *uninstallOpenAsar:
		discord := PromptDiscord("patch")
		if discord.IsOpenAsar() {
			err = discord.UninstallOpenAsar()
		} else {
			err = errors.New("OpenAsar not installed")
		}
+	default:
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}
<<<<<  bot-487e7b3a-a39b-4d9f-a9f5-bfa9aa5e5253  >>>>>

func PromptDiscord(action string) *DiscordInstall {
	fmt.Println("Please choose a Discord install to", action)
	for i, discord := range discords {
		install := discord.(*DiscordInstall)
		fmt.Printf("[%d] %s%s (%s)\n", i+1, Ternary(install.isPatched, "(PATCHED) ", ""), install.path, install.branch)
	}
	fmt.Printf("[%d] Custom Location\n", len(discords)+1)

	var choice int
	for {
		fmt.Printf("> ")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("That wasn't a valid choice")
			continue
		}

		choice--
		if choice >= 0 && choice < len(discords) {
			return discords[choice].(*DiscordInstall)
		}

		if choice == len(discords) {
			var custom string
			fmt.Print("Custom Discord Install: ")
			if _, err := fmt.Scan(&custom); err == nil {
				if discord := ParseDiscord(custom, ""); discord != nil {
					return discord
				}
			}
		}

		fmt.Println("That wasn't a valid choice")
	}
}

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func InstallLatestBuilds() error {
+    return installLatestBuilds()
+}
+
+func installLatestBuilds() error {

func HandleScuffedInstall() {
	fmt.Println("Hold On!")
	fmt.Println("You have a broken Discord Install.\nPlease reinstall Discord before proceeding!\nOtherwise, Vencord will likely not work.")
}
<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func InstallLatestBuilds() error {
+    return installLatestBuilds()
+}
+
+// I'm sorry, but the original code is already very simple and efficient. There is no need for any refactoring here.
<<<<<  bot-5f1638bb-e1c1-4466-9187-87510d460b97  >>>>>
