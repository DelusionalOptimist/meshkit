package coder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/layer5io/meshkit/cmd/errorutil/internal/component"

	mesherr "github.com/layer5io/meshkit/cmd/errorutil/internal/error"
	"github.com/sirupsen/logrus"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func walkAnalyze(rootDir string, skipDirs []string, errorsInfo *mesherr.InfoAll) error {
	return walk(rootDir, skipDirs, false, false, errorsInfo)
}

func walkUpdate(rootDir string, skipDirs []string, updateAll bool, errorsInfo *mesherr.InfoAll) error {
	return walk(rootDir, skipDirs, true, updateAll, errorsInfo)
}

func walk(rootDir string, skipDirs []string, update bool, updateAll bool, errorsInfo *mesherr.InfoAll) error {
	subDirsToSkip := append([]string{".git", ".github"}, skipDirs...)
	logrus.Info(fmt.Sprintf("root directory: %s", rootDir))
	logrus.Info(fmt.Sprintf("subdirs to skip: %v", subDirsToSkip))
	comp, err := component.New(rootDir)
	if err != nil {
		return err
	}

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		logger := logrus.WithFields(logrus.Fields{"path": path})
		if err != nil {
			logger.WithFields(logrus.Fields{"error": fmt.Sprintf("%v", err)}).Warn("failure accessing path")
			return err
		}
		if info.IsDir() && contains(subDirsToSkip, info.Name()) {
			logger.Infof("skipping directory %s", info.Name())
			return filepath.SkipDir
		}
		if info.IsDir() {
			logger.Debug("handling dir")
		} else {
			if includeFile(path) {
				isErrorsGoFile := isErrorGoFile(path)
				logger.WithFields(logrus.Fields{"iserrorsfile": fmt.Sprintf("%v", isErrorsGoFile)}).Debug("handling Go file")
				err := handleFile(path, update, updateAll, errorsInfo, comp)
				if err != nil {
					return err
				}
			} else {
				logger.Debug("skipping file")
			}
		}
		return nil
	})
	if update {
		err = comp.Write()
	}
	return err
}

func isErrorGoFile(path string) bool {
	_, file := filepath.Split(path)
	return file == "error.go"
}

func includeFile(path string) bool {
	if strings.HasSuffix(path, "_test.go") {
		return false
	}
	if filepath.Ext(path) == ".go" {
		return true
	}
	return false
}
