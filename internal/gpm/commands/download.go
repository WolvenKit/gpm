/*
 Copyright (c) 2020 - 2021 the WolvenKit contributors.

 Licensed under the GNU Affero General Public License v3.0 (the "License").
*/

package commands

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadCommand() *cli.Command {
	command := cli.Command{
		Name:     "Download",
		Aliases:  []string{"d", "--download", "-d"},
		Usage:    "Download the specified mod",
		Category: "Download",
		Action: func(context *cli.Context) error {
			//DownloadMod()
			return nil
		},
	}

	return &command
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
