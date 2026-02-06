## go-sqlcipher (fork)

Fork of [mutecomm/go-sqlcipher](https://github.com/mutecomm/go-sqlcipher) with the SQLCipher amalgamation updated to **v4.13.0** (SQLite 3.51.2).

### Why this fork?

The upstream package bundles SQLCipher v4.4.2 (SQLite 3.33.0), which lacks the `sqlite3_vtab_in` APIs introduced in SQLite 3.38.0. These APIs are required by [sqlite-vec](https://github.com/asg017/sqlite-vec) for vector similarity search. This fork updates the amalgamation so that SQLCipher and sqlite-vec can coexist in a single encrypted database.

### Changes from upstream

- **SQLCipher v4.13.0** amalgamation (`sqlite3.c`, `sqlite3.h`, `sqlite3ext.h`) based on SQLite 3.51.2
- **Platform-native crypto** — CommonCrypto on macOS, OpenSSL on Linux/Windows (libtomcrypt was removed in SQLCipher v4.6+)
- **FTS5** full-text search module enabled
- **`SQLITE_EXTRA_INIT`/`SQLITE_EXTRA_SHUTDOWN`** flags added (required by SQLCipher v4.6+)

### Verified compatible with

| Component | Version |
|-----------|---------|
| SQLCipher | 4.13.0 community |
| sqlite-vec | v0.1.6 |
| FTS5 | built-in |

All three work together in a single encrypted database. All upstream go-sqlcipher tests pass.

### Installation

```
go get github.com/miclip/go-sqlcipher/v4
```

Or use a `replace` directive if your code imports the upstream module:

```
replace github.com/mutecomm/go-sqlcipher/v4 => github.com/miclip/go-sqlcipher/v4 <version>
```

### Build requirements

Requires `CGO_ENABLED=1` and platform crypto libraries:

- **macOS**: Xcode command line tools (CommonCrypto + Security framework, included by default)
- **Linux**: `libssl-dev` / `openssl-devel`
- **Windows**: OpenSSL development headers

### Documentation

To create and open encrypted database files use the following DSN parameters:

```go
key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
dbname := fmt.Sprintf("db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
db, _ := sql.Open("sqlite3", dbname)
```

`_pragma_key` is the hex encoded 32 byte key (must be 64 characters long).
`_pragma_cipher_page_size` is the page size of the encrypted database (set if
you want a different value than the default size).

```go
key := url.QueryEscape("secret")
dbname := fmt.Sprintf("db?_pragma_key=%s&_pragma_cipher_page_size=4096", key)
db, _ := sql.Open("sqlite3", dbname)
```

This uses a passphrase directly as `_pragma_key` with the key derivation function in
SQLCipher. Do not forget the `url.QueryEscape()` call in your code!

See also [PRAGMA key](https://www.zetetic.net/sqlcipher/sqlcipher-api/#PRAGMA_key).

Use the function
[sqlite3.IsEncrypted()](https://godoc.org/github.com/mutecomm/go-sqlcipher#IsEncrypted)
to check whether a database file is encrypted or not.

Examples can be found under the `./_example` directory.

### License

The code of the originating packages is covered by their respective licenses.
See [LICENSE](LICENSE) file for details.
