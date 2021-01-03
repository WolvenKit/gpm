/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "go.uber.org/zap"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download the specified mod",
    Run: func(cmd *cobra.Command, args []string) {
        //DownloadMod("","","","")
    },
}

// Downloads mod from the Mod Registry
func DownloadMod(logger *zap.SugaredLogger, url string, downloadDir string, identifier string, fileType string) (error, string) {
	logger.Debugf("Downloading %s%s from %s", identifier, fileType, url)
    response, err := http.Get(url)
	if err != nil {
		return err, ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.Errorf("Received response code %s", response.StatusCode)
	}

	p := filepath.FromSlash(fmt.Sprintf("%s/%s%s", downloadDir, identifier, fileType))
    logger.Debugf("Saving archive to %s", p)
    file, err := os.Create(p)
	if err != nil {
		return err, ""
	}
	defer file.Close()

    _, err = io.Copy(file, response.Body)
	if err != nil {
		return err, ""
	}
    logger.Debugf("Archive saved at %s", file.Name())

	return nil, p
}
