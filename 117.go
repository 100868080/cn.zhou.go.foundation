m := url.Values{"lang": {"en"}}
m.Add("item", "1")
m.Add("item", "2")
fmt.println(m.Get("lang"))
fmt.Println(m.Get("q"))
fmt.println(m.Get("item"))
fmt.Println(m["item"])
m = nil
fmt.println(m.Get("item"))
m.Add("item", "3")