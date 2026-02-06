package sqlite3

/*
// make go-sqlite3 use embedded library without code changes
#cgo CFLAGS: -DUSE_LIBSQLITE3

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// use platform-native crypto: CommonCrypto on macOS, OpenSSL on Linux/Windows
#cgo darwin CFLAGS: -DSQLCIPHER_CRYPTO_CC
#cgo darwin LDFLAGS: -framework Security
#cgo linux CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo linux LDFLAGS: -lcrypto
#cgo windows CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo windows LDFLAGS: -lcrypto

// sqlcipher extra init/shutdown hooks (required by SQLCipher 4.6+)
#cgo CFLAGS: -DSQLITE_EXTRA_INIT=sqlcipher_extra_init
#cgo CFLAGS: -DSQLITE_EXTRA_SHUTDOWN=sqlcipher_extra_shutdown

// enable FTS5 full-text search
#cgo CFLAGS: -DSQLITE_ENABLE_FTS5

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1
*/
import "C"
