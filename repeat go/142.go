package main

import (
	"math"
	"net/http"
)

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := ParseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr:"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}
