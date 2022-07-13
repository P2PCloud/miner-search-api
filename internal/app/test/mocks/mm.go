package mocks

import (
	"testing"
)

// MCtrl wraps test.TB for minimock.
type MCtrl struct {
	T testing.TB
}

// Fatal comply Tester interface.
func (c MCtrl) Fatal(args ...interface{}) {
	c.T.Error(args...)
}

// Fatalf comply Tester interface.
func (c MCtrl) Fatalf(format string, args ...interface{}) {
	c.T.Errorf(format, args...)
}

// Error comply Tester interface.
func (c MCtrl) Error(args ...interface{}) {
	c.T.Error(args...)
}

// Errorf comply Tester interface.
func (c MCtrl) Errorf(format string, args ...interface{}) {
	c.T.Errorf(format, args...)
}

// FailNow comply Tester interface.
func (c MCtrl) FailNow() {}
