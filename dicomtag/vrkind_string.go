// Code generated by "stringer -type VRKind"; DO NOT EDIT

package dicomtag

import "fmt"

//const _VRKind_name = "VRStringListVRBytesVRStringVRUInt16ListVRUInt32ListVRInt16ListVRInt32ListVRFloat32ListVRFloat64ListVRSequenceVRItemVRTagListVRDateVRPixelData"

var vrKindNames = map[VRKind]string{
	VRStringList:  "VRStringList",
	VRBytes:       "VRBytes",
	VRString:      "VRString",
	VRUInt16List:  "VRUInt16List",
	VRUInt32List:  "VRUInt32List",
	VRInt16List:   "VRInt16List",
	VRInt32List:   "VRInt32List",
	VRFloat32List: "VRFloat32List",
	VRFloat64List: "VRFloat64List",
	VRSequence:    "VRSequence",
	VRItem:        "VRItem",
	VRTagList:     "VRTagList",
	VRDate:        "VRDate",
	VRDateTime:    "VRDateTime",
	VRTime:        "VRTime",
	VRPixelData:   "VRPixelData",
}

//var _VRKind_index = [...]uint8{0, 12, 19, 27, 39, 51, 62, 73, 86, 99, 109, 115, 124, 130, 141}

func (i VRKind) String() string {
	if result, ok := vrKindNames[i]; ok {
		return result
	}
	return fmt.Sprintf("VRKind(%d)", i)
}
