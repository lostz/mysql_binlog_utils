package mysql_binlog_util

const (
	UNKNOWN_EVENT            = 0
	START_EVENT_V3           = 1
	QUERY_EVENT              = 2
	STOP_EVENT               = 3
	ROTATE_EVENT             = 4
	INTVAR_EVENT             = 5
	LOAD_EVENT               = 6
	SLAVE_EVENT              = 7
	CREATE_FILE_EVENT        = 8
	APPEND_BLOCK_EVENT       = 9
	EXEC_LOAD_EVENT          = 10
	DELETE_FILE_EVENT        = 11
	NEW_LOAD_EVENT           = 12
	RAND_EVENT               = 13
	USER_VAR_EVENT           = 14
	FORMAT_DESCRIPTION_EVENT = 15
	XID_EVENT                = 16
	BEGIN_LOAD_QUERY_EVENT   = 17
	EXECUTE_LOAD_QUERY_EVENT = 18
	TABLE_MAP_EVENT          = 19
	PRE_GA_WRITE_ROWS_EVENT  = 20
	PRE_GA_UPDATE_ROWS_EVENT = 21
	PRE_GA_DELETE_ROWS_EVENT = 22
	WRITE_ROWS_EVENT_V1      = 23
	UPDATE_ROWS_EVENT_V1     = 24
	DELETE_ROWS_EVENT_V1     = 25
	INCIDENT_EVENT           = 26
	HEARTBEAT_LOG_EVENT      = 27
	IGNORABLE_LOG_EVENT      = 28
	ROWS_QUERY_LOG_EVENT     = 29
	WRITE_ROWS_EVENT         = 30
	UPDATE_ROWS_EVENT        = 31
	DELETE_ROWS_EVENT        = 32
	GTID_LOG_EVENT           = 33
	ANONYMOUS_GTID_LOG_EVENT = 34
	PREVIOUS_GTIDS_LOG_EVENT = 35
)

const (
	LOG_EVENT_FIXED_HEADER_LEN = 19
	MAX_ALLOWED_PACKET         = 1024 * 1024 * 1024
)

type EventFixedHeader struct {
	Timestamp    int
	EventType    int
	ServerId     int
	EventLength  int
	NextPosition int
	Flags        int
}

type EventFixedData struct {
	Bytes []byte
}

type EventVariableData struct {
	Bytes []byte
}
