// generated by stringer -type=FrameSize; DO NOT EDIT

package ast

import "fmt"

const _FrameSize_name = "SzFixedSzVar8SzVar16"

var _FrameSize_index = [...]uint8{0, 7, 13, 20}

func (i FrameSize) String() string {
	if i < 0 || i >= FrameSize(len(_FrameSize_index)-1) {
		return fmt.Sprintf("FrameSize(%d)", i)
	}
	return _FrameSize_name[_FrameSize_index[i]:_FrameSize_index[i+1]]
}
