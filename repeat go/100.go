package main

import (
	"os"
)
var rmdirs []func()
for _,d:=range tempDirs(){
	dir:=d:  //NOte:necessary!
	os.MkdirAll(dir,0755)
	rmdirs=append(rmdirs,func()){
		os.RemoveAll(dir)
	})
}
for _,rmdirs:=range rmdirs{
	rmdirs()    //clean up
}

var rmdirs[]func
for _,dir:=range tempDirs(){
	os.MkdirAll(dir,0755)
	rmdirs=append(rmdirs,func()){
		os.RemoveAll(dir)
	})
}

for _,dir:=range tempDirs(){
	dir:=dir
}

var rmdirs [] func()
dirs:=tempDirs()

for i:=0;i<len(dirs);i++{
	os.MkdirAll(dirs[i],0755)
	rmdirs=append(rmdirs,func()){
		os.RemoveAll(dirs[i])   //NOTE:incorrect!
	})
}