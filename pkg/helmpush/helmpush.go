/*
(c) Copyright 2018, Gemalto. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helmpush

import (
	"github.com/gemalto/gokube/pkg/download"
	"github.com/gemalto/gokube/pkg/gokube"
	"github.com/gemalto/gokube/pkg/utils"
	"os"
)

const (
	URL = "https://github.com/chartmuseum/helm-push/releases/download/v%s/helm-push_%s_windows_amd64.tar.gz"
)

// InstallPlugin ...
func InstallPlugin(helmPushVersion string) {
	if helmPushVersion[0] == 118 {
		helmPushVersion = helmPushVersion[1:]
	}
	var pluginHome = utils.GetUserHome() + "/.helm/plugins/helm-push"
	utils.CreateDir(pluginHome)
	utils.CreateDir(pluginHome + "/bin")
	if _, err := os.Stat(pluginHome + "/bin/helmpush.exe"); os.IsNotExist(err) {
		download.DownloadFromUrl("helm-push v"+helmPushVersion, URL, helmPushVersion)
		utils.MoveFile(gokube.GetTempDir()+"/bin/helmpush.exe", pluginHome+"/bin/helmpush.exe")
		utils.MoveFile(gokube.GetTempDir()+"/plugin.yaml", pluginHome+"/plugin.yaml")
		utils.RemoveDir(gokube.GetTempDir())
	}
}

// DeletePlugin ...
func DeletePlugin() {
	var pluginHome = utils.GetUserHome() + "/.helm/plugins/helm-push"
	_, err := os.Stat(pluginHome)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}
	utils.RemoveFiles(pluginHome + "/*")
	utils.RemoveDir(pluginHome)
}
