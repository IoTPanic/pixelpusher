package db

var createDeviceTableStatement = `CREATE TABLE IF NOT EXISTS fixtures (
	id INTEGER PRIMARY KEY,
	name VARCHAR(255),
	longID VARCHAR(100),
	hostname VARCHAR(100),
	port INTEGER,
	connector BOOLEAN,
	key VARCHAR(255),
	use_key BOOLEAN,
	channel_cnt INTEGER
)`

var createChannelsTableStatement = `CREATE TABLE IF NOT EXISTS channels (
	id INTEGER PRIMARY KEY,
	name VARCHAR(100),
	device INTEGER,
	type INTEGER,
	color VARCHAR(100),
	max_length INTEGER
)
`

var createMatrixexTableStatment = `CREATE TABLE IF NOT EXISTS matrixes (
	id INTEGER PRIMARY KEY,
	device INTEGER,
	channel INTEGER,
	width INTEGER,
	height INTEGER,
	type INTEGER,
	coloring VARCHAR(100),
	offset INTEGER,
	brightness INTEGER
)
`

var deleteDevicesTablesStatment = `DROP TABLE devices`
