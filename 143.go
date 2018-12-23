var w io.Writer
w = os.Stdout
f := w.(*os.File)
c := w.(*bytes.Buffer)

var w io.Write
w = os.Stdout
rw := w.(io.ReadWriter)
w = new(ByteCount)
re = w.(io.ReadWriter)

var w io.Writer = os.Stdout
f, ok := w.(*os.File)

b, ok := w.(*bytes.Buffer)