/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ytt

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type Ytt struct {
	LayerContributor libpak.DependencyLayerContributor
	Logger           bard.Logger
}

func NewYtt(dependency libpak.BuildpackDependency, cache libpak.DependencyCache) (Ytt, libcnb.BOMEntry) {
	contributor, entry := libpak.NewDependencyLayer(dependency, cache, libcnb.LayerTypes{
		Launch: true,
	})
	return Ytt{LayerContributor: contributor}, entry
}

func (w Ytt) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	w.LayerContributor.Logger = w.Logger

	return w.LayerContributor.Contribute(layer, func(artifact *os.File) (libcnb.Layer, error) {
		w.Logger.Bodyf("Copying to %s", layer.Path)

		binDir := filepath.Join(layer.Path, "bin")

		if err := os.MkdirAll(binDir, 0755); err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to mkdir\n%w", err)
		}

		if err := copyFile(artifact, filepath.Join(binDir, "ytt")); err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to copy Ytt\n%w", err)
		}

		return layer, nil
	})
}

func (w Ytt) Name() string {
	return w.LayerContributor.LayerName()
}

func copyFile(sourceFile *os.File, destinationFile string) error {
	input, err := io.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	// needs to be executable
	err = os.WriteFile(destinationFile, input, 0755)
	if err != nil {
		return err
	}

	return nil
}
