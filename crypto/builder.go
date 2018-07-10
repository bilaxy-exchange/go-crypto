// This file is part of Ark Go Crypto.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

func buildSignedTransaction(transaction *Transaction, passphrase string, secondPassphrase string) *Transaction {
	transaction.Timestamp = GetTime()
	transaction.Sign(passphrase)

	if len(secondPassphrase) > 0 {
		transaction.SecondSign(secondPassphrase)
	}

	transaction.Id = transaction.GetId()

	return transaction
}

func BuildTransfer(recipient string, amount uint64, vendorField string, passphrase string, secondPassphrase string) *Transaction {
	transaction := &Transaction{
		Type:        TRANSACTION_TYPES.Transfer,
		Fee:         TRANSACTION_FEES.Transfer,
		RecipientId: recipient,
		Amount:      amount,
		VendorField: vendorField,
		Asset:       &TransactionAsset{},
	}

	return buildSignedTransaction(transaction, passphrase, secondPassphrase)
}

func BuildSecondSignatureRegistration(passphrase string, secondPassphrase string) *Transaction {
	transaction := &Transaction{
		Type:  TRANSACTION_TYPES.SecondSignatureRegistration,
		Fee:   TRANSACTION_FEES.SecondSignatureRegistration,
		Asset: &TransactionAsset{},
	}

	publicKey, _ := PublicKeyFromPassphrase(passphrase)

	transaction.Asset.Signature = &SecondSignatureRegistrationAsset{
		PublicKey: HexEncode(publicKey.Serialize()),
	}

	return buildSignedTransaction(transaction, passphrase, secondPassphrase)
}

func BuildDelegateRegistration(username string, passphrase string, secondPassphrase string) *Transaction {
	transaction := &Transaction{
		Type:  TRANSACTION_TYPES.DelegateRegistration,
		Fee:   TRANSACTION_FEES.DelegateRegistration,
		Asset: &TransactionAsset{},
	}

	transaction.Asset.Delegate = &DelegateAsset{
		Username: username,
	}

	return buildSignedTransaction(transaction, passphrase, secondPassphrase)
}

func BuildVote(vote, passphrase, secondPassphrase string) *Transaction {
	transaction := &Transaction{
		Type:  TRANSACTION_TYPES.Vote,
		Fee:   TRANSACTION_FEES.Vote,
		Asset: &TransactionAsset{},
	}

	transaction.RecipientId, _ = AddressFromPassphrase(passphrase)
	transaction.Asset.Votes = append(transaction.Asset.Votes, vote)

	return buildSignedTransaction(transaction, passphrase, secondPassphrase)
}

// func BuildMultiSignatureRegistration() *Transaction {}