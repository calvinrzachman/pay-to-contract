package paytocontract

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

/*
	PAY TO CONTRACT DEFINITIONS

	KEY PAIR
	let (x, P) be a public/private key pair with the public key P given by:

		P = x*G									(1)

	for private key x and elliptic curve generator G

	TWEAKED PUBLIC KEY
	An altered public key derived from an internal public key P and a non-random
	elliptic curve point C = c*G. The equation is as follows:

		Q = P + C								(2)

	Adding the internal public key P to another elliptic curve point C allows
	for the encoding of arbitrary data into the tweaked public key through user
	selection of c. Most basically, this is useful as it enables us to encode more information
	in a single public key and then use this information to provide more "spend-time" flexibility


	TWEAKED SECRET
	Schnorr signatures spending UTXO locked to the tweaked public key Q can easily be spent
	by modifying the original private key x prior to signing. Substitution of (1) into (2)
	gives:

		Q = x*G + c*G
		Q = ( x + c )*G = w*G					(3)

	with w = (x+c) as the shared secret

*/

var (
	// Curve represents the secp256k1 elliptic curve
	Curve = btcec.S256()

	// N represents the size of the finite cyclic group defined over the given elliptic curve
	N = Curve.Params().N

	// P represents the prime order of the finite field over which we take the elliptic curve...
	P = Curve.Params().P
)

/*

	PAY TO CONTRACT METHODS

	Define the various methods for implementing Pay to Contract (a Taproot precursor)

	Step 1: Determine whether we can tweak a public key - have it commit
	to arbitrary data and still produce a valid Schnorr signature (CONFIRMED)

	Step 2: Implement the M.A.S.T component of Taproot - The data
	to commit to will be the merkle root of this tree

*/

// TweakAdd alters a given public key (P) committing it to arbitrary data and returns
// an elliptic curve point (qX, qY) representing the tweaked key.
// Payments to the tweaked key (Q) commit funds to knowledge of this data.
// For this reason the data is referred to as a contract. The contract is the hash digest
// of the data and is interpreted as a 32-byte big endian integer.
// Under BIP-Taproot the data is itself a commitment in form of a merkle root
// whose leaves are comprised of Bitcoin scripts specifying the various spending
// conditions for the transaction output (see M.A.S.T)
func TweakAdd(publicKey *ecdsa.PublicKey, contract [32]byte) (qX, qY *big.Int) {
	// Calculate the contract point C = H(...arbitrary data)*G
	cX, cY := Curve.ScalarBaseMult(contract[:])

	// Add internal public key to contract point Q = P + c*G
	qX, qY = Curve.Add(publicKey.X, publicKey.Y, cX, cY)
	return qX, qY
}

// TweakSecretKey computes the modified secret key (w) corresponding to the
// tweaked public Q = w*G which can be used to spend UTXO with a ScriptPubKey
// which commits to a tweaked public key.
// Usage: Modify the original private key (x) prior to signing
func TweakSecretKey(x *big.Int, contract [32]byte) *big.Int {
	// Compute the tweaked secret ( x + c )
	c := new(big.Int).SetBytes(contract[:])
	tweakSecret := new(big.Int)
	tweakSecret.Add(x, c)
	tweakSecret.Mod(tweakSecret, Curve.N)
	return tweakSecret
}
