package process

import (
	"supervisor-go/config"
	"testing"
)

var procs *Manager = NewManager()

func TestProcessMgrAdd(t *testing.T) {
	entry := &config.Entry{ConfigDir: ".", Group: "test", Name: "program:test1"}
	procs.Clear()
	procs.Add("test1", NewProcess("supervisor-go", entry))

	if procs.Find("test1") == nil {
		t.Error("fail to add process")
	}
}

func TestProcMgrRemove(t *testing.T) {
	procs.Clear()
	procs.Add("test1", &Process{})
	proc := procs.Remove("test1")

	if proc == nil {
		t.Error("fail to remove process")
	}

	proc = procs.Remove("test1")
	if proc != nil {
		t.Error("fail to remove process")
	}
}
