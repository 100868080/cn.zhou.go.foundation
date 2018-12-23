package gzip

import (
	"io"
)

func NewWrite(w io.Writer) io.WriteCloser
func NewReader(r io.Reader) (io.ReadWriteCloser, error)


#include <bzlib.H>
int bz2compress(bz_stream*s,int action,char *in,unsigned *inlen,char *out,unsigned *outlen){
	s->next_in=in;
	s->avail_in=*inlen;
	s->next_out=*outlen;
	s->avail_out=*outlen;
	int r=BZ2_bzCompress(s,action);
	*inlen-=s>avail_in;
	*outlen-=s->avail_out;
	s->next_in=s->next_out=NULL;
	return r;
	
}