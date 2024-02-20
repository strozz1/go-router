package routes

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

// struct defining a segment on a route, defining the value and the handler.
// If the segment is not a final path, the handler must be nil.
// This struct represents a piece of a full route.
type Segment struct {
	value       string
	handler     http.HandlerFunc
	segmentType SegmentType
}

// type of endpoint. It can be static(0) for static route or dynamic(1).
type SegmentType int

const (
	StaticType  = 0
	DynamicType = 1
)

// create a Segment struct from a string
func FromString(val string) (Segment,error) {
    var segment Segment
	segType := StaticType
	if IsDinamic(val) {
		segType = DynamicType
        if(!ValidateDinRoute(val)){
            return segment,errors.New("dinamic route format is invalid!: "+val)
        }
	}else{
        if(!ValidateRoute(val)){
            return segment,errors.New("route format is invalid: "+val)
        }
    }

    segment =Segment{
        value: val,
		segmentType: SegmentType(segType),
    }

   	return segment,nil
}

// Function implemented from Comparable interface
func (p Segment) Equals(p2 Segment) bool {
	if p.segmentType == StaticType {
		return (p.value == p2.value)
	}
	return true

}
func (p Segment) Debug() string {
        write := "" + p.value + ", " + fmt.Sprintf("%d",p.segmentType) + "."
	return write
}

func IsDinamic(value string) bool {
    // regex: matches if string contains aZ and/or - and more aZ
	r, _ := regexp.MatchString("\\/{.*}", value)
	return r
}


func ValidateRoute(val string) bool{
	r, _ := regexp.MatchString("(\\/[a-zA-Z]+(-?[a-zA-Z]+)*)$|\\/$", val)
    return r
}

func ValidateDinRoute(val string) bool{
	r, _ := regexp.MatchString("{[a-zA-Z]+(-?[a-zA-Z]+)*}", val)
    return r
}
