package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewPacketID(portID, channelID string, sequence uint64) PacketID {
	return PacketID{
		ChannelId: channelID,
		PortId:    portID,
		Sequence:  sequence,
	}
}

// NewFee creates and returns a new Fee struct encapsulating the receive, acknowledgement and timeout fees as sdk.Coins
func NewFee(recvFee, ackFee, timeoutFee sdk.Coins) Fee {
	return Fee{
		RecvFee:    recvFee,
		AckFee:     ackFee,
		TimeoutFee: timeoutFee,
	}
}

// Total returns the total amount for a given Fee
func (m Fee) Total() sdk.Coins {
	return m.RecvFee.Add(m.AckFee...).Add(m.TimeoutFee...)
}

// Validate asserts that each Fee is valid and all three Fees are not empty or zero
func (m Fee) Validate() error {
	var errFees []string
	if !m.AckFee.IsValid() {
		errFees = append(errFees, "ack feerefunder invalid")
	}
	if !m.RecvFee.IsValid() {
		errFees = append(errFees, "recv feerefunder invalid")
	}
	if !m.TimeoutFee.IsValid() {
		errFees = append(errFees, "timeout feerefunder invalid")
	}

	if len(errFees) > 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "contains invalid fees: %s", strings.Join(errFees, " , "))
	}

	// if all three feerefunder's are zero or empty return an error
	if m.AckFee.IsZero() && m.RecvFee.IsZero() && m.TimeoutFee.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "all fees are zero")
	}

	return nil
}
