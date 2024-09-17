package query

import (
	"fmt"
	"strings"
)

// Query define all the queries
type Query struct {
	Key              string
	EndpointTemplate string
}

// Endpoint return the end point string
func (q Query) Endpoint(args ...string) string {
	count := strings.Count(q.EndpointTemplate, "%s")
	a := args[:count]

	in := make([]interface{}, len(a))
	for i := range in {
		in[i] = a[i]
	}

	return fmt.Sprintf(q.EndpointTemplate, in...)
}

// Path return the path
func (q Query) Path(args ...string) string {
	temp := []string{args[0], q.Key}
	args = append(temp, args[1:]...)
	return fmt.Sprintf("custom/%s", strings.Join(args, "/"))
}

// query endpoints supported by the thorchain Querier
var (
	QueryTx                  = Query{Key: "tx", EndpointTemplate: "/%s/tx/{%s}"}
	QueryTxVoterOld          = Query{Key: "txvoterold", EndpointTemplate: "/%s/tx/{%s}/signers"}
	QueryTxVoter             = Query{Key: "txvoter", EndpointTemplate: "/%s/tx/details/{%s}"}
	QueryTxStages            = Query{Key: "txstages", EndpointTemplate: "/%s/tx/stages/{%s}"}
	QueryTxStatus            = Query{Key: "txstatus", EndpointTemplate: "/%s/tx/status/{%s}"}
	QueryKeysignArray        = Query{Key: "keysign", EndpointTemplate: "/%s/keysign/{%s}"}
	QueryKeysignArrayPubkey  = Query{Key: "keysignpubkey", EndpointTemplate: "/%s/keysign/{%s}/{%s}"}
	QueryKeygensPubkey       = Query{Key: "keygenspubkey", EndpointTemplate: "/%s/keygen/{%s}/{%s}"}
	QueryQueue               = Query{Key: "outqueue", EndpointTemplate: "/%s/queue"}
	QueryHeights             = Query{Key: "heights", EndpointTemplate: "/%s/lastblock"}
	QueryChainHeights        = Query{Key: "chainheights", EndpointTemplate: "/%s/lastblock/{%s}"}
	QueryInboundAddresses    = Query{Key: "inboundaddresses", EndpointTemplate: "/%s/inbound_addresses"}
	QueryNetwork             = Query{Key: "network", EndpointTemplate: "/%s/network"}
	QueryPOL                 = Query{Key: "pol", EndpointTemplate: "/%s/pol"}
	QueryBalanceModule       = Query{Key: "balancemodule", EndpointTemplate: "/%s/balance/module/{%s}"}
	QueryVaultsAsgard        = Query{Key: "vaultsasgard", EndpointTemplate: "/%s/vaults/asgard"}
	QueryVault               = Query{Key: "vault", EndpointTemplate: "/%s/vault/{%s}"}
	QueryVaultPubkeys        = Query{Key: "vaultpubkeys", EndpointTemplate: "/%s/vaults/pubkeys"}
	QueryConstantValues      = Query{Key: "constants", EndpointTemplate: "/%s/constants"}
	QueryVersion             = Query{Key: "version", EndpointTemplate: "/%s/version"}
	QueryMimirValues         = Query{Key: "mimirs", EndpointTemplate: "/%s/mimir"}
	QueryMimirWithKey        = Query{Key: "mimirwithkey", EndpointTemplate: "/%s/mimir/key/{%s}"}
	QueryMimirAdminValues    = Query{Key: "adminmimirs", EndpointTemplate: "/%s/mimir/admin"}
	QueryMimirNodesValues    = Query{Key: "nodesmimirs", EndpointTemplate: "/%s/mimir/nodes"}
	QueryMimirNodesAllValues = Query{Key: "nodesmimirsall", EndpointTemplate: "/%s/mimir/nodes_all"}
	QueryMimirNodeValues     = Query{Key: "nodemimirs", EndpointTemplate: "/%s/mimir/node/{%s}"}
	QueryPendingOutbound     = Query{Key: "pendingoutbound", EndpointTemplate: "/%s/queue/outbound"}
	QueryScheduledOutbound   = Query{Key: "scheduledoutbound", EndpointTemplate: "/%s/queue/scheduled"}
	QuerySwapperClout        = Query{Key: "swapperclout", EndpointTemplate: "/%s/clout/swap/{%s}"}
	QuerySwapQueue           = Query{Key: "swapqueue", EndpointTemplate: "/%s/queue/swap"}
	QueryTssKeygenMetrics    = Query{Key: "tss_keygen_metric", EndpointTemplate: "/%s/metric/keygen/{%s}"}
	QueryTssMetrics          = Query{Key: "tss_metric", EndpointTemplate: "/%s/metrics"}
	QueryTHORName            = Query{Key: "thorname", EndpointTemplate: "/%s/thorname/{%s}"}
	QueryQuoteSwap           = Query{Key: "quoteswap", EndpointTemplate: "/%s/quote/swap"}
	QueryQuoteSaverDeposit   = Query{Key: "quotesaverdeposit", EndpointTemplate: "/%s/quote/saver/deposit"}
	QueryQuoteSaverWithdraw  = Query{Key: "quotesaverwithdraw", EndpointTemplate: "/%s/quote/saver/withdraw"}
	QueryQuoteLoanOpen       = Query{Key: "quoteloanopen", EndpointTemplate: "/%s/quote/loan/open"}
	QueryQuoteLoanClose      = Query{Key: "quoteloanclose", EndpointTemplate: "/%s/quote/loan/close"}
	QueryInvariants          = Query{Key: "invariants", EndpointTemplate: "/%s/invariants"}
	QueryInvariant           = Query{Key: "invariant", EndpointTemplate: "/%s/invariant/{%s}"}
	QueryBlock               = Query{Key: "block", EndpointTemplate: "/%s/block"}

	// queries only available on regtest builds
	QueryExport = Query{Key: "export", EndpointTemplate: "/%s/export"}
)

// Queries all queries
var Queries = []Query{
	QueryTxStages,
	QueryTxStatus,
	QueryTxVoter,
	QueryTxVoterOld,
	QueryTx,
	QueryKeysignArray,
	QueryKeysignArrayPubkey,
	QueryQueue,
	QueryHeights,
	QueryChainHeights,
	QueryInboundAddresses,
	QueryNetwork,
	QueryPOL,
	QueryBalanceModule,
	QueryVaultsAsgard,
	QueryVaultPubkeys,
	QueryVault,
	QueryKeygensPubkey,
	QueryConstantValues,
	QueryVersion,
	QueryMimirValues,
	QueryMimirWithKey,
	QueryMimirAdminValues,
	QueryMimirNodesAllValues,
	QueryMimirNodesValues,
	QueryMimirNodeValues,
	QueryPendingOutbound,
	QueryScheduledOutbound,
	QuerySwapperClout,
	QuerySwapQueue,
	QueryTssMetrics,
	QueryTssKeygenMetrics,
	QueryTHORName,
	QueryQuoteSwap,
	QueryQuoteSaverDeposit,
	QueryQuoteSaverWithdraw,
	QueryQuoteLoanOpen,
	QueryQuoteLoanClose,
	QueryInvariants,
	QueryInvariant,
	QueryBlock,
	QueryExport,
}
