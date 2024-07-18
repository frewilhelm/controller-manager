/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package digest

import (
	"crypto"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"fmt"

	"github.com/opencontainers/go-digest"
	_ "github.com/opencontainers/go-digest/blake3"
)

const (
	SHA1 digest.Algorithm = "sha1"
)

var (
	// Canonical is the primary digest algorithm used to calculate checksums.
	Canonical = digest.SHA256
)

func init() {
	// Register SHA-1 algorithm for support of e.g. Git commit SHAs.
	digest.RegisterAlgorithm(SHA1, crypto.SHA1)
}

// AlgorithmForName returns the digest algorithm for the given name, or an
// error of type digest.ErrDigestUnsupported if the algorithm is unavailable.
func AlgorithmForName(name string) (digest.Algorithm, error) {
	a := digest.Algorithm(name)
	if !a.Available() {
		return "", fmt.Errorf("%w: %s", digest.ErrDigestUnsupported, name)
	}
	return a, nil
}
