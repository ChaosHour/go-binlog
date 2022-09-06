# Changed files and directories
[]: # Version: 1.0.0
[]: # Date: 2022-10-05


What was done and why?  

I'm learning Golang and using [Github Copilot](https://copilot.github.com/). 
These two amazing projects gave me the inspiration
to modify this project. 

[binlog2sql](https://github.com/danfengcao/binlog2sql) it's written in Python. 

[go-mysql-binlog](https://github.com/liipx/go-mysql-binlog) is a great project, written in Golang.



**Why Golang?**
Because it clicks better for me than anything else at this moment in time.
I like the fact, that anyone can run it without having to mess with dependency hell.


Did I create it? 
Nope, once again [liipx](https://github.com/liipx/go-mysql-binlog) created this amazing repo and I'm just playing arround to see what I can do with it.


## What files did I change?

```bash
bash-3.2$ diff -rq  go-mysql-binlog go-binlog
Only in go-mysql-binlog: .git
Only in go-mysql-binlog: .gitignore
Only in go-mysql-binlog: LICENSE
Only in go-mysql-binlog: README.md
Only in go-mysql-binlog: checksum.go
Only in go-mysql-binlog: const.go
Only in go-mysql-binlog: decoder.go
Only in go-mysql-binlog: doc
Only in go-mysql-binlog: event.go
Only in go-mysql-binlog: rows_event.go
Only in go-mysql-binlog: test
Only in go-mysql-binlog: util.go

Only in go-binlog: cmd
Only in go-binlog: go.mod
Only in go-binlog: main.go


bash-3.2$ tree go-binlog
go-binlog
├── CHANGED.md
├── LICENSE
├── README.md
├── cmd
│   └── go-binlog
│       ├── LICENSE
│       ├── checksum.go
│       ├── const.go
│       ├── decoder.go
│       ├── event.go
│       ├── rows_event.go
│       └── util.go
├── go.mod
└── main.go

2 directories, 12 files
```

## Note:
I used parts of what was suggested for a main.go file from the authors github repo and modded the rest for my needs.


## Files changed:

    
> for i in $( ls -1 go-binlog/cmd/go-binlog/*); do TB=$( basename ${i}); echo "Current file is ${TB}"; diff -bnrwW 10  go-mysql-binlog/${TB} go-binlog/cmd/go-binlog/${TB}; sleep 0.1 ; done
```Go
Current file is LICENSE
Current file is checksum.go
Current file is const.go
a95 38
	TransactionContextEvent   = 0x27
	ViewChangeEvent           = 0x28
	XAEvent                   = 0x29
	BinlogCheckpointEvent     = 0x2a
	BinlogDumpGTIDEvent       = 0x2b
	BinlogDumpMGTIDEvent      = 0x2c
	BinlogDumpBGTIDEvent      = 0x2d
	BinlogTransaction         = 0x2e
	RotateEventV2             = 0x2f
	IntvarEventV2             = 0x30
	RandEventV2               = 0x31
	UserVarEventV2            = 0x32
	FormatDescriptionEventV2  = 0x33
	XIDEventV2                = 0x34
	BinlogCheckpointEventV2   = 0x35
	IntvarEventV3             = 0x36
	HeartbeatEventV2          = 0x37
	IgnorableEventV2          = 0x38
	RowsQueryEventV2          = 0x39
	TableMapEventV3           = 0x3a
	WriteRowsEventV3          = 0x3b
	UpdateRowsEventV3         = 0x3c
	DeleteRowsEventV3         = 0x3d
	GtidEventV1               = 0x3e
	AnonymousGtidEventV1      = 0x3f
	PreviousGtidEventV1       = 0x40
	BeginLoadQueryEventV3     = 0x41
	ExecuteLoadQueryEventV3   = 0x42
	TableMapEventV4           = 0x43
	WriteRowsEventV4          = 0x44
	UpdateRowsEventV4         = 0x45
	DeleteRowsEventV4         = 0x46
	WriteRowsEventV5          = 0x47
	UpdateRowsEventV5         = 0x48
	DeleteRowsEventV5         = 0x49
	PartitionedRowsQueryEvent = 0x4a
	HeartbeatEventV3          = 0x4b
	HeartbeatLogEvent         = 0x4c
a135 23
	WriteRowsEventV3:         "WRITE_ROWS_EVENTv3",
	UpdateRowsEventV3:        "UPDATE_ROWS_EVENTv3",
	DeleteRowsEventV3:        "DELETE_ROWS_EVENTv3",
	TransactionContextEvent:  "TRANSACTION_CONTEXT_EVENT",
	ViewChangeEvent:          "VIEW_CHANGE_EVENT",
	XAEvent:                  "XA_EVENT",
	BinlogCheckpointEvent:    "BINLOG_CHECKPOINT_EVENT",
	BinlogDumpGTIDEvent:      "BINLOG_DUMP_GTID_EVENT",
	BinlogDumpMGTIDEvent:     "BINLOG_DUMP_MGTID_EVENT",
	BinlogDumpBGTIDEvent:     "BINLOG_DUMP_BGTID_EVENT",
	BinlogTransaction:        "BINLOG_TRANSACTION_EVENT",
	RotateEventV2:            "ROTATE_EVENTv2",
	IntvarEventV2:            "INTVAR_EVENTv2",
	RandEventV2:              "RAND_EVENTv2",
	UserVarEventV2:           "USER_VAR_EVENTv2",
	FormatDescriptionEventV2: "FORMAT_DESCRIPTION_EVENTv2",
	XIDEventV2:               "XID_EVENTv2",
	BinlogCheckpointEventV2:  "BINLOG_CHECKPOINT_EVENTv2",
	IntvarEventV3:            "INTVAR_EVENTv3",
	HeartbeatEventV2:         "HEARTBEAT_EVENTv2",
	IgnorableEventV2:         "IGNORABLE_EVENTv2",
	RowsQueryEventV2:         "ROWS_QUERY_EVENTv2",
	TableMapEventV3:          "TABLE_MAP_EVENTv3",
Current file is decoder.go
Current file is event.go
d175 1
a175 2
	//desc.MySQLVersion = string(bytes.Trim(data[pos:pos+50], string(0x00)))
	desc.MySQLVersion = string(bytes.Trim(data[pos:pos+50], "\x00"))
a395 28

// Binlog::Rows_query_log_event - Kurt Larsen
// https://dev.mysql.com/doc/internals/en/rows-query-log-event.html
// The rows_query_log_event is added to the binlog as last event to tell the reader what query was executed.
type BinRowsQueryEvent struct {
	BaseEventBody
	SlaveProxyID  uint32
	ExecutionTime uint32
	SchemaName    string
	Query         string
}

//	Binlog_version_event
//	https://dev.mysql.com/doc/internals/en/binlog-version-event.html
//	The binlog_version_event is added to the binlog as last event to tell the reader what binlog version was used.
type BinLogVersionEvent struct {
	BaseEventBody
	Version uint16
}

//	Binlog_checksum_event
//	https://dev.mysql.com/doc/internals/en/binlog-checksum-event.html
//	The binlog_checksum_event is added to the binlog as last event to tell the reader what binlog checksum was used.
type BinLogChecksumEvent struct {
	BaseEventBody
	Flags    uint16
	Checksum uint32
}
Current file is rows_event.go
a182 2
	case WriteRowsEventV3, UpdateRowsEventV3, DeleteRowsEventV3:
		e.Version = 3
Current file is util.go
```


## Removed the doc & test dirs and moved files to cmd


Did Github Copilot do all of the changes?  Yes, it did.  
I knew in my head what I wanted it to do, but because I am not a seasond Go Programmer, Github Copilot was kind to assist.
