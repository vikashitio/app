package handlers

import (
	"fmt"
	"template/database"

	"github.com/gofiber/fiber/v2"
)

type TxResponse struct {
	Txhash    string `json:"txhash"`
	Height    string `json:"height"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Timestamp string `json:"timestamp"`
}

// Handler to get balance by address
func TestAPIS(c *fiber.Ctx) error {
	// PostgreSQL connection string

	// Prepare the insert statement
	insertQuery := `
        INSERT INTO checkcryptoaddress (name, symbol)
        VALUES ($1, $2)`

	// Sample data to insert
	coins := [][]interface{}{{"0x", "zrx"},
		{"Aave Coin", "aave"},
		{"Algorand", "algo"},
		{"Apecoin", "ape"},
		{"API3", "api3"},
		{"Aragon", "ant"},
		{"Arbitrum", "arb"},
		{"Augur", "rep"},
		{"AugurV2", "repv2"},
		{"AuroraCoin", "aur"},
		{"Avalanche", "avax"},
		{"Axie Infinity", "axs"},
		{"Bancor", "bnt"},
		{"Band Protocol", "band"},
		{"Bankex", "bkx"},
		{"Basic Attention Token", "bat"},
		{"BeaverCoin", "bvc"},
		{"Biconomy", "bico"},
		{"Binance", "bnb"},
		{"BioCoin", "bio"},
		{"Bitcoin", "btc"},
		{"BitcoinCash", "bch"},
		{"BitcoinGold", "btg"},
		{"BitcoinPrivate", "btcp"},
		{"Bitcoin SV", "bsv"},
		{"BitcoinZ", "btcz"},
		{"BlockTrade", "btt"},
		{"Blur", "blur"},
		{"Bonk", "bonk"},
		{"BTU Protocol", "btu"},
		{"Callisto", "clo"},
		{"Cardano", "ada"},
		{"Celo", "celo"},
		{"Chainlink", "link"},
		{"Chiliz", "chz"},
		{"Civic", "cvc"},
		{"Compound", "comp"},
		{"Cred", "lba"},
		{"Crypto.com Coin", "cro"},
		{"Curve DAO", "crv"},
		{"CUSD", "cusd"},
		{"Dash", "dash"},
		{"Decentraland", "mana"},
		{"Decred", "dcr"},
		{"DigiByte", "dgb"},
		{"District0x", "dnt"},
		{"DogeCoin", "doge"},
		{"Enjin Coin", "enj"},
		{"EOS", "eos"},
		{"Ethereum", "eth"},
		{"EthereumClassic", "etc"},
		{"Ethereum Name Service", "ens"},
		{"EthereumPow", "ethw"},
		{"EtherZero", "etz"},
		{"Expanse", "exp"},
		{"Fetch.ai", "fet"},
		{"FirmaChain", "fct"},
		{"Flare", "flr"},
		{"FreiCoin", "frc"},
		{"GameCredits", "game"},
		{"GarliCoin", "grlc"},
		{"Gnosis", "gno"},
		{"Golem", "glm"},
		{"Golem (GNT)", "gnt"},
		{"Hashflow", "hft"},
		{"HedgeTrade", "hedg"},
		{"Hush", "hush"},
		{"HyperSpace", "xsc"},
		{"iExec RLC", "rlc"},
		{"Illuvium", "ilv"},
		{"Immutable", "imx"},
		{"Injective", "inj"},
		{"Komodo", "kmd"},
		{"LBRY Credits", "lbc"},
		{"Lido DAO Token", "ldo"},
		{"LiteCoin", "ltc"},
		{"loki", "loki"},
		{"Loom Network", "loom"},
		{"Magic", "magic"},
		{"Maker", "mkr"},
		{"Marlin", "pond"},
		{"Mask Network", "mask"},
		{"Matchpool", "gup"},
		{"Matic", "matic"},
		{"MegaCoin", "mec"},
		{"Melon", "mln"},
		{"Metal", "mtl"},
		{"Monero", "xmr"},
		{"Multi-collateral DAI", "dai"},
		{"NameCoin", "nmc"},
		{"Nano", "nano"},
		{"Nem", "xem"},
		{"Neo", "neo"},
		{"NeoGas", "gas"},
		{"Numeraire", "nmr"},
		{"Ocean Protocol", "ocean"},
		{"Odyssey", "ocn"},
		{"OmiseGO", "omg"},
		{"Onyx Protocol", "xcn"},
		{"Optimism", "op"},
		{"Origin Protocol", "ogn"},
		{"Paxos", "pax"},
		{"PayPal USD", "pyusd"},
		{"PeerCoin", "ppc"},
		{"PIVX", "pivx"},
		{"Polkadot", "dot"},
		{"Polymath", "poly"},
		{"PrimeCoin", "xpm"},
		{"ProtoShares", "pts"},
		{"Qtum", "qtum"},
		{"Quant", "qnt"},
		{"Quantum Resistant Ledger", "qrl"},
		{"RaiBlocks", "xrb"},
		{"Ripio Credit Network", "rcn"},
		{"Ripple", "xrp"},
		{"Salt", "salt"},
		{"Serve", "serv"},
		{"Siacoin", "sc"},
		{"Skale", "skl"},
		{"SnowGem", "sng"},
		{"Solana", "sol"},
		{"SolarCoin", "slr"},
		{"SOLVE", "solve"},
		{"Spendcoin", "spnd"},
		{"Status", "snt"},
		{"Stellar", "xlm"},
		{"Storj", "storj"},
		{"Storm", "storm"},
		{"StormX", "stmx"},
		{"SuperVerse", "super"},
		{"Swarm City", "swt"},
		{"Synthetix Network", "snx"},
		{"Tap", "xtp"},
		{"Tellor", "trb"},
		{"TEMCO", "temco"},
		{"TenX", "pay"},
		{"Tether", "usdt"},
		{"Tezos", "xtz"},
		{"The Graph", "grt"},
		{"The Sandbox", "sand"},
		{"Tron", "trx"},
		{"TrueUSD", "tusd"},
		{"Unifi Protocol DAO", "unfi"},
		{"Uniswap Coin", "uni"},
		{"USD Coin", "usdc"},
		{"VeChain", "vet"},
		{"VertCoin", "vtc"},
		{"Viberate", "vib"},
		{"VoteCoin", "vot"},
		{"Vulcan Forged PYR", "pyr"},
		{"Waves", "waves"},
		{"Wings", "wings"},
		{"Yearn.finance", "yfi"},
		{"ZCash", "zec"},
		{"ZClassic", "zcl"},
		{"ZenCash", "zen"}}

	// Insert data
	for _, coin := range coins {
		err := database.DB.Db.Exec(insertQuery, coin[0], coin[1])

		fmt.Println(err)

	}

	fmt.Println("Data inserted successfully!")
	return nil

}
