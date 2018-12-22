type Movie struct{
	Title string
	Year int 'json:"released"'
	Color bool 'json:"color,omitempty"'
	Actors []string
}
var movies=[]Movie{
	{Title:"Casablanca",Year:1942,Coor:false,
	Actors:[]string{"Humphrey Bogart","Ingrid Bergman"}},
	{Title:"Cool Hand Luke",Year:1967,Color:true,
	Actors:[]string{"Paul newman"}},
	{Title:"Bullitt",Year:1968,Color:true,
	Actors:[]string{"Steve McQueen","Jacqueline Bisset"}},
	//...
}