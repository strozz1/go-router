package routes

import (
	"net/http"
)

// struct defining a segment on a route, defining the value and the handler.
// If the segment is not a final path, the handler must be nil.
// This struct represents a piece of a full route.
type Segment struct{
    value string
    handler http.HandlerFunc
    segmentType SegmentType
}

// type of endpoint. It can be static(0) for static route or dynamic(1). 
type SegmentType int 

const (
    StaticType = 0
    DynamicType = 1
)


// create a Segment struct from a string
func FromString(val string) Segment{
    return Segment{
        value: val,
        segmentType: StaticType,
    }
}


// Function implemented from Comparable interface
func (p Segment) Equals(p2 Segment) bool {
    return (p.value == p2.value)
}
func (p Segment) Debug() string {
    return p.value
}

