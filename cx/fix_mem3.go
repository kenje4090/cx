package cxcore

import (
	"log"
)
//NOTE: Temp file for resolving GetFinalOffset issue
//TODO: What should this function be called?

// GetFinalOffset ...
// Get byte position to write operand result back to?
// Note: CXArgument is bloated, should pass in required though something else

//NEEDS COMMENT. WTF DOES THIS DO?
//TODO:
//GetFinalOffset
//->
//GetFinalOffsetI32
//GetFinalOffsetF32
//GetfinalOffsetI16
//ETC
func GetFinalOffset(fp int, arg *CXArgument) int {
	// defer RuntimeError(PROGRAM)
	// var elt *CXArgument
	finalOffset := arg.Offset
	// var fldIdx int

	// elt = arg

	//Todo: find way to eliminate this check
	if finalOffset < PROGRAM.StackSize {
		// Then it's in the stack, not in data or heap and we need to consider the frame pointer.
		finalOffset += fp
	}

	// elt = arg
	//TODO: Eliminate all op codes with more than one return type
	//TODO: Eliminate this loop
	//Q: How can CalculateDereferences change offset?
	//Why is finalOffset fed in as a pointer?
	CalculateDereferences(arg, &finalOffset, fp)
	for _, fld := range arg.Fields {
		// elt = fld
		finalOffset += fld.Offset
		CalculateDereferences(fld, &finalOffset, fp)
	}

	return finalOffset
}


//OMFG. set ENABLE_MIRACLE_BUG to true and do `make build; make test`
//var ENABLE_MIRACLE_BUG bool = true //uses GetFinalOffset for everything
var ENHANCED_DEBUGING bool = true //runs asserts to find error

//this is simplest version of function that works for atomic types
func GetOffsetAtomicSimple(fp int, arg *CXArgument) int {
	finalOffset := arg.Offset
	if finalOffset < PROGRAM.StackSize {
		finalOffset += fp //check if on stack
	}
	return finalOffset
}

var ENABLE_MIRACLE_BUG bool    = false
//this is version with type assertions
func GetOffsetAtomic(fp int, arg *CXArgument) int {
	if ENABLE_MIRACLE_BUG == false {
		return GetFinalOffset(fp, arg)
	}

	finalOffset := arg.Offset
	//Todo: find way to eliminate this check
	if finalOffset < PROGRAM.StackSize {
		// Then it's in the stack, not in data or heap and we need to consider the frame pointer.
		finalOffset += fp
	}

	if ENHANCED_DEBUGING {
		offset1 := finalOffset //save value
		CalculateDereferences(arg, &offset1, fp)
		if offset1 != finalOffset {
			log.Panicf("fix_mem3.go, GetOffsetAtomic(), offfset1 != finalOffset, offset1= %d, finalOffset= %d \n", offset1, finalOffset)
		}
		if len(arg.Fields) != 0 {
			log.Panic("fix_mem4.go, GetOffsetAtomic(): arg.Fields cannot be greater than 0 for atomic types\n")
		}
	}
	return finalOffset
}

// GetOffset_i8 ...
func GetOffset_i8(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp,arg)
}

// GetOffset_i16 ...
func GetOffset_i16(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_i32 ...
func GetOffset_i32(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_i64 ...
func GetOffset_i64(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_ui8 ...
func GetOffset_ui8(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_ui16 ...
func GetOffset_ui16(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_ui32 ...
func GetOffset_ui32(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_ui64 ...
func GetOffset_ui64(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_f32 ...
func GetOffset_f32(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_f64 ...
func GetOffset_f64(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_bool ...
//NOTE: BOOL is not ready for migration yet
func GetOffset_bool(fp int, arg *CXArgument) int {
	//return GetFinalOffset(fp, arg)
	return GetOffsetAtomic(fp, arg)
}

// GetOffset_str ...
func GetOffset_str(fp int, arg *CXArgument) int {
	return GetFinalOffset(fp, arg)
}
