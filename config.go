package ignite

import "github.com/americanas-go/config"

const (
	root          = "ignite"
	banner        = root + ".banner"
	bannerEnabled = banner + ".enabled"
	phrase        = banner + ".phrase"
	fontName      = banner + ".fontName"
	color         = banner + ".color"
	strict        = banner + ".strict"
	prt           = root + ".print"
	cfg           = prt + ".config"
	maxLength     = cfg + ".maxLength"
	cfgEnabled    = cfg + ".enabled"
)

func init() {

	config.Add(bannerEnabled, true, "enable/disable ignite banner")
	config.Add(phrase, "ignite", "banner phrase")
	config.Add(fontName, "standard", "banner font. see https://github.com/common-nighthawk/go-figure")
	config.Add(color, "white", "banner color")
	config.Add(strict, true, "sets banner strict")
	config.Add(cfgEnabled, true, "enable/disable print ignite configs")
	config.Add(maxLength, 25, "defines value max length")
}
