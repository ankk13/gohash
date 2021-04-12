package gohash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/fnv"
)

// Client manage all hash function
type Client struct{}

// FNV32 hashes using fnv32 algorithm
func (c *Client) FNV32(text string) uint32 {
	algorithm := fnv.New32()
	return uint32Hasher(algorithm, text)
}

// FNV32a hashes using fnv32a algorithm
func (c *Client) FNV32a(text string) uint32 {
	algorithm := fnv.New32a()
	return uint32Hasher(algorithm, text)
}

// FNV64 hashes using fnv64 algorithm
func (c *Client) FNV64(text string) uint64 {
	algorithm := fnv.New64()
	return uint64Hasher(algorithm, text)
}

// FNV64a hashes using fnv64a algorithm
func (c *Client) FNV64a(text string) uint64 {
	algorithm := fnv.New64a()
	return uint64Hasher(algorithm, text)
}

// MD5 hashes using md5 algorithm
func (c *Client) MD5(text string) string {
	algorithm := md5.New()
	return stringHasher(algorithm, text)
}

// SHA1 hashes using sha1 algorithm
func (c *Client) SHA1(text string) string {
	algorithm := sha1.New()
	return stringHasher(algorithm, text)
}

// SHA256 hashes using sha256 algorithm
func (c *Client) SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

// SHA512 hashes using sha512 algorithm
func (c *Client) SHA512(text string) string {
	algorithm := sha512.New()
	return stringHasher(algorithm, text)
}

// Hash util
func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func uint32Hasher(algorithm hash.Hash32, text string) uint32 {
	algorithm.Write([]byte(text))
	return algorithm.Sum32()
}

func uint64Hasher(algorithm hash.Hash64, text string) uint64 {
	algorithm.Write([]byte(text))
	return algorithm.Sum64()
}
