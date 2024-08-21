package models

import (
	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/btree"
	"github.com/google/uuid"
)

type FactorSourceKind string

const (
	DEVICE              FactorSourceKind = "Device"
	LEDGER              FactorSourceKind = "Ledger"
	ARCULUS             FactorSourceKind = "Arculus"
	YUBI_KEY            FactorSourceKind = "YubiKey"
	OFF_DEVICE_MNEMONIC FactorSourceKind = "OffDeviceMnemonic"
	SECURITY_QUESTION   FactorSourceKind = "SecurityQuestion"
)

// / MATRIX OF FACTOR INSTANCES
type FactorInstanceSet btree.BTreeG[FactorInstance]

// Add interface unifying the two types
type ThresholdFactors FactorInstanceSet
type OverrideFactors FactorInstanceSet

type MatrixOfFactorInstances struct {
	thresholdFactors ThresholdFactors
	threshold        uint8
	overrideFactors  OverrideFactors
}

// / FACTOR SOURCE
type FactorSourceID struct {
	factorSourceKind FactorSourceKind
	id               uuid.UUID
}

type DerivationIndex uint8
type DerivationPath DerivationIndex

type FactorInstance struct {
	factorSourceID FactorSourceID
	derivationPath DerivationPath
}

// / ACCOUNT
type Account struct {
	address AccountAddress
}

type AccountAddress struct {
	name            string
	derivationIndex DerivationIndex
}

// / SECURITY STATE
type SecurityState interface {
	__State__()
}

type SecurityStateSecurified struct {
	matrixOfFactorInstances MatrixOfFactorInstances
}

type SecurityStateUnsecurified FactorInstance

func (s SecurityStateSecurified) __State__() {}

func (s SecurityStateUnsecurified) __State__() {}

// / TRANSACTION
type AccountsToSign btree.BTreeG[Account]
type TransactionID uuid.UUID

type Transaction struct {
	accountsToSign AccountsToSign
	id             TransactionID
}

// / SIGNATURES
type Signature struct {
	factorSourceID FactorSourceID
	transactionID  TransactionID
	derivationPath DerivationPath
}

// https://github.com/elliotchance/orderedmap
type SignedTransactions orderedmap.OrderedMap[Transaction, btree.BTreeG[Signature]]

type SucessfullySignedTransactions SignedTransactions
type UnsucessfullySignedTransactions SignedTransactions

type BatchSigningResult struct {
	sucessfullySignedTransactions   SucessfullySignedTransactions
	unsucessfullySignedTransactions UnsucessfullySignedTransactions
}

/// goroutine async
func batchSign(transactions btree.BTreeG[Transaction]) (BatchSigningResult, error) {
	return BatchSigningResult{}, nil
}