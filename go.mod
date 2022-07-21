module github.com/xxf098/go-tun2socks-build

go 1.15

require (
	github.com/eycorsican/go-tun2socks v1.16.8
	github.com/v2fly/v2ray-core/v4 v4.45.2+incompatible
	golang.org/x/mobile v0.0.0-20191210151939-1a1fef82734d
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f
)

// replace with v2ray-core path
replace github.com/v2fly/v2ray-core/v4 v4.45.2+incompatible => ../v2ray-core
