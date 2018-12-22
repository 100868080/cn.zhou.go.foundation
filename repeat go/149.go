package main

import (
	"database/sql"
	"fmt"
)

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec(
		"SELECT*FROM tracks WHERE artist=?AND?<=year AND year<=?", artist, minYear, maxYear)
}

func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return sqlQuoteString(s)
	} else {
		panic(fmt.Sprintf("unexpected type %T:%v", x, x))
	}
}

switch x.(type){
	case nil:   //....
	case int,uint:    //...
	case bool:  //...
	case string:  //......
	default:  ///...
}

switch x:=x.(type)  {/*...*/}
