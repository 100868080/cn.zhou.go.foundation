package main

import (
	"sync"
	"image"
)
func loadIcons(){
	icons=map[string]image.Image{
		"spades.png": loadIcon("spades.png"),
		"hearts.png": loadIcon("hearts.png"),
		"daimonds.png": loadIcon("daimonds.png"),
		"clubs": loadIcond("clubs.png")
		
	}
}


func Icon(name string) image.Image{
	if icons==nil{
		loadIcons()
		
	}
	return icons[name]
}

func loadIcons(){
	icons=make(map[string]image.Image)
	icons["spades.png"]=loadIcon("spades.png")
	icons["hearts.png"]=loadIcon("hearts.png")
	icons["diamonds.png"]=loadIcon("daimonds.png")
	icons["clubs.png"]=loadIcon("clubs.png")
}

var mu sync.Mutex
var icons map[string]image.Image

func Icon(name string) image.Image{
	mu.Lock()
	defer mu.Unlock()
	if icons ==nil{
		loadIcons()
	}
	return icons[name]
}

var mu sync.RWMutex
var icons map[string]image.Image
func Icon(name string) image.Image{
	mu.RLock()
	if icons!=nil{
		icon:=icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()
mu.Lock()
if icons==nil{
	loadIcons()
}
icon:=icons[name]
mu.Unlock()
return icon

}

var loadIconsOnce sync.Once
var icons map[string]image.Image
func Icon(name string)image.Image{
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}