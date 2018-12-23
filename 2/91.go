value,ok:=vache.Lookup(key)
if !ok{
	//...cache[key] does not exist...
}

fmt.println(err)fmt.printf("%v",err)
resp,err:=http.Get(url)
if err!=nil{
	return nill,err
}
doc,err:=html.Parse(resp.Body)
resp.Body.Closr()
if err!=nil{
	return nil,fmt.Errorf("parsing %s as HTML:%v",url,err)
}

genesis:crashed:no parachute:G-switch failed:bad relay orientation