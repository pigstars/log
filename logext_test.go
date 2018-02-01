package log

import (
    "os"
	"testing"
)

func TestLogStd(t *testing.T) {

	SetOutputLevel(Ldebug)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")

	SetOutputLevel(Linfo)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")
}

func TestLogNew(t *testing.T){
    Logger := New(os.Stderr,"", Llevel|LstdFlags|Lshortfile, 2)
	Logger.Debugf("Debug: foo\n")
	Logger.Debug("Debug: foo")

	Logger.Infof("Info: foo\n")
	Logger.Info("Info: foo")

	Logger.Warnf("Warn: foo\n")
	Logger.Warn("Warn: foo")

	Logger.Errorf("Error: foo\n")
	Logger.Error("Error: foo")

	Logger.SetOutputLevel(Linfo)

	Logger.Debugf("Debug: foo\n")
	Logger.Debug("Debug: foo")

	Logger.Infof("Info: foo\n")
	Logger.Info("Info: foo")

	Logger.Warnf("Warn: foo\n")
	Logger.Warn("Warn: foo")

	Logger.Errorf("Error: foo\n")
	Logger.Error("Error: foo")
}
