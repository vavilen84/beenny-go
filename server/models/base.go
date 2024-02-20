package models

type Blockchain int

type DropStatus int

type BillingPlan int

const (
	_ Blockchain = iota
	EthereumBlockchain
	PolygonBlockchain
	SolanaBlockchain
	ArbitrumBlockchain
	AvalancheBlockchain
	BinanceBlockchain
	CardanoBlockchain
	CeloBlockchain
	CronosBlockchain
	FantomBlockchain
	FlareBlockchain
	IMXBlockchain
	TronBlockchain
	WaxBlockchain
	ZiliqaBlockchain
	TonBlockchain
	AptosBlockchain
	OtherBlockchain
)

const (
	_ DropStatus = iota
	UnPublishedDropStatus
	InReviewDropStatus
	PublishedDropStatus
)

const (
	_ BillingPlan = iota
	StandardBillingPlan
	FeaturedBillingPlan
)

const (
	AtLeastOneWebsiteOrGroupLinkErrMsg = "Must be set at least one of: WebsiteURL, DiscordURL, TwitterURL"
	BlockchainNameRequiredErrMsg       = "Please, set Blockchain name"
)
