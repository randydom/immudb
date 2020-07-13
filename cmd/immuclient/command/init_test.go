/*
Copyright 2019-2020 vChain, Inc.

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

package immuclient

import (
	"testing"

	"github.com/codenotary/immudb/cmd/helper"
	"github.com/codenotary/immudb/cmd/immuclient/immuclienttest"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/codenotary/immudb/pkg/server/servertest"
	"github.com/spf13/cobra"
)

func TestInit(t *testing.T) {
	opts := helper.Options{}
	cm := Init(&opts)
	if len(cm.Commands()) != 29 {
		t.Fatal("fail immuclient commands, wrong number of expected commanfs")
	}
}

func TestConnect(t *testing.T) {
	options := server.Options{}.WithAuth(true).WithInMemoryStore(true)
	bs := servertest.NewBufconnServer(options)
	bs.Start()
	cmd := commandline{
		immucl: immuclienttest.NewClient(&immuclienttest.PasswordReader{}, bs.Dialer, nil),
	}
	_ = cmd.connect(&cobra.Command{}, []string{})
	cmd.disconnect(&cobra.Command{}, []string{})
}
