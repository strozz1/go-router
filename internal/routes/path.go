package routes

import (
	"net/http"
)

// type def for a path content, defining the value and the handler.
// If the path is not a final path, the handler must be nil
type PathContent struct{
    value string
    handler http.HandlerFunc
}

func FromString(val string) PathContent{
    return PathContent{
        value: val,
    }
}



func (p PathContent) Equals(p2 PathContent) bool {
    return (p.value == p2.value)
}



type Path struct{
    routes Tree[PathContent]

} 


func EmptyPath() Path{
    return Path{
        routes: EmptyTree[PathContent](),
    }
}

func (p *Path) AddPath(routes []string, handler http.HandlerFunc) *PathContent{

    current := p.routes.root
    var tmp PathContent
    for _,n := range routes{
        tmp = FromString(n)
        find := current.GetChild(tmp)

        //if nil create new node and append to current
        if(find == nil){
            find = &Node[PathContent]{
                value: tmp,
                children : []*Node[PathContent]{},
            }
            current.children = append(current.children, find)
        }

        current = find
    }
    content := &current.value
    content.handler=handler
    return content
}

func (p *Path) Search(routes []string) *PathContent{
    current := p.routes.root
    var tmp PathContent
    for _,n := range routes{
        tmp = FromString(n)
        find := current.GetChild(tmp)

        

        //if nil create new node and append to current
        if(find == nil){
            return nil
        }
        current = find
    }
    return &current.value
 

}
