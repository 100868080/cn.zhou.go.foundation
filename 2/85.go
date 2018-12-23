/*package math

func Sin(x float64) float //implemented in assembly language
*/
package html

import (
	"io"
)

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumeNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	key, val string
}

func Parse(r io.Reader) (*Node, error)
