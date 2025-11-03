package coin

import (
	"errors"
	"fmt"
	"strconv"
)

const BlockchainEthereum = "Ethereum"

func GetCoinForId(id string) (Coin, error) {
	for _, c := range Coins {
		if c.Handle == id {
			return c, nil
		}
	}
	return Coin{}, errors.New("unknown id " + id)
}

func IsEVM(coinID uint) bool {
	return Coins[coinID].Blockchain == BlockchainEthereum
}

// nolint:cyclop
func GetCoinExploreURL(c Coin, tokenID, tokenType string) (string, error) {
	switch c.ID {
	case ETHEREUM:
		return fmt.Sprintf("https://etherscan.io/token/%s", tokenID), nil
	case TRON:
		if _, err := strconv.ParseUint(tokenID, 10, 64); err == nil {
			return fmt.Sprintf("https://tronscan.io/#/token/%s", tokenID), nil
		}

		return fmt.Sprintf("https://tronscan.io/#/token20/%s", tokenID), nil
	case BINANCE:
		return fmt.Sprintf("https://explorer.binance.org/asset/%s", tokenID), nil
	case SMARTCHAIN:
		return fmt.Sprintf("https://bscscan.com/token/%s", tokenID), nil
	case EOS:
		return fmt.Sprintf("https://bloks.io/account/%s", tokenID), nil
	case NEO:
		return fmt.Sprintf("https://neo.tokenview.com/en/token/0x%s", tokenID), nil
	case NULS:
		return fmt.Sprintf("https://nulscan.io/token/info?contractAddress=%s", tokenID), nil
	case WANCHAIN:
		return fmt.Sprintf("https://www.wanscan.org/token/%s", tokenID), nil
	case SOLANA:
		return fmt.Sprintf("https://solscan.io/token/%s", tokenID), nil
	case TOMOCHAIN:
		return fmt.Sprintf("https://tomoscan.io/token/%s", tokenID), nil
	case KAVA:
		return "https://www.mintscan.io/kava", nil
	case ONTOLOGY:
		return "https://explorer.ont.io", nil
	case GOCHAIN:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", tokenID), nil
	case THETA:
		return "https://explorer.thetatoken.org/", nil
	case THUNDERTOKEN:
		return fmt.Sprintf("https://viewblock.io/thundercore/address/%s", tokenID), nil
	case CLASSIC:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/tokens/%s", tokenID), nil
	case VECHAIN:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", tokenID), nil
	case WAVES:
		return fmt.Sprintf("https://wavesexplorer.com/assets/%s", tokenID), nil
	case XDAI:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/tokens/%s", tokenID), nil
	case POA:
		return fmt.Sprintf("https://blockscout.com/poa/core/tokens/%s", tokenID), nil
	case POLYGON:
		return fmt.Sprintf("https://polygonscan.com/token/%s", tokenID), nil
	case OPTIMISM:
		return fmt.Sprintf("https://optimistic.etherscan.io/token/%s", tokenID), nil
	case AVALANCHEC:
		return fmt.Sprintf("https://snowtrace.io/token/%s", tokenID), nil
	case ARBITRUM:
		return fmt.Sprintf("https://arbiscan.io/token/%s", tokenID), nil
	case FANTOM:
		return fmt.Sprintf("https://ftmscan.com/token/%s", tokenID), nil
	case TERRA:
		return fmt.Sprintf("https://finder.terra.money/mainnet/address/%s", tokenID), nil
	case RONIN:
		return fmt.Sprintf("https://explorer.roninchain.com/token/%s", tokenID), nil
	case CELO:
		return fmt.Sprintf("https://explorer.celo.org/mainnet/address/%s", tokenID), nil
	case ELROND:
		if tokenType == "ESDT" {
			return fmt.Sprintf("https://explorer.multiversx.com/tokens/%s", tokenID), nil
		}

		return fmt.Sprintf("https://explorer.multiversx.com/collections/%s", tokenID), nil
	case HECO:
		return fmt.Sprintf("https://hecoinfo.com/token/%s", tokenID), nil
	case OASIS:
		return fmt.Sprintf("https://explorer.oasis.updev.si/token/%s", tokenID), nil
	case CRONOS:
		return fmt.Sprintf("https://cronos.org/explorer/address/%s/token-transfers", tokenID), nil
	case STELLAR:
		return fmt.Sprintf("https://stellar.expert/explorer/public/asset/%s", tokenID), nil
	case KCC:
		return fmt.Sprintf("https://explorer.kcc.io/token/%s", tokenID), nil
	case AURORA:
		return fmt.Sprintf("https://aurorascan.dev/address/%s", tokenID), nil
	case ALGORAND:
		return fmt.Sprintf("https://algoexplorer.io/asset/%s", tokenID), nil
	case KAVAEVM:
		return fmt.Sprintf("https://explorer.kava.io/token/%s", tokenID), nil
	case METER:
		return fmt.Sprintf("https://scan.meter.io/address/%s", tokenID), nil
	case EVMOS:
		return fmt.Sprintf("https://evm.evmos.org/address/%s", tokenID), nil
	case OKC:
		return fmt.Sprintf("https://www.oklink.com/en/okc/address/%s", tokenID), nil
	case APTOS:
		switch tokenType {
		case "APTOSFA":
			return fmt.Sprintf("https://explorer.aptoslabs.com/fungible_asset/%s?network=mainnet", tokenID), nil
		default:
			return fmt.Sprintf("https://explorer.aptoslabs.com/coin/%s?network=mainnet", tokenID), nil
		}
	case MOONBEAM:
		return fmt.Sprintf("https://moonscan.io/token/%s", tokenID), nil
	case KLAYTN:
		return fmt.Sprintf("https://kaiascan.io/token/%s", tokenID), nil
	case METIS:
		return fmt.Sprintf("https://andromeda-explorer.metis.io/token/%s", tokenID), nil
	case MOONRIVER:
		return fmt.Sprintf("https://moonriver.moonscan.io/token/%s", tokenID), nil
	case BOBA:
		return fmt.Sprintf("https://bobascan.com/token/%s", tokenID), nil
	case TON:
		return fmt.Sprintf("https://tonscan.org/address/%s", tokenID), nil
	case POLYGONZKEVM:
		return fmt.Sprintf("https://explorer.public.zkevm-test.net/address/%s", tokenID), nil
	case ZKSYNC:
		return fmt.Sprintf("https://explorer.zksync.io/address/%s", tokenID), nil
	case SUI:
		return fmt.Sprintf("https://explorer.sui.io/address/%s", tokenID), nil
	case STRIDE:
		return fmt.Sprintf("https://www.mintscan.io/stride/account/%s", tokenID), nil
	case NEUTRON:
		return fmt.Sprintf("https://www.mintscan.io/neutron/account/%s", tokenID), nil
	case IOTEXEVM:
		return fmt.Sprintf("https://iotexscan.io/address/%s#transactions", tokenID), nil
	case CRYPTOORG:
		return fmt.Sprintf("https://crypto.org/explorer/account/%s", tokenID), nil
	case TEZOS:
		return fmt.Sprintf("https://tzstats.com/%s", tokenID), nil
	case CFXEVM:
		return fmt.Sprintf("https://evm.confluxscan.net/address/%s", tokenID), nil
	case ACALA:
		if tokenType == "custom_token" {
			return fmt.Sprintf("https://acala.subscan.io/custom_token?customTokenId=%s", tokenID), nil
		}
		return fmt.Sprintf("https://acala.subscan.io/system_token_detail?unique_id=%s", tokenID), nil
	case ACALAEVM:
		return fmt.Sprintf("https://blockscout.acala.network/token/%s", tokenID), nil
	case BASE:
		return fmt.Sprintf("https://basescan.org/token/%s", tokenID), nil
	case CARDANO:
		return fmt.Sprintf("https://cexplorer.io/asset/%s", tokenID), nil
	case NEON:
		return fmt.Sprintf("https://neonscan.org/token/%s", tokenID), nil
	case MANTLE:
		return fmt.Sprintf("https://explorer.mantle.xyz/address/%s", tokenID), nil
	case LINEA:
		return fmt.Sprintf("https://explorer.linea.build/token/%s", tokenID), nil
	case OPBNB:
		return fmt.Sprintf("https://opbnbscan.com/token/%s", tokenID), nil
	case MANTA:
		return fmt.Sprintf("https://pacific-explorer.manta.network/token/%s", tokenID), nil
	case ZETACHAIN:
		return fmt.Sprintf("https://explorer.zetachain.com/address/%s", tokenID), nil
	case ZETAEVM:
		return fmt.Sprintf("https://explorer.zetachain.com/address/%s", tokenID), nil
	case BITCOIN:
		return fmt.Sprintf("https://unisat.io/brc20/%s", tokenID), nil
	case BLAST:
		return fmt.Sprintf("https://blastscan.io/token/%s", tokenID), nil
	case SCROLL:
		return fmt.Sprintf("https://scrollscan.com/token/%s", tokenID), nil
	case ZKLINKNOVA:
		return fmt.Sprintf("https://explorer.zklink.io/address/%s", tokenID), nil
	case RIPPLE:
		return fmt.Sprintf("https://xrpscan.com/account/%s", tokenID), nil
	case SONIC:
		return fmt.Sprintf("https://sonicscan.org/token/%s", tokenID), nil
	case TIA:
		return "https://www.mintscan.io/celestia", nil
	case DYDX:
		return "https://www.mintscan.io/dydx", nil
	case PLASMA:
		return fmt.Sprintf("https://plasmascan.to/token/%s", tokenID), nil
	case MONAD:
		return fmt.Sprintf("https://explorer.monad.xyz/token/%s", tokenID), nil
	}

	return "", errors.New("no explorer for coin: " + c.Handle)
}

// nolint:cyclop
func GetAddressExploreURL(c Coin, address string) (string, error) {
	switch c.ID {
	case ETHEREUM:
		return fmt.Sprintf("https://etherscan.io/address/%s", address), nil
	case TRON:
		return fmt.Sprintf("https://tronscan.io/#/address/%s", address), nil
	case BINANCE:
		return fmt.Sprintf("https://explorer.binance.org/address/%s", address), nil
	case SMARTCHAIN:
		return fmt.Sprintf("https://bscscan.com/address/%s", address), nil
	case EOS:
		return fmt.Sprintf("https://bloks.io/account/%s", address), nil
	case NULS:
		return fmt.Sprintf("https://nulscan.io/token/info?address=%s", address), nil
	case WANCHAIN:
		return fmt.Sprintf("https://www.wanscan.org/address/%s", address), nil
	case SOLANA:
		return fmt.Sprintf("https://solscan.io/account/%s", address), nil
	case TOMOCHAIN:
		return fmt.Sprintf("https://tomoscan.io/address/%s", address), nil
	case KAVA:
		return "https://www.mintscan.io/kava", nil
	case ONTOLOGY:
		return "https://explorer.ont.io", nil
	case GOCHAIN:
		return fmt.Sprintf("https://explorer.gochain.io/addr/%s", address), nil
	case THETA:
		return "https://explorer.thetatoken.org/", nil
	case THUNDERTOKEN:
		return fmt.Sprintf("https://explorer-mainnet.thundercore.com/address/%s", address), nil
	case CLASSIC:
		return fmt.Sprintf("https://blockscout.com/etc/mainnet/address/%s", address), nil
	case VECHAIN:
		return fmt.Sprintf("https://explore.vechain.org/accounts/%s", address), nil
	case WAVES:
		return fmt.Sprintf("https://wavesexplorer.com/addresses/%s", address), nil
	case XDAI:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/address/%s", address), nil
	case POLYGON:
		return fmt.Sprintf("https://polygonscan.com/address/%s", address), nil
	case OPTIMISM:
		return fmt.Sprintf("https://optimistic.etherscan.io/address/%s", address), nil
	case AVALANCHEC:
		return fmt.Sprintf("https://snowtrace.io/address/%s", address), nil
	case ARBITRUM:
		return fmt.Sprintf("https://arbiscan.io/address/%s", address), nil
	case FANTOM:
		return fmt.Sprintf("https://explorer.fantom.network/address/%s", address), nil // not working
	case TERRA:
		return fmt.Sprintf("https://finder.terra.money/mainnet/address/%s", address), nil
	case RONIN:
		return fmt.Sprintf("https://explorer.roninchain.com/address/%s", address), nil
	case CELO:
		return fmt.Sprintf("https://explorer.celo.org/mainnet/address/%s", address), nil
	case ELROND:
		return fmt.Sprintf("https://explorer.multiversx.com/accounts/%s", address), nil
	case CRONOS:
		return fmt.Sprintf("https://explorer.cronos.org/address/%s", address), nil
	case STELLAR:
		return fmt.Sprintf("https://stellar.expert/explorer/public/account/%s", address), nil
	case KCC:
		return fmt.Sprintf("https://explorer.kcc.io/address/%s", address), nil
	case AURORA:
		return fmt.Sprintf("https://aurorascan.dev/address/%s", address), nil
	case KAVAEVM:
		return fmt.Sprintf("https://explorer.kava.io/address/%s", address), nil
	case METER:
		return fmt.Sprintf("https://scan.meter.io/address/%s", address), nil
	case APTOS:
		return fmt.Sprintf("https://explorer.aptoslabs.com/account/%s?network=mainnet", address), nil
	case MOONBEAM:
		return fmt.Sprintf("https://moonscan.io/address/%s", address), nil
	case KLAYTN:
		return fmt.Sprintf("https://kaiascan.io/address/%s", address), nil
	case METIS:
		return fmt.Sprintf("https://andromeda-explorer.metis.io/address/%s", address), nil
	case MOONRIVER:
		return fmt.Sprintf("https://moonriver.moonscan.io/address/%s", address), nil
	case BOBA:
		return fmt.Sprintf("https://bobascan.com/address/%s", address), nil
	case TON:
		return fmt.Sprintf("https://tonscan.org/address/%s", address), nil
	case POLYGONZKEVM:
		return fmt.Sprintf("https://explorer.public.zkevm-test.net/address/%s", address), nil
	case ZKSYNC:
		return fmt.Sprintf("https://explorer.zksync.io/address/%s", address), nil
	case SUI:
		return fmt.Sprintf("https://explorer.sui.io/address/%s", address), nil
	case STRIDE:
		return fmt.Sprintf("https://www.mintscan.io/stride/account/%s", address), nil
	case NEUTRON:
		return fmt.Sprintf("https://www.mintscan.io/neutron/account/%s", address), nil
	case IOTEXEVM:
		return fmt.Sprintf("https://iotexscan.io/address/%s#transactions", address), nil
	case CRYPTOORG:
		return fmt.Sprintf("https://crypto.org/explorer/account/%s", address), nil
	case TEZOS:
		return fmt.Sprintf("https://tzstats.com/%s", address), nil
	case CFXEVM:
		return fmt.Sprintf("https://evm.confluxscan.net/address/%s", address), nil
	case ACALA:
		return fmt.Sprintf("https://acala.subscan.io/account/%s", address), nil
	case ACALAEVM:
		return fmt.Sprintf("https://blockscout.acala.network/address/%s", address), nil
	case BASE:
		return fmt.Sprintf("https://basescan.org/address/%s", address), nil
	case CARDANO:
		return fmt.Sprintf("https://cexplorer.io/address/%s", address), nil
	case NEON:
		return fmt.Sprintf("https://neonscan.org/address/%s", address), nil
	case MANTLE:
		return fmt.Sprintf("https://explorer.mantle.xyz/address/%s", address), nil
	case LINEA:
		return fmt.Sprintf("https://explorer.linea.build/address/%s", address), nil
	case OPBNB:
		return fmt.Sprintf("https://opbnbscan.com/address/%s", address), nil
	case MANTA:
		return fmt.Sprintf("https://pacific-explorer.manta.network/address/%s", address), nil
	case ZETACHAIN:
		return fmt.Sprintf("https://explorer.zetachain.com/address/%s", address), nil
	case ZETAEVM:
		return fmt.Sprintf("https://explorer.zetachain.com/address/%s", address), nil
	case BLAST:
		return fmt.Sprintf("https://blastscan.io/address/%s", address), nil
	case SCROLL:
		return fmt.Sprintf("https://scrollscan.com/address/%s", address), nil
	case ZKLINKNOVA:
		return fmt.Sprintf("https://explorer.zklink.io/address/%s", address), nil
	case RIPPLE:
		return fmt.Sprintf("https://xrpscan.com/account/%s", address), nil
	case SONIC:
		return fmt.Sprintf("https://sonicscan.org/address/%s", address), nil
	case TIA:
		return "https://www.mintscan.io/celestia", nil
	case DYDX:
		return "https://www.mintscan.io/dydx", nil
	case PLASMA:
		return fmt.Sprintf("https://plasmascan.to/address/%s", address), nil
	case MONAD:
		return fmt.Sprintf("https://explorer.monad.xyz/address/%s", address), nil
	}

	return "", errors.New("no explorer for coin: " + c.Handle)
}
