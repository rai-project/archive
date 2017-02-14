package archive

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"

	"github.com/Sirupsen/logrus"
	"github.com/Unknwon/com"
	"github.com/mholt/archiver"

	"github.com/rai-project/utils"
)

const (
	FileExtension = ".tar.bz2"
	MimeType      = "application/gzip"
)

var (
	log                *logrus.Entry
	directorySizeLimit = int64(200 * com.MByte)
	format             = archiver.TarBz2
)

func Zip(targetFile string, inputDir string) (string, error) {
	if !com.IsDir(inputDir) {
		msg := "Directory " + inputDir + " not found"
		log.Error(msg)
		return "", errors.New(msg)
	}
	dirSize, err := utils.DirSize(inputDir)
	if err != nil {
		msg := "Cannot get size of inputDirectory."
		log.WithField("directory", inputDir).Error(msg)
		return "", errors.New(msg)
	}
	if dirSize > directorySizeLimit {
		msg := fmt.Sprintf(
			"Directory size limit exceeded (%v). The directory must be %v bytes or less.",
			dirSize,
			directorySizeLimit,
		)
		log.WithField("directory", inputDir).
			WithField("directory_size", dirSize).
			WithField("limit", directorySizeLimit).
			Error(msg)
		return "", errors.New(msg)
	}
	allfiles, err := com.GetFileListBySuffix(inputDir, "")
	pp.Println(allfiles)
	if err != nil {
		log.WithError(err).Error("Failed to get directory " + inputDir + " contents.")
		return "", err
	}
	files := []string{}
	for _, file := range allfiles {
		if strings.Contains(file, ".git") {
			continue
		}
		files = append(files, file)
	}
	err = format.Make(targetFile, files)
	if err != nil {
		log.WithError(err).Error("Failed to create archive for " + inputDir + ".")
		return "", err
	}
	if err != nil {
		log.WithError(err).Error("Failed to create archive for " + inputDir + ".")
		return "", err
	}
	return targetFile, nil
}

func Unzip(targetDir string, fileName string) (string, error) {
	err := format.Open(fileName, targetDir)
	if err != nil {
		log.WithError(err).Error("Failed to open archive " + fileName + ".")
		return "", err
	}
	return targetDir, nil
}
