var _fmt.Stringer=&select

var _fmt.Stringer=s 
  os.Stdout.Write([] byte("hello"))
os.Stdout.Close()

var w io.Writer
w=os.Stdout
w.Write([]byte ("hello"))
w.Close()


var any interface{}
any=true
any=12.34
any="hello"
any=map[string]int{"one":1}
any=new{bytes.Buffer}

var w io.Writer=new(bytes.Buffer)

var _io.Writer=(*bytes.Buffer) (nil)

type Artifact interface{
	Title()string
	Creators() []string
	Created () time.Time
}

type Text interface{
	Pages()int
	Words()int
	PageSize()int
}
type Audio interface{
	Stream() (io.ReadCloser,error)
	RunningTime()time.Duration
	Format() string 
}
type Video interface{
	Stream() (io.ReadCloser,error)
	RunningTime()time.Duration
	Format()string
	Resolution()  (x,y int)
}

type Streamer interface{
	Stream() (io.ReadCloser,error)
	RunningTime()time.Duration
	Format()string
}
