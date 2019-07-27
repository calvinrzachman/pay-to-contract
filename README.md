# pay-to-contract
A simple Go implementation of Pay To Contract


## Overview
Pay to Contract is a scheme which extends the classic Bitcoin Pay-to-Public-Key model, bringing notable improvements to the privacy, security, and fungibility of Bitcoin payments. It achieves this through alteration of the public key which allows for the encoding of arbitrary data.

The ability to have the public key commit to arbitrary data at "send-time" enables...
This process forms the underpinnings of [Taproot](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2019-May/016914.html)
 
Pay to Contrat/Taproot together with Schnorr signatures and a collaborative signature scheme like [MuSig](https://eprint.iacr.org/2018/068.pdf), allows any and all multi-signature and complex script based transactions to hit the blockchain with the same form as a normal Pay to Public Key Hash (P2PKH).

This allows for smart contract structure in which spends with the approval of all contract participants, such as collaborative closes of channels on Lightning Network, appear no different than any other transaction.



To spend a Pay to Contract ouput...

 

### Helpful Links
[Taproot Is Coming: What It Is, and How It Will Benefit Bitcoin](https://bitcoinmagazine.com/articles/taproot-coming-what-it-and-how-it-will-benefit-bitcoin)
