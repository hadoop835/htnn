// Copyright The HTNN Authors.
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

//go:build mcp_output

package helper

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	controlleroutput "mosn.io/htnn/controller/internal/controller/output"
	"mosn.io/htnn/controller/pkg/procession"
)

type OutputSuite struct {
}

func (o *OutputSuite) Name() string {
	return "mcp"
}

func (o *OutputSuite) Get(ctx context.Context, c client.Client) procession.Output {
	output, err := controlleroutput.NewMcpOutput(ctx)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		defer GinkgoRecover()

		mc := NewMcpClient(c)
		defer mc.Close()
		mc.Init()
		mc.Handle()
	}()
	return output
}
