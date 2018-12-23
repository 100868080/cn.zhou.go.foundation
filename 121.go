package main

import (
	"fmt"
)
func (s *IntSet)String()string{
	var bufbytes.Buffer
	buf.WriteByte('{')
	for i,word:=range s.words{
		if word==0{
			continue
		}
		for j:=0;j<64;j++{
			if word&(1<<uint(j))!=0{
				if buf.Len()>len("{"){
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf,"%d",64*i+j)"}")}}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}