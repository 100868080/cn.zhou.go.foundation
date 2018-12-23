data,err:=json.Marshar(movies)
if err!=nil{
	log.Fatalf("Json marshaling failed:%s",err)
}
fmt.prinf("%s\n",data)

[{"Title":"Casablanca","released":1942,"Actor":["Humphrey Bogart",
"Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,
"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,
"Actors":["Steve McQueen","Jacqueline Bisset"]}]

data:=err:=json.MarshalIndent(movies,"","    ")
if err!=nil{
	log.Fatalf("JSON marshaling failed:%s",err)
}
fmt.printf("%s\n",data)