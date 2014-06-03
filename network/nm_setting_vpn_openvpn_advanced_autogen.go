// This file is automatically generated, please don't edit manully.
package network

import (
	"fmt"
)

// Get key type
func getSettingVpnOpenvpnAdvancedKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_OPENVPN_KEY_PORT:
		t = ktypeUint32
	case NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS:
		t = ktypeUint32
	case NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU:
		t = ktypeUint32
	case NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE:
		t = ktypeUint32
	case NM_SETTING_VPN_OPENVPN_KEY_MSSFIX:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM:
		t = ktypeBoolean
	}
	return
}

// Check is key in current setting section
func isKeyInSettingVpnOpenvpnAdvanced(key string) bool {
	switch key {
	case NM_SETTING_VPN_OPENVPN_KEY_PORT:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_MSSFIX:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnOpenvpnAdvancedDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PORT:
		value = 1194
	case NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS:
		value = 0
	case NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO:
		value = false
	case NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP:
		value = false
	case NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV:
		value = false
	case NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU:
		value = 1500
	case NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE:
		value = 1300
	case NM_SETTING_VPN_OPENVPN_KEY_MSSFIX:
		value = false
	case NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM:
		value = false
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnOpenvpnAdvancedKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnOpenvpnAdvancedKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PORT:
		value = getSettingVpnOpenvpnKeyPortJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS:
		value = getSettingVpnOpenvpnKeyRenegSecondsJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO:
		value = getSettingVpnOpenvpnKeyCompLzoJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP:
		value = getSettingVpnOpenvpnKeyProtoTcpJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV:
		value = getSettingVpnOpenvpnKeyTapDevJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU:
		value = getSettingVpnOpenvpnKeyTunnelMtuJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE:
		value = getSettingVpnOpenvpnKeyFragmentSizeJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_MSSFIX:
		value = getSettingVpnOpenvpnKeyMssfixJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM:
		value = getSettingVpnOpenvpnKeyRemoteRandomJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnOpenvpnAdvancedKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnOpenvpnAdvancedKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_PORT:
		err = setSettingVpnOpenvpnKeyPortJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS:
		err = setSettingVpnOpenvpnKeyRenegSecondsJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO:
		err = setSettingVpnOpenvpnKeyCompLzoJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP:
		err = setSettingVpnOpenvpnKeyProtoTcpJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV:
		err = setSettingVpnOpenvpnKeyTapDevJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU:
		err = setSettingVpnOpenvpnKeyTunnelMtuJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE:
		err = setSettingVpnOpenvpnKeyFragmentSizeJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_MSSFIX:
		err = setSettingVpnOpenvpnKeyMssfixJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM:
		err = setSettingVpnOpenvpnKeyRemoteRandomJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnOpenvpnKeyPortExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT)
}
func isSettingVpnOpenvpnKeyRenegSecondsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS)
}
func isSettingVpnOpenvpnKeyCompLzoExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO)
}
func isSettingVpnOpenvpnKeyProtoTcpExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP)
}
func isSettingVpnOpenvpnKeyTapDevExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV)
}
func isSettingVpnOpenvpnKeyTunnelMtuExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU)
}
func isSettingVpnOpenvpnKeyFragmentSizeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE)
}
func isSettingVpnOpenvpnKeyMssfixExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX)
}
func isSettingVpnOpenvpnKeyRemoteRandomExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM)
}

// Ensure section and key exists and not empty
func ensureSectionSettingVpnOpenvpnAdvancedExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME))
	}
}
func ensureSettingVpnOpenvpnKeyPortNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyPortExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyRenegSecondsNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyRenegSecondsExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyCompLzoNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyCompLzoExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyProtoTcpNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyProtoTcpExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyTapDevNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyTapDevExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyTunnelMtuNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyTunnelMtuExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyFragmentSizeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyFragmentSizeExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyMssfixNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyMssfixExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyRemoteRandomNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyRemoteRandomExists(data) {
		rememberError(errs, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingVpnOpenvpnKeyPort(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnOpenvpnKeyRenegSeconds(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnOpenvpnKeyCompLzo(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnOpenvpnKeyProtoTcp(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnOpenvpnKeyTapDev(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnOpenvpnKeyTunnelMtu(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnOpenvpnKeyFragmentSize(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnOpenvpnKeyMssfix(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnOpenvpnKeyRemoteRandom(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM)
	value = interfaceToBoolean(ivalue)
	return
}

// Setter
func setSettingVpnOpenvpnKeyPort(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT, value)
}
func setSettingVpnOpenvpnKeyRenegSeconds(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS, value)
}
func setSettingVpnOpenvpnKeyCompLzo(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO, value)
}
func setSettingVpnOpenvpnKeyProtoTcp(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP, value)
}
func setSettingVpnOpenvpnKeyTapDev(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV, value)
}
func setSettingVpnOpenvpnKeyTunnelMtu(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU, value)
}
func setSettingVpnOpenvpnKeyFragmentSize(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE, value)
}
func setSettingVpnOpenvpnKeyMssfix(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX, value)
}
func setSettingVpnOpenvpnKeyRemoteRandom(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM, value)
}

// JSON Getter
func getSettingVpnOpenvpnKeyPortJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_PORT))
	return
}
func getSettingVpnOpenvpnKeyRenegSecondsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS))
	return
}
func getSettingVpnOpenvpnKeyCompLzoJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO))
	return
}
func getSettingVpnOpenvpnKeyProtoTcpJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP))
	return
}
func getSettingVpnOpenvpnKeyTapDevJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV))
	return
}
func getSettingVpnOpenvpnKeyTunnelMtuJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU))
	return
}
func getSettingVpnOpenvpnKeyFragmentSizeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE))
	return
}
func getSettingVpnOpenvpnKeyMssfixJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_MSSFIX))
	return
}
func getSettingVpnOpenvpnKeyRemoteRandomJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM))
	return
}

// JSON Setter
func setSettingVpnOpenvpnKeyPortJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_PORT))
}
func setSettingVpnOpenvpnKeyRenegSecondsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS))
}
func setSettingVpnOpenvpnKeyCompLzoJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO))
}
func setSettingVpnOpenvpnKeyProtoTcpJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP))
}
func setSettingVpnOpenvpnKeyTapDevJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV))
}
func setSettingVpnOpenvpnKeyTunnelMtuJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU))
}
func setSettingVpnOpenvpnKeyFragmentSizeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE))
}
func setSettingVpnOpenvpnKeyMssfixJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_MSSFIX))
}
func setSettingVpnOpenvpnKeyRemoteRandomJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM, valueJSON, getSettingVpnOpenvpnAdvancedKeyType(NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM))
}

// Logic JSON Setter

// Remover
func removeSettingVpnOpenvpnKeyPort(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT)
}
func removeSettingVpnOpenvpnKeyRenegSeconds(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS)
}
func removeSettingVpnOpenvpnKeyCompLzo(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_COMP_LZO)
}
func removeSettingVpnOpenvpnKeyProtoTcp(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PROTO_TCP)
}
func removeSettingVpnOpenvpnKeyTapDev(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TAP_DEV)
}
func removeSettingVpnOpenvpnKeyTunnelMtu(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU)
}
func removeSettingVpnOpenvpnKeyFragmentSize(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE)
}
func removeSettingVpnOpenvpnKeyMssfix(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_MSSFIX)
}
func removeSettingVpnOpenvpnKeyRemoteRandom(data connectionData) {
	removeSettingKey(data, NM_SETTING_ALIAS_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_REMOTE_RANDOM)
}
