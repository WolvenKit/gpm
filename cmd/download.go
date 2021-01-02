/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package cmd

import (
    "errors"
	"fmt"
    "github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download the specified mod",
    Run: func(cmd *cobra.Command, args []string) {
        DownloadMod("","","","")
    },
}

// Downloads mod from the Mod Registry
func DownloadMod(url string, downloadDir string, identifier string, fileType string) (error, string) {
	response, err := http.Get(url)
	if err != nil {
		return err, ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code"), ""
	}

	p := filepath.FromSlash(fmt.Sprintf("%s/%s%s", downloadDir, identifier, fileType))
	file, err := os.Create(p)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	// Write the response's bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err, ""
	}

	return nil, p
}
