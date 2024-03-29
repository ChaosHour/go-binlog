# go-binlog

> It's not perfect and more work needs to be done. But it's a start.


## Install 

```Go

mkdir -p  go/{src,bin,pkg}  

go install github.com/ChaosHour/go-binlog@latest


[root@primary ~]# go install github.com/ChaosHour/go-binlog@latest
go: downloading github.com/ChaosHour/go-binlog v0.0.0-20220906163523-b9774708b3e6

[root@primary ~]# ls -lrt go/bin
total 2056
-rwxr-xr-x. 1 root root 2104276 Sep  6 18:44 go-binlog

```

> I'm using these in my my.cnf file:

```ini
binlog_format = ROW
binlog_row_image = FULL
```

## Usage

```Go
go-binlog -h
Please specify a MySQL binary log file
  -f string
    	MySQL binary log file
  -h	Print help

 go-binlog -f /db/logs01/mysql-bin.000003
 
 [root@primary go-binlog]# go-binlog -f /db/logs01/mysql-bin.000003
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-05 02:09:12 +0000 UTC, ServerID:1662343752, EventSize:269, EventEndPos:2680, Flag:0x0
&{{} 3929 0 0 47 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 33 0 33 0 33 0 12 1 103 114 111 117 112 111 110 95 112 114 111 100 117 99 116 105 111 110 0] groupon_production CREATE TABLE IF NOT EXISTS dictionary (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  word varchar(100) NOT NULL,
  mean varchar(300) NOT NULL,
  PRIMARY KEY (id)
)}

go-binlog -f /db/logs01/mysql-bin.000003
Got FORMAT_DESCRIPTION_EVENT:
	Type:FORMAT_DESCRIPTION_EVENT, Time:2022-09-04 23:58:55 +0000 UTC, ServerID:1662335935, EventSize:116, EventEndPos:120, Flag:0x1
&{{} 4 5.6.51-91.0-log 1662335935 19 [56 13 0 8 0 18 0 4 4 4 4 18 0 0 92 0 4 26 8 0 0 0 8 8 8 2 0 0 0 10 10 10 25 25 0 1 30 104 219 224] true}
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-04 23:58:58 +0000 UTC, ServerID:1662335938, EventSize:142, EventEndPos:262, Flag:0x0
&{{} 4 0 0 26 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0]  SET PASSWORD FOR 'root'@'localhost'='*CC44899BBE450A06A0823407493390266377825C'}
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-04 23:58:58 +0000 UTC, ServerID:1662335938, EventSize:172, EventEndPos:434, Flag:0x0
&{{} 5 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'proxy'@'192.168.50.%' IDENTIFIED BY PASSWORD '*D7CE0B8C570E52CB6BA166CD39364EF036ED7975'}
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-04 23:58:58 +0000 UTC, ServerID:1662335938, EventSize:131, EventEndPos:565, Flag:0x0
&{{} 5 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT USAGE ON *.* TO 'proxy'@'192.168.50.%'}
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-04 23:58:59 +0000 UTC, ServerID:1662335939, EventSize:171, EventEndPos:736, Flag:0x0
&{{} 6 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'repl'@'192.168.50.%' IDENTIFIED BY PASSWORD '*79536349C907C75D6CB57238E75EB257B429A751'}
Got QUERY_EVENT:
	Type:QUERY_EVENT, Time:2022-09-04 23:58:59 +0000 UTC, ServerID:1662335939, EventSize:142, EventEndPos:878, Flag:0x0
&{{} 6 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT REPLICATION SLAVE ON *.* TO 'repl'@'192.168.50.%'}
```

## Testing New features

You can now select a start and end position to read from the binary log file.

```Go
go-binlog on  new-features via 🐹 v1.20.5 
❯ go run main.go -h                                  
Please specify a MySQL binary log file
  -d string
        MySQL database name
  -e int
        End position
  -f string
        MySQL binary log file
  -h    Print help
  -s int
        Start position
  -t string
        MySQL table name

```

```Go
go-binlog on  new-features [!?] via 🐹 v1.20.5 
❯ go run main.go -s 1988 -e 2152 -f mysql-bin.000003
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:161, EventEndPos:1988, Flag:0x0
LogPos: 1988
&{{} 11 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT SELECT,PROCESS,REPLICATION CLIENT ON *.* TO 'monuser'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:164, EventEndPos:2152, Flag:0x0
LogPos: 2152
&{{} 12 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'dba_util'@'%' IDENTIFIED BY PASSWORD '*EC6A383BDA085D07ECA7FE38DEA96477ECA56C76'}
Read 2 events
End position of binlog: 2273

go-binlog on  new-features [!?] via 🐹 v1.20.5 
❯ go run main.go -f mysql-bin.000003                
Got FORMAT_DESCRIPTION_EVENT: 
        Type:FORMAT_DESCRIPTION_EVENT, Time:2022-09-05 16:46:55 -0700 PDT, ServerID:1662421615, EventSize:116, EventEndPos:120, Flag:0x1
LogPos: 120
&{{} 4 5.6.51-91.0-log 1662421615 19 [56 13 0 8 0 18 0 4 4 4 4 18 0 0 92 0 4 26 8 0 0 0 8 8 8 2 0 0 0 10 10 10 25 25 0 1 96 100 190 152] true}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:58 -0700 PDT, ServerID:1662421618, EventSize:142, EventEndPos:262, Flag:0x0
LogPos: 262
&{{} 5 0 0 26 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0]  SET PASSWORD FOR 'root'@'localhost'='*CC44899BBE450A06A0823407493390266377825C'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:58 -0700 PDT, ServerID:1662421618, EventSize:172, EventEndPos:434, Flag:0x0
LogPos: 434
&{{} 6 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'proxy'@'192.168.50.%' IDENTIFIED BY PASSWORD '*D7CE0B8C570E52CB6BA166CD39364EF036ED7975'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:58 -0700 PDT, ServerID:1662421618, EventSize:131, EventEndPos:565, Flag:0x0
LogPos: 565
&{{} 6 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT USAGE ON *.* TO 'proxy'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:171, EventEndPos:736, Flag:0x0
LogPos: 736
&{{} 7 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'repl'@'192.168.50.%' IDENTIFIED BY PASSWORD '*79536349C907C75D6CB57238E75EB257B429A751'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:142, EventEndPos:878, Flag:0x0
LogPos: 878
&{{} 7 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT REPLICATION SLAVE ON *.* TO 'repl'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:172, EventEndPos:1050, Flag:0x0
LogPos: 1050
&{{} 8 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'chaos'@'192.168.50.%' IDENTIFIED BY PASSWORD '*D37C49F9CBEFBF8B6F4B165AC703AA271E079004'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:131, EventEndPos:1181, Flag:0x0
LogPos: 1181
&{{} 8 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT USAGE ON *.* TO 'chaos'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:139, EventEndPos:1320, Flag:0x0
LogPos: 1320
&{{} 8 0 0 55 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 2 116 101 115 116 0 109 121 115 113 108 0]  GRANT ALL ON `test`.* TO 'chaos'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:171, EventEndPos:1491, Flag:0x0
LogPos: 1491
&{{} 10 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'sstuser'@'localhost' IDENTIFIED BY PASSWORD '*B865CAE8F340F6CE1485A06F4492BB49718DF1EC'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:46:59 -0700 PDT, ServerID:1662421619, EventSize:162, EventEndPos:1653, Flag:0x0
LogPos: 1653
&{{} 10 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT REPLICATION CLIENT,RELOAD,LOCK TABLES ON *.* TO 'sstuser'@'localhost'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:174, EventEndPos:1827, Flag:0x0
LogPos: 1827
&{{} 11 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'monuser'@'192.168.50.%' IDENTIFIED BY PASSWORD '*EC6A383BDA085D07ECA7FE38DEA96477ECA56C76'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:161, EventEndPos:1988, Flag:0x0
LogPos: 1988
&{{} 11 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT SELECT,PROCESS,REPLICATION CLIENT ON *.* TO 'monuser'@'192.168.50.%'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:164, EventEndPos:2152, Flag:0x0
LogPos: 2152
&{{} 12 0 0 34 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 12 1 109 121 115 113 108 0]  CREATE USER 'dba_util'@'%' IDENTIFIED BY PASSWORD '*EC6A383BDA085D07ECA7FE38DEA96477ECA56C76'}
Got QUERY_EVENT: 
        Type:QUERY_EVENT, Time:2022-09-05 16:47:00 -0700 PDT, ServerID:1662421620, EventSize:121, EventEndPos:2273, Flag:0x0
LogPos: 2273
&{{} 12 0 0 50 [0 0 0 0 0 1 0 0 0 16 0 0 0 0 6 3 115 116 100 4 8 0 8 0 33 0 11 4 114 111 111 116 9 108 111 99 97 108 104 111 115 116 12 1 109 121 115 113 108 0]  GRANT ALL ON *.* TO 'dba_util'@'%'}
Read 15 events
End position of binlog: 2273
```


```bash
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


[root@primary ~]# mysql --version
mysql  Ver 14.14 Distrib 5.6.51-91.0, for Linux (x86_64) using  6.2

[root@primary ~]# uname -msrn
Linux master 3.10.0-1160.76.1.el7.x86_64 x86_64

[root@primary ~]# hostnamectl
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