package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	path "path/filepath"
	"strconv"
)

const OpenAsarDownloadLink = "https://github.com/GooseMod/OpenAsar/releases/download/nightly/app.asar"

<<<<<<<<<<<<<  ✨ Codeium AI Suggestion  >>>>>>>>>>>>>>
func FindAsarFile(dir string) (*os.File, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
+    files := []string{"app.asar", "_app.asar"}
+    for _, file := range files {
+        f, err := os.Open(filepath.Join(dir, file))
+        if err != nil {
+            continue
+        }
+        defer f.Close()
+        stats, err := f.Stat()
+        if err == nil && !stats.IsDir() {
+            return f, nil
+        }
+    }
+    return nil, fmt.Errorf("Install at %s has no asar file", dir)
}
<<<<<  bot-faf6cdf9-49c7-4c80-a047-57e0db47d617  >>>>>

func (di *DiscordInstall) IsOpenAsar() (retBool bool) {
	if di.isOpenAsar != nil {
		return *di.isOpenAsar
	}

	defer func() {
		fmt.Println("Checking if", di.path, "is using OpenAsar:", retBool)
		di.isOpenAsar = &retBool
	}()

	for _, version := range di.versions {
		fmt.Println(version, path.Join(version, ".."))
		asarFile, err := FindAsarFile(path.Join(version, ".."))
		if err != nil {
			fmt.Println(err)
			continue
		}

		b, err := io.ReadAll(asarFile)
		_ = asarFile.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if bytes.Contains(b, []byte("OpenAsar")) {
			return true
		}
	}

	return false
}

func (di *DiscordInstall) InstallOpenAsar() error {
	for _, version := range di.versions {
		dir := path.Join(version, "..")
		asarFile, err := FindAsarFile(dir)
		if err != nil {
			return err
		}
		_ = asarFile.Close()

		if err = os.Rename(asarFile.Name(), path.Join(dir, "app.asar.original")); err != nil {
			return err
		}

		res, err := http.Get(OpenAsarDownloadLink)
		if err != nil {
			return err
		} else if res.StatusCode >= 300 {
			return errors.New("Failed to fetch OpenAsar - " + strconv.Itoa(res.StatusCode) + ": " + res.Status)
		}

		outFile, err := os.Create(asarFile.Name())
		if err != nil {
			return err
		}

		if _, err = io.Copy(outFile, res.Body); err != nil {
			return err
		}
	}

	di.isOpenAsar = Ptr(true)
	return nil
}

func (di *DiscordInstall) UninstallOpenAsar() error {
	for _, version := range di.versions {
		dir := path.Join(version, "..")
		originalAsar := path.Join(dir, "app.asar.original")
		if !ExistsFile(originalAsar) {
			return errors.New("No app.asar.original. Reinstall Discord")
		}

		asarFile, err := FindAsarFile(dir)
		if err != nil {
			return err
		}
		_ = asarFile.Close()

		if err = os.Rename(originalAsar, asarFile.Name()); err != nil {
			return err
		}
	}

	di.isOpenAsar = Ptr(false)
	return nil
}
