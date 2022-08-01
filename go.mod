module github.com/xxf098/go-tun2socks-build

go 1.15

require (
	github.com/eycorsican/go-tun2socks v1.16.8
	github.com/xtls/xray-core v1.5.9+incompatible
	golang.org/x/mobile v0.0.0-20220722155234-aaac322e2105
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d
)

// replace with v2ray-core path
replace github.com/xtls/xray-core v1.5.9+incompatible => ../xray-core
