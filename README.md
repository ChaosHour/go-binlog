# go-binlog


## Install 

```Go
go get github.com/ChaosHour/go-binlog

or  

go install github.com/ChaosHour/go-binlog@latest
```


```bash
What was done and why?  

I'm learning Golang and using [Github Copilot](https://copilot.github.com/). 
These two amazing projects gave me the inspiration
to create this project. 

[binlog2sql](https://github.com/danfengcao/binlog2sql) it's written in Python. 

[go-mysql-binlog](https://github.com/liipx/go-mysql-binlog) is a great project, written in Golang.



**Why Golang?**
Because it clicks better for me than anything else at this moment in time.
I like the fact, that anyone can run it without having to mess with dependency hell.


Did I create it? 
Nope, once again [liipx](https://github.com/liipx/go-mysql-binlog) created this amazing repo and I'm just playing arround to see what I can do with it.
```


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
├── cmd
│   ├── LICENSE
│   ├── checksum.go
│   ├── const.go
│   ├── decoder.go
│   ├── event.go
│   ├── rows_event.go
│   └── util.go
├── go.mod
└── main.go

1 directory, 9 files
```

## Note:
I used parts of what was suggested for a main.go file from the authors github repo and modded the rest for my needs.


## Files changed:

```Go     
bash-3.2$ diff -rN  go-mysql-binlog/const.go go-binlog/cmd/const.go
60,95c60,133
< 	UnknownEvent           = 0x00
< 	StartEventV3           = 0x01
< 	QueryEvent             = 0x02
< 	StopEvent              = 0x03
< 	RotateEvent            = 0x04
< 	IntvarEvent            = 0x05
< 	LoadEvent              = 0x06
< 	SlaveEvent             = 0x07
< 	CreateFileEvent        = 0x08
< 	AppendBlockEvent       = 0x09
< 	ExecLoadEvent          = 0x0a
< 	DeleteFileEvent        = 0x0b
< 	NewLoadEvent           = 0x0c
< 	RandEvent              = 0x0d
< 	UserVarEvent           = 0x0e
< 	FormatDescriptionEvent = 0x0f
< 	XIDEvent               = 0x10
< 	BeginLoadQueryEvent    = 0x11
< 	ExecuteLoadQueryEvent  = 0x12
< 	TableMapEvent          = 0x13
< 	WriteRowsEventV0       = 0x14
< 	UpdateRowsEventV0      = 0x15
< 	DeleteRowsEventV0      = 0x16
< 	WriteRowsEventV1       = 0x17
< 	UpdateRowsEventV1      = 0x18
< 	DeleteRowsEventV1      = 0x19
< 	IncidentEvent          = 0x1a
< 	HeartbeatEvent         = 0x1b
< 	IgnorableEvent         = 0x1c
< 	RowsQueryEvent         = 0x1d
< 	WriteRowsEventV2       = 0x1e
< 	UpdateRowsEventV2      = 0x1f
< 	DeleteRowsEventV2      = 0x20
< 	GTIDEvent              = 0x21
< 	AnonymousGTIDEvent     = 0x22
< 	PreviousGTIDEvent      = 0x23
---
> 	UnknownEvent              = 0x00
> 	StartEventV3              = 0x01
> 	QueryEvent                = 0x02
> 	StopEvent                 = 0x03
> 	RotateEvent               = 0x04
> 	IntvarEvent               = 0x05
> 	LoadEvent                 = 0x06
> 	SlaveEvent                = 0x07
> 	CreateFileEvent           = 0x08
> 	AppendBlockEvent          = 0x09
> 	ExecLoadEvent             = 0x0a
> 	DeleteFileEvent           = 0x0b
> 	NewLoadEvent              = 0x0c
> 	RandEvent                 = 0x0d
> 	UserVarEvent              = 0x0e
> 	FormatDescriptionEvent    = 0x0f
> 	XIDEvent                  = 0x10
> 	BeginLoadQueryEvent       = 0x11
> 	ExecuteLoadQueryEvent     = 0x12
> 	TableMapEvent             = 0x13
> 	WriteRowsEventV0          = 0x14
> 	UpdateRowsEventV0         = 0x15
> 	DeleteRowsEventV0         = 0x16
> 	WriteRowsEventV1          = 0x17
> 	UpdateRowsEventV1         = 0x18
> 	DeleteRowsEventV1         = 0x19
> 	IncidentEvent             = 0x1a
> 	HeartbeatEvent            = 0x1b
> 	IgnorableEvent            = 0x1c
> 	RowsQueryEvent            = 0x1d
> 	WriteRowsEventV2          = 0x1e
> 	UpdateRowsEventV2         = 0x1f
> 	DeleteRowsEventV2         = 0x20
> 	GTIDEvent                 = 0x21
> 	AnonymousGTIDEvent        = 0x22
> 	PreviousGTIDEvent         = 0x23
> 	TransactionContextEvent   = 0x27
> 	ViewChangeEvent           = 0x28
> 	XAEvent                   = 0x29
> 	BinlogCheckpointEvent     = 0x2a
> 	BinlogDumpGTIDEvent       = 0x2b
> 	BinlogDumpMGTIDEvent      = 0x2c
> 	BinlogDumpBGTIDEvent      = 0x2d
> 	BinlogTransaction         = 0x2e
> 	RotateEventV2             = 0x2f
> 	IntvarEventV2             = 0x30
> 	RandEventV2               = 0x31
> 	UserVarEventV2            = 0x32
> 	FormatDescriptionEventV2  = 0x33
> 	XIDEventV2                = 0x34
> 	BinlogCheckpointEventV2   = 0x35
> 	IntvarEventV3             = 0x36
> 	HeartbeatEventV2          = 0x37
> 	IgnorableEventV2          = 0x38
> 	RowsQueryEventV2          = 0x39
> 	TableMapEventV3           = 0x3a
> 	WriteRowsEventV3          = 0x3b
> 	UpdateRowsEventV3         = 0x3c
> 	DeleteRowsEventV3         = 0x3d
> 	GtidEventV1               = 0x3e
> 	AnonymousGtidEventV1      = 0x3f
> 	PreviousGtidEventV1       = 0x40
> 	BeginLoadQueryEventV3     = 0x41
> 	ExecuteLoadQueryEventV3   = 0x42
> 	TableMapEventV4           = 0x43
> 	WriteRowsEventV4          = 0x44
> 	UpdateRowsEventV4         = 0x45
> 	DeleteRowsEventV4         = 0x46
> 	WriteRowsEventV5          = 0x47
> 	UpdateRowsEventV5         = 0x48
> 	DeleteRowsEventV5         = 0x49
> 	PartitionedRowsQueryEvent = 0x4a
> 	HeartbeatEventV3          = 0x4b
> 	HeartbeatLogEvent         = 0x4c
100,135c138,196
< 	UnknownEvent:           "UNKNOWN_EVENT",
< 	StartEventV3:           "START_EVENT_V3",
< 	QueryEvent:             "QUERY_EVENT",
< 	StopEvent:              "STOP_EVENT",
< 	RotateEvent:            "ROTATE_EVENT",
< 	IntvarEvent:            "INTVAR_EVENT",
< 	LoadEvent:              "LOAD_EVENT",
< 	SlaveEvent:             "SLAVE_EVENT",
< 	CreateFileEvent:        "CREATE_FILE_EVENT",
< 	AppendBlockEvent:       "APPEND_BLOCK_EVENT",
< 	ExecLoadEvent:          "EXEC_LOAD_EVENT",
< 	DeleteFileEvent:        "DELETE_FILE_EVENT",
< 	NewLoadEvent:           "NEW_LOAD_EVENT",
< 	RandEvent:              "RAND_EVENT",
< 	UserVarEvent:           "USER_VAR_EVENT",
< 	FormatDescriptionEvent: "FORMAT_DESCRIPTION_EVENT",
< 	XIDEvent:               "XID_EVENT",
< 	BeginLoadQueryEvent:    "BEGIN_LOAD_QUERY_EVENT",
< 	ExecuteLoadQueryEvent:  "EXECUTE_LOAD_QUERY_EVENT",
< 	TableMapEvent:          "TABLE_MAP_EVENT",
< 	WriteRowsEventV0:       "WRITE_ROWS_EVENTv0",
< 	UpdateRowsEventV0:      "UPDATE_ROWS_EVENTv0",
< 	DeleteRowsEventV0:      "DELETE_ROWS_EVENTv0",
< 	WriteRowsEventV1:       "WRITE_ROWS_EVENTv1",
< 	UpdateRowsEventV1:      "UPDATE_ROWS_EVENTv1",
< 	DeleteRowsEventV1:      "DELETE_ROWS_EVENTv1",
< 	IncidentEvent:          "INCIDENT_EVENT",
< 	HeartbeatEvent:         "HEARTBEAT_EVENT",
< 	IgnorableEvent:         "IGNORABLE_EVENT",
< 	RowsQueryEvent:         "ROWS_QUERY_EVENT",
< 	WriteRowsEventV2:       "WRITE_ROWS_EVENTv2",
< 	UpdateRowsEventV2:      "UPDATE_ROWS_EVENTv2",
< 	DeleteRowsEventV2:      "DELETE_ROWS_EVENTv2",
< 	GTIDEvent:              "GTID_EVENT",
< 	AnonymousGTIDEvent:     "ANONYMOUS_GTID_EVENT",
< 	PreviousGTIDEvent:      "PREVIOUS_GTIDS_EVENT",
---
> 	UnknownEvent:             "UNKNOWN_EVENT",
> 	StartEventV3:             "START_EVENT_V3",
> 	QueryEvent:               "QUERY_EVENT",
> 	StopEvent:                "STOP_EVENT",
> 	RotateEvent:              "ROTATE_EVENT",
> 	IntvarEvent:              "INTVAR_EVENT",
> 	LoadEvent:                "LOAD_EVENT",
> 	SlaveEvent:               "SLAVE_EVENT",
> 	CreateFileEvent:          "CREATE_FILE_EVENT",
> 	AppendBlockEvent:         "APPEND_BLOCK_EVENT",
> 	ExecLoadEvent:            "EXEC_LOAD_EVENT",
> 	DeleteFileEvent:          "DELETE_FILE_EVENT",
> 	NewLoadEvent:             "NEW_LOAD_EVENT",
> 	RandEvent:                "RAND_EVENT",
> 	UserVarEvent:             "USER_VAR_EVENT",
> 	FormatDescriptionEvent:   "FORMAT_DESCRIPTION_EVENT",
> 	XIDEvent:                 "XID_EVENT",
> 	BeginLoadQueryEvent:      "BEGIN_LOAD_QUERY_EVENT",
> 	ExecuteLoadQueryEvent:    "EXECUTE_LOAD_QUERY_EVENT",
> 	TableMapEvent:            "TABLE_MAP_EVENT",
> 	WriteRowsEventV0:         "WRITE_ROWS_EVENTv0",
> 	UpdateRowsEventV0:        "UPDATE_ROWS_EVENTv0",
> 	DeleteRowsEventV0:        "DELETE_ROWS_EVENTv0",
> 	WriteRowsEventV1:         "WRITE_ROWS_EVENTv1",
> 	UpdateRowsEventV1:        "UPDATE_ROWS_EVENTv1",
> 	DeleteRowsEventV1:        "DELETE_ROWS_EVENTv1",
> 	IncidentEvent:            "INCIDENT_EVENT",
> 	HeartbeatEvent:           "HEARTBEAT_EVENT",
> 	IgnorableEvent:           "IGNORABLE_EVENT",
> 	RowsQueryEvent:           "ROWS_QUERY_EVENT",
> 	WriteRowsEventV2:         "WRITE_ROWS_EVENTv2",
> 	UpdateRowsEventV2:        "UPDATE_ROWS_EVENTv2",
> 	DeleteRowsEventV2:        "DELETE_ROWS_EVENTv2",
> 	GTIDEvent:                "GTID_EVENT",
> 	AnonymousGTIDEvent:       "ANONYMOUS_GTID_EVENT",
> 	PreviousGTIDEvent:        "PREVIOUS_GTIDS_EVENT",
> 	WriteRowsEventV3:         "WRITE_ROWS_EVENTv3",
> 	UpdateRowsEventV3:        "UPDATE_ROWS_EVENTv3",
> 	DeleteRowsEventV3:        "DELETE_ROWS_EVENTv3",
> 	TransactionContextEvent:  "TRANSACTION_CONTEXT_EVENT",
> 	ViewChangeEvent:          "VIEW_CHANGE_EVENT",
> 	XAEvent:                  "XA_EVENT",
> 	BinlogCheckpointEvent:    "BINLOG_CHECKPOINT_EVENT",
> 	BinlogDumpGTIDEvent:      "BINLOG_DUMP_GTID_EVENT",
> 	BinlogDumpMGTIDEvent:     "BINLOG_DUMP_MGTID_EVENT",
> 	BinlogDumpBGTIDEvent:     "BINLOG_DUMP_BGTID_EVENT",
> 	BinlogTransaction:        "BINLOG_TRANSACTION_EVENT",
> 	RotateEventV2:            "ROTATE_EVENTv2",
> 	IntvarEventV2:            "INTVAR_EVENTv2",
> 	RandEventV2:              "RAND_EVENTv2",
> 	UserVarEventV2:           "USER_VAR_EVENTv2",
> 	FormatDescriptionEventV2: "FORMAT_DESCRIPTION_EVENTv2",
> 	XIDEventV2:               "XID_EVENTv2",
> 	BinlogCheckpointEventV2:  "BINLOG_CHECKPOINT_EVENTv2",
> 	IntvarEventV3:            "INTVAR_EVENTv3",
> 	HeartbeatEventV2:         "HEARTBEAT_EVENTv2",
> 	IgnorableEventV2:         "IGNORABLE_EVENTv2",
> 	RowsQueryEventV2:         "ROWS_QUERY_EVENTv2",
> 	TableMapEventV3:          "TABLE_MAP_EVENTv3",



bash-3.2$ diff -rN  go-mysql-binlog/event.go go-binlog/cmd/event.go
175c175,176
< 	desc.MySQLVersion = string(bytes.Trim(data[pos:pos+50], string(0x00)))
---
> 	//desc.MySQLVersion = string(bytes.Trim(data[pos:pos+50], string(0x00)))
> 	desc.MySQLVersion = string(bytes.Trim(data[pos:pos+50], "\x00"))
395a397,424
>
> // Binlog::Rows_query_log_event - Kurt Larsen
> // https://dev.mysql.com/doc/internals/en/rows-query-log-event.html
> // The rows_query_log_event is added to the binlog as last event to tell the reader what query was executed.
> type BinRowsQueryEvent struct {
> 	BaseEventBody
> 	SlaveProxyID  uint32
> 	ExecutionTime uint32
> 	SchemaName    string
> 	Query         string
> }
>
> //	Binlog_version_event
> //	https://dev.mysql.com/doc/internals/en/binlog-version-event.html
> //	The binlog_version_event is added to the binlog as last event to tell the reader what binlog version was used.
> type BinLogVersionEvent struct {
> 	BaseEventBody
> 	Version uint16
> }
>
> //	Binlog_checksum_event
> //	https://dev.mysql.com/doc/internals/en/binlog-checksum-event.html
> //	The binlog_checksum_event is added to the binlog as last event to tell the reader what binlog checksum was used.
> type BinLogChecksumEvent struct {
> 	BaseEventBody
> 	Flags    uint16
> 	Checksum uint32
> }



bash-3.2$ diff -rN  go-mysql-binlog/rows_event.go go-binlog/cmd/rows_event.go
182a183,184
> 	case WriteRowsEventV3, UpdateRowsEventV3, DeleteRowsEventV3:
> 		e.Version = 3
```


## Removed the doc & test dirs and moved files to cmd

```bash
Did Github Copilot do all of the changes?  Yes, it did.  
I knew in my head what I wanted it to do, but because I am not a seasond Go Programmer, Github Copilot was kind to assist.


I'm I taking creadit for this?  
Not at all. I'm just testing and sharring what's been done so far.  More to come.



Testing Env:

bash-3.2$ vagrant --version
Vagrant 2.3.0

bash-3.2$ vboxmanage --version
6.1.36r152435

bash-3.2$ go version
go version go1.19 darwin/amd64


bash-3.2$ ansible --version
ansible [core 2.13.2]
python version = 3.10.6 (main, Aug 11 2022, 13:49:25) [Clang 13.1.6 (clang-1316.0.21.2.5)]
jinja version = 3.1.2
libyaml = True



(updated-pythian-code-challenge) bash-3.2$ vagrant status
Current machine states:

proxysql                    running (virtualbox)
primary                     running (virtualbox)
replica                     running (virtualbox)
etlreplica                  running (virtualbox)


[root@master ~]# mysql --version
mysql  Ver 14.14 Distrib 5.6.51-91.0, for Linux (x86_64) using  6.2

[root@master ~]# uname -msrn
Linux master 3.10.0-1160.76.1.el7.x86_64 x86_64

[root@master ~]# hostnamectl
   Static hostname: master
         Icon name: computer-vm
           Chassis: vm
        Machine ID: ca6cc295a4224d489bc39e5c50f68c35
           Boot ID: 0627bd18e3814cfd8cf48c7b122d783d
    Virtualization: kvm
  Operating System: CentOS Linux 7 (Core)
       CPE OS Name: cpe:/o:centos:centos:7
            Kernel: Linux 3.10.0-1160.76.1.el7.x86_64
      Architecture: x86-64
```