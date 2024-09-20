package db

import "github.com/google/wire"

// ProviderSetDB is a Wire provider set that provides *sql.DB using InitializeDB.
var ProviderSetDB = wire.NewSet(InitializeDB)
