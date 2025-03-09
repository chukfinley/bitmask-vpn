//go:build linux
// +build linux

// Copyright (C) 2018, 2020 LEAP
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package helper

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"

	"0xacab.org/leap/bitmask-vpn/pkg/config"
)

const (
	openvpnUser       = "nobody"
	openvpnGroup      = "nogroup"
	LogFolder         = "/var/log/"
	systemOpenvpnPath = "/usr/sbin/openvpn"
)

var (
	snapOpenvpnPath = "/snap/bin/" + config.ProviderConfig.BinaryName + ".openvpn"
)

func getPlatformOpenvpnFlags() []string {
	return []string{
		"--script-security", "1",
		"--user", openvpnUser,
		"--group", openvpnGroup,
	}
}

func parseCliArgs() {
	// linux helper does not reply to args
}

func initializeService(port int) {}

func daemonize() {}

func getOpenvpnPath() string {
	if os.Getenv("SNAP") != "" {
		return snapOpenvpnPath
	}
	return systemOpenvpnPath
}

func kill(cmd *exec.Cmd) error {
	return cmd.Process.Signal(os.Interrupt)
}

func firewallStart(gateways []string, mode string) error {
	log.Warn().Msg("Start firewall: do nothing, not implemented")
	return nil
}

func firewallStop() error {
	log.Warn().Msg("Stop firewall: do nothing, not implemented")
	return nil
}

func firewallIsUp() bool {
	log.Warn().Msg("IsUp firewall: do nothing, not implemented")
	return false
}
