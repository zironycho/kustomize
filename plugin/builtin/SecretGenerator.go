// +build plugin

/*
Copyright 2019 The Kubernetes Authors.

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

package main

import (
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
	"sigs.k8s.io/kustomize/pkg/types"
	"sigs.k8s.io/yaml"
)

type plugin struct {
	ldr ifc.Loader
	rf  *resmap.Factory
	types.GeneratorOptions
	types.SecretArgs
}

var KustomizePlugin plugin

func (p *plugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, config []byte) (err error) {
	p.GeneratorOptions = types.GeneratorOptions{}
	p.SecretArgs = types.SecretArgs{}
	err = yaml.Unmarshal(config, p)
	p.ldr = ldr
	p.rf = rf
	return
}

func (p *plugin) Generate() (resmap.ResMap, error) {
	argsList := make([]types.SecretArgs, 1)
	argsList[0] = p.SecretArgs
	return p.rf.NewResMapFromSecretArgs(
		p.ldr, &p.GeneratorOptions, argsList)
}
