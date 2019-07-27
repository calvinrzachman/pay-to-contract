# pay-to-contract
A simple Go implementation of Pay To Contract

[![GoDoc](https://camo.githubusercontent.com/8609cfcb531fa0f5598a3d4353596fae9336cce3/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f79616e6777656e6d61692f686f772d746f2d6164642d62616467652d696e2d6769746875622d726561646d653f7374617475732e737667)](https://godoc.org/github.com/calvinrzachman/pay-to-contract)

## Overview
Pay to Contract is a scheme which extends the classic Bitcoin Pay-to-Public-Key model, providing a path to notable improvements to the privacy, security, and fungibility of Bitcoin payments. Most basically, Pay to Contract is useful as it enables us to encode more information in a single public key and then use this information to provide more "spend-time" flexibility. It achieves this through alteration of the public key which allows for the encoding of arbitrary data.

The ability to have the public key commit to arbitrary data at "send-time" enables... and forms the underpinnings of [Taproot](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2019-May/016914.html) Bitcoin Improvement Proposal
 
Pay to Contrat/Taproot together with Schnorr signatures and a collaborative signature scheme like [MuSig](https://eprint.iacr.org/2018/068.pdf), allows any and all multi-signature and complex script based transactions to hit the blockchain with the same form as a normal Pay to Public Key Hash (P2PKH).

This allows for smart contract structure in which spends with the approval of all contract participants, such as collaborative closes of channels on Lightning Network, appear no differently than any other transaction.

Under [BIP-Taproot](https://github.com/sipa/bips/blob/bip-schnorr/bip-taproot.mediawiki) the data is itself a commitment in form of a merkle root whose leaves are comprised of Bitcoin scripts specifying the various spending conditions for the transaction output (see [M.A.S.T](https://bitcointechtalk.com/what-is-a-bitcoin-merklized-abstract-syntax-tree-mast-33fdf2da5e2f)).

Taproot Pay to Contract outputs can be spent in 2 ways: 
- Key spend - provide a BIP-Schnorr signature for the tweaked public key
- Script spend - reveal the root of a merklized abstract syntax tree (MAST), a proof that the spending condition connects to the merkle root, and finally the script inputs 

## Usage 
Install using:
 
    go get -u github.com/calvinrzachman/pay-to-contract

In your code:

```go
import "github.com/calvinrzachman/pay-to-contract"

```

## Note

This implementation is for learning purposes only. Code from this repository should not be used in a production environment.

### Helpful Links
[Taproot Is Coming: What It Is, and How It Will Benefit Bitcoin](https://bitcoinmagazine.com/articles/taproot-coming-what-it-and-how-it-will-benefit-bitcoin)
