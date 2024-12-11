package enum

type Provider string

var (
	ProviderStripe Provider = "stripe"
	ProviderLink   Provider = "link"
)

func (p Provider) String() string { return string(p) }
