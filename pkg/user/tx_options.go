package user

import (
	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type TxOption func(builder sdkclient.TxBuilder) sdkclient.TxBuilder

func SetGasLimit(limit uint64) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetGasLimit(limit)
		return builder
	}
}

func SetFeeAmount(fees sdk.Coins) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetFeeAmount(fees)
		return builder
	}
}

func SetFee(fees uint64) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(appconsts.BondDenom, sdk.NewInt(int64(fees)))))
		return builder
	}
}

func SetMemo(memo string) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetMemo(memo)
		return builder
	}
}

func SetFeePayer(feePayer sdk.AccAddress) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetFeePayer(feePayer)
		return builder
	}
}

func SetTip(tip *tx.Tip) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetTip(tip)
		return builder
	}
}

func SetTimeoutHeight(height uint64) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetTimeoutHeight(height)
		return builder
	}
}

func SetFeeGranter(feeGranter sdk.AccAddress) TxOption {
	return func(builder sdkclient.TxBuilder) sdkclient.TxBuilder {
		builder.SetFeeGranter(feeGranter)
		return builder
	}
}

// InheritTxConfig sets all of the accessible configurations from a given tx
// into a given client.TxBuilder
func InheritTxConfig(builder sdkclient.TxBuilder, tx authsigning.Tx) sdkclient.TxBuilder {
	if gas := tx.GetGas(); gas != 0 {
		builder.SetGasLimit(gas)
	}

	if feeAmmount := tx.GetFee(); !feeAmmount.AmountOf("utia").Equal(sdk.NewInt(0)) {
		builder.SetFeeAmount(tx.GetFee())
	}

	if memo := tx.GetMemo(); memo != "" {
		builder.SetMemo(tx.GetMemo())
	}

	if tip := tx.GetTip(); tip != nil {
		builder.SetTip(tip)
	}

	if timeoutHeight := tx.GetTimeoutHeight(); timeoutHeight != 0 {
		builder.SetTimeoutHeight(timeoutHeight)
	}

	signers := tx.GetSigners()
	// Note: if there are multiple signers in a PFB, then this could create an
	// invalid signature. This is not an issue at this time because we currently
	// ignore pfbs with multiple signers
	if len(signers) == 1 {
		if feePayer := tx.FeeGranter(); !feePayer.Equals(signers[0]) {
			builder.SetFeeGranter(tx.FeeGranter())
		}
	}

	if feeGranter := tx.FeeGranter(); !feeGranter.Empty() {
		builder.SetFeeGranter(tx.FeeGranter())
	}

	return builder
}
