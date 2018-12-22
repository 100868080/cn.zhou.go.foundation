log.println(findLinks(url))
links, err := findLinks(url)
log.Println(links, err)

func Size(rect image.Rectangle)(width,height int)
func Split(path string)(dir,file string)
func HourMinse(t time.Time) (hour,minute,second int)


func CountWordsAndImages(url string) (words,image int,err error){
	resp,err:=http.Get(url)
	if err!=nil{
		return 
	}
	doc,err:=html.Parse(resp.Body)
	resp.Body.Close()
	if err!=mil{
		err=fmt.Errorf("parsing HTML:%s",err)
		return 
	}
	words,image=CountWordsAndImages(doc)
	return
}
func CountWordsAndImages(n*html.Node)(words,image int){/* ... */}

return words,image,err