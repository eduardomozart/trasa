//go:build !windows
// +build !windows

package hygiene

/**
 * Copyright (C) 2020 Seknox Pte Ltd.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

type DeviceWindows struct {
}

func (d DeviceWindows) IsAutoLoginEnabled() (bool, error) {
	return false, nil
}

func (d DeviceWindows) IsDeviceEncrypted() (bool, error) {
	return false, nil
}

func (d DeviceWindows) GetInstalledPackages() ([]string, error) {
	return nil, nil
}

// slower
func getOSNameVersionFromSysinfo() (string, string, string, error) {
	return "", "", "", nil
}

func getOSNameVersionFromPS() (string, string, string, error) {
	return "", "", "", nil
}

func (d DeviceWindows) GetOSNameVersion() (string, string, string, error) {
	return "", "", "", nil
}

func (d DeviceWindows) IsRemoteConnectionEnabled() (bool, error) {
	return false, nil
}

func (d DeviceWindows) EndpointSecurity() (string, string, bool, error) {
	return "", "", false, nil
}

func (d DeviceWindows) ScreenLockEnabled() (bool, error) {
	return false, nil
}

func (d DeviceWindows) GetPasswordLastUpdated() (string, error) {
	return "", nil

}
func (d DeviceWindows) GetCriticalAutoUpdateStatus() (bool, error) {
	return false, nil

}
func (d DeviceWindows) GetPendingUpdates() ([]string, error) {
	return nil, nil
}

func (d DeviceWindows) IsFireWallSet() (bool, error) {
	return false, nil
}

func (d DeviceWindows) GetNetwork() (string, string, string, error) {
	return "", "", "", nil
}

func (d DeviceWindows) GetLatestSecurityPatch() (string, error) {
	return "", nil
}

func (d DeviceWindows) IdleDeviceScreenLockTime() (string, error) {
	return "", nil
}
