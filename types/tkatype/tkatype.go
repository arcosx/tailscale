// Copyright (c) 2022 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tkatype defines types for working with the tka package.
//
// Do not add extra dependencies to this package unless they are tiny,
// because this package encodes wire types that should be lightweight to use.
package tkatype

// KeyID references a verification key stored in the key authority. A keyID
// uniquely identifies a key. KeyIDs are all 32 bytes.
//
// For 25519 keys: We just use the 32-byte public key.
//
// Even though this is a 32-byte value, we use a byte slice because
// CBOR-encoded byte slices have a different prefix to CBOR-encoded arrays.
// Encoding as a byte slice allows us to change the size in the future if we
// ever need to.
type KeyID []byte

// MarshaledSignature represents a marshaled tka.NodeKeySignature.
type MarshaledSignature []byte

// AUMSigHash represents the BLAKE2s digest of an Authority Update
// Message (AUM), sans any signatures.
type AUMSigHash [32]byte

// Signature describes a signature over an AUM, which can be verified
// using the key referenced by KeyID.
type Signature struct {
	KeyID     KeyID  `cbor:"1,keyasint"`
	Signature []byte `cbor:"2,keyasint"`
}
