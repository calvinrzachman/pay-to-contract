# pay-to-contract
A simple Go implementation of Pay To Contract

## Overview
Pay to Contract is a scheme which extends the classic Bitcoin Pay-to-Public-Key model, bringing notable improvements to the privacy, security, and fungibility of Bitcoin payments. It achieves this through alteration of the public key which allows for the encoding of arbitrary data.

The ability to have the public key commit to arbitrary data at "send-time" enables...
This process forms the underpinnings of [Taproot](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2019-May/016914.html)
 
Pay to Contrat/Taproot together with Schnorr signatures and a collaborative signature scheme like [MuSig](https://eprint.iacr.org/2018/068.pdf), allows any and all multi-signature and complex script based transactions to hit the blockchain with the same form as a normal Pay to Public Key Hash (P2PKH).

This allows for smart contract structure in which spends with the approval of all contract participants, such as collaborative closes of channels on Lightning Network, to appear no different than any other transaction.

Under [BIP-Taproot](https://github.com/sipa/bips/blob/bip-schnorr/bip-taproot.mediawiki) the data is itself a commitment in form of a merkle root whose leaves are comprised of Bitcoin scripts specifying the various spending conditions for the transaction output (see [M.A.S.T](https://bitcointechtalk.com/what-is-a-bitcoin-merklized-abstract-syntax-tree-mast-33fdf2da5e2f)).

Pay to Contract outputs can be spent in 2 ways: 
- Key spend - provide a BIP-Schnorr signature for the tweaked public key
- Script spend - reveal the root of a merklized abstract syntax tree (MAST), a proof that the spending condition connects to the merkle root, and finally the script inputs 

## Note

This implementation is for learning purposes only. Code from this repository should not be used in a production environment.

### Helpful Links
[Taproot Is Coming: What It Is, and How It Will Benefit Bitcoin](https://bitcoinmagazine.com/articles/taproot-coming-what-it-and-how-it-will-benefit-bitcoin)
