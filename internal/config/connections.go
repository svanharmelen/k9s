// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

// Connection
type Connection struct {
	Commands []struct {
		Command     string `json:"command" yaml:"command"`
		WaitForPort int    `json:"waitForPort" yaml:"waitForPort"`
	} `json:"commands,omitempty" yaml:"commands,omitempty"`
	KubeConfig string `json:"kubeConfig" yaml:"kubeConfig"`
}
