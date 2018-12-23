for _,filename:=range filanames{
	f,err:=os.Open(filename)
	if err!=nil{
		return err 
	}
	defer f.Close()    //NOTE:risky;could run out of file
	descriptors
	//...process f...
}

for _,filename:=range filenames{
	if err!:=doFile(filename);err!=nil{
		return err 
	}
}

func doFile(filename string)error{
	f,err:=os.Open(filename)
	if err!=nil{
		return err
	}
	defer f.Close()
	//...process f...
}

func fetch(url string)(filename string,n int64,err error){
	resp,err:=http.Get(url)
	if err!=nil{
		return "",0,err
	}
	defer resp.Body.Close()
	local:=path.Base(resp.Reqest.URL.Path)
	if local=="/"{
		local="index.html"
}
f,err:=os.Create(local)
if err!=nil{
	return "",0,err
}
n,err=io.Copy(f,resp.Body)
//Close file,but prefer error from Copy,if any.
if closeErr:=f.Close();err==nil{
	err=closeErr
}
return local,n,err
}