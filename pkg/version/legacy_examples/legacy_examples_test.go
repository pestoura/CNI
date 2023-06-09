// Copyright 2016 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package legacy_examples_test

import (
	"os"
	"path/filepath"
	"runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/containernetworking/cni/pkg/version/legacy_examples"
)

var _ = Describe("The v0.1.0 Example", func() {
	It("builds ok", func() {
		example := legacy_examples.V010
		pluginPath, err := example.Build()
		Expect(err).NotTo(HaveOccurred())

		expectedBaseName := example.Name
		if runtime.GOOS == "windows" {
			expectedBaseName += ".exe"
		}
		Expect(filepath.Base(pluginPath)).To(Equal(expectedBaseName))

		Expect(os.RemoveAll(pluginPath)).To(Succeed())
	})
})
