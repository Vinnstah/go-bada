# Requirements

## Coordinator 
Input - List of `Transactions` that should be signed. Should be signed by enough `FactorSources` so that all transactions are valid.

Each `Transaction` might have between 1-n `Accounts` required to sign.

Each `Account` can either have single factor or multifactor. Mutlifactor is a "mode" and might require several factors but does not require several factors. Account have a property, `securityState` which is an enum with two cases; `securified` (associated value: `MatrixOfFactorInstances`) and `unsecurified` (associated value one `FactorInstance`). `Account` also has an `Address` which can be a tuple of `Name` (String) and `Index` (u8) matching all `DerivationPath` indicies. 

`MatrixOfFactorInstances` have 3 properties, 2 of which as lists of `FactorInstances`. `ThresholdFactors` is one of the lists. The other is `OverideFactors`. The thrid property is `ThresholdCount`, which describes n of m factors needs in order to sign. The same `FactorInstance` can not occur in both lists, and atleast one list need atleast one `FactorInstance`.

Enum `FactorSourceKind`: `Device`, `Ledger`, `Arculus`, `Yubikey`, `OffDeviceMnemonic` and `SecurityQuestion`. Signing order is: Ledger, SecurityQuestion, OffDeviceMnemonic, Arculus, Yubikey and lastly Device.

`FactorSourceKind` `Device` supports paralell signing. `Coordinator` will call a request on `SignWithFactorSourcesInParalellDriver` passing all `Device` `FactorSources` in one go. All the other kinds; one `FactorSource` at a time. There a 2 different types of requests: `BatchParalellSigningRequest` and `BatchSerialSigningRequest`.

`BatchSerialSigningRequest` takes a `FactorSourceID` and HashMap with `TransactionID` as key and list of `DerivationPath` to sigh the `TransactionID` with. 

`BatchParalellSigningRequest` contains a HashMap with `FactorSourceID` as key and `BatchSerialSigningRequest` as value. The key `FactorSourceID` in `BatchParalellSigningRequest` should match the value in `BatchSerialSigningRequest`

`BatchSerialSigningResponse` contains a HashMap with `TransactionID` as key and a list of `Signature`s as value. A `Signature` contains the `FactorSourceID`, `TransactionID` and `DerivationPath`.

`BatchParalellSigningResponse` has a HashMap with `FactorSourceID` as key and `BatchSerialSigningResponse` as values. 

The user has the option to skip signing with a `FactorSource` and before doing so should be prompted which `Transaction`s would fail by doing so.

`FactorSourceID` is a struct, which has a `FactorSourceKind` and a `UUID`

`FactorSource` has a `FactorSourceID`

`FactorInstance` has a reference to the `FactorSource` which created it, which is the `FactorSourceID`

`FactorInstance` also has a `DerivationPath` which is a `u8`.

`Transaction` has a `TransactionID` which is a `UUID` and a list of `AccountsRequiredToSign`. It is the `TransactionID` which is supposed to be signed.

One public func `batchSign`. Most likely async. Should return two lists of `SucessfullySignedTransactions` and `UnsuccesfullySignedTransactions`.

The whole `batchSign` should return early if all `Transactions` are already valid and signed. 

# Note
We can use `FactorSource` and `FactorSourceID` interchangably. 