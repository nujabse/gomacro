/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * api.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package common

type Size uint8 // 1, 2, 4 or 8

// symbolic assembly code: instruction or its arguments
type AsmCode interface {
	asmcode()
}

// assembled machine code.
// Executable if compiled for the same architecture
// the program is running on - see Asm.Func()
type MachineCode struct {
	ArchId ArchId
	Bytes  []uint8
}

// argument of assembly instructions
type Arg interface {
	RegId() RegId // register used by Arg, or NoReg if Arg is Const
	Kind() Kind
	Const() bool
	asmcode()
}

// memory area where spill registers can be saved
type Save struct {
	reg              Reg      // points to memory area
	start, next, end SaveSlot // memory area indexes
	bitmap           []bool   // bitmap of used/free indexes
}

func SizeOf(a Arg) Size {
	size := a.Kind().Size()
	if size == 0 {
		errorf("unsupported register/memory kind: %v", a.Kind())
	}
	return size
}
