module go-blockchain

go 1.13

require (
	github.com/boltdb/bolt v1.3.1
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
)

replace (
	github.com/boltdb/bolt v1.3.1 => ./vendor/github.com/boltdb/bolt
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad => ./vendor/golang.org/x/crypto
)
