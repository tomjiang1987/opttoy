// Code generated by "stringer -type=operator operator.go"; DO NOT EDIT.

package main

import "fmt"

const _operator_name = "unknownOprootOpdefineListOpdefineOpdefineFieldOpruleListOpruleHeaderOpruleOpbindOprefOpmatchInvokeOpmatchFieldsOpmatchAndOpmatchNotOpmatchAnyOpreplaceListOpconstructOptagsOpstringOp"

var _operator_index = [...]uint8{0, 9, 15, 27, 35, 48, 58, 70, 76, 82, 87, 100, 113, 123, 133, 143, 156, 167, 173, 181}

func (i operator) String() string {
	if i < 0 || i >= operator(len(_operator_index)-1) {
		return fmt.Sprintf("operator(%d)", i)
	}
	return _operator_name[_operator_index[i]:_operator_index[i+1]]
}
