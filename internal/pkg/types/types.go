package types

const (
	MD5     = "md5"
	SHA1    = "sha1"
	SHA256  = "sha256"
	SHA3    = "sha3"
	SHA3224 = "sha3-224"
	SHA3256 = "sha3-256"
	SHA3384 = "sha3-384"
	SHA3512 = "sha3-512"

	FNV1  = "fnv-1"
	FNV1A = "fnv-1a"
	XXH64 = "xxh64"
	XXH3  = "xxh3"

	SDHASH = "sdhash"
	SSDEEP = "ssdeep"
	TLSH   = "tlsh"

	CRC32IEEE = "crc32-ieee"
	CRC64ECMA = "crc64-ecma"
	CRC64ISO  = "crc64-iso"
)

type Heap int

const (
	Regular Heap = iota
	Deflate
	Ignore
	Stdin
	Stdout
	Stderr
	Plugin
	Chat
)

type Invoke int

const (
	None Invoke = iota
	Compare
	Counts
	Entropy
	HashSum
	Strings
	Timeline
	Unique
)

type Null struct{}
