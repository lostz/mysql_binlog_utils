package mysql_binlog_util

import (
	"os"
	"os/exec"
	"testing"
)

func TestDumpBinlogFromPos0(t *testing.T) {
	defer os.Remove("./test/test-mysql-bin-dump")
	if err := DumpBinlogFromPos("./test/test-mysql-bin", 107, "./test/test-mysql-bin-dump"); nil != err {
		t.Errorf("expect no err, but got %v", err)
	}
	if err := exec.Command("sh", "-c", "./test/mysqlbinlog ./test/test-mysql-bin-dump > /dev/null").Run(); nil != err {
		t.Errorf("expect dump log could be parsed by mysqlbinlog, but failed with err=%v", err)
	}
}

func TestDumpBinlogFromPos1(t *testing.T) {
	defer os.Remove("./test/test-mysql-bin-dump")
	if err := DumpBinlogFromPos("./test/test-mysql-bin", 24959, "./test/test-mysql-bin-dump"); nil != err {
		t.Errorf("expect no err, but got %v", err)
	}
	if err := exec.Command("sh", "-c", "./test/mysqlbinlog ./test/test-mysql-bin-dump > /dev/null").Run(); nil != err {
		t.Errorf("expect dump log could be parsed by mysqlbinlog, but failed with err=%v", err)
	}
}
