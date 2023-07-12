package messages

import (
	"fmt"
	"strings"

	"github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"

	// authztypes "github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/forbole/juno/v5/types"
)

// MessageNotSupported returns an error telling that the given message is not supported
func MessageNotSupported(msg sdk.Msg) error {
	return fmt.Errorf("message type not supported: %s", proto.MessageName(msg))
}

// MessageAddressesParser represents a function that extracts all the
// involved addresses from a provided message (both accounts and validators)
type MessageAddressesParser = func(tx *types.Tx) ([]string, error)

// CosmosMessageAddressesParser represents a MessageAddressesParser that parses a
// Chain message and returns all the involved addresses (both accounts and validators)
var CosmosMessageAddressesParser = DefaultMessagesParser

// DefaultMessagesParser represents the default messages parser that simply returns the list
// of all the signers of a message
func DefaultMessagesParser(tx *types.Tx, chainPrefix string) ([]string, error) {
	allAddressess := parseAddressesFromEvents(tx, chainPrefix)
	return allAddressess, nil
}

// function to remove duplicate values
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func parseAddressesFromEvents(tx *types.Tx, chainPrefix string) []string {
	var allAddressess []string
	for _, cc := range tx.Events {
		for _, dd := range cc.Attributes {
			if strings.Contains(dd.Value, "/") {
				continue
			}
			if strings.Contains(dd.Value, chainPrefix) {
				allAddressess = append(allAddressess, dd.Value)
			}
		}

	}
	allInvolvedAddresses := removeDuplicates(allAddressess)

	fmt.Printf("\n all involved addresses %v \n", allInvolvedAddresses)

	return allInvolvedAddresses
}
