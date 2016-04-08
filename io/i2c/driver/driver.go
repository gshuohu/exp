// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driver contains interfaces to be implemented by various I2C implementations.
package driver // import "golang.org/x/exp/io/i2c/driver"

// Opener is an interface to be implemented by the I2C driver to open
// a connection to an I2C device with the specified bus number and I2C address.
type Opener interface {
	Open(bus, addr int) (Conn, error)
}

// Conn represents an active connection to an I2C device.
type Conn interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close() error
}