//go:build !ios && !android
// +build !ios,!android

package v2ray

import (
	_ "github.com/xtls/xray-core/app/commander"
	_ "github.com/xtls/xray-core/app/log/command"
	_ "github.com/xtls/xray-core/app/proxyman/command"
	_ "github.com/xtls/xray-core/app/stats/command"

	_ "github.com/xtls/xray-core/app/reverse"

	_ "github.com/xtls/xray-core/transport/internet/domainsocket"
)
