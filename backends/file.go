/*
 * This file is part of remco.
 * © 2016 The Remco Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package backends

import (
	"github.com/HeavyHorst/easyKV/file"
	berr "github.com/HeavyHorst/remco/backends/error"
	"github.com/HeavyHorst/remco/log"
	"github.com/HeavyHorst/remco/template"
	"github.com/Sirupsen/logrus"
)

// FileConfig represents the config for the file backend.
type FileConfig struct {
	Filepath string
	template.Backend
}

// Connect creates a new fileClient and fills the underlying template.Backend with the file-Backend specific data.
func (c *FileConfig) Connect() (template.Backend, error) {
	if c == nil {
		return template.Backend{}, berr.ErrNilConfig
	}

	c.Backend.Name = "file"
	log.WithFields(logrus.Fields{
		"backend":  c.Backend.Name,
		"filepath": c.Filepath,
	}).Info("Set filepath")

	client, err := file.New(c.Filepath)
	if err != nil {
		return c.Backend, err
	}
	c.Backend.ReadWatcher = client
	return c.Backend, nil
}
