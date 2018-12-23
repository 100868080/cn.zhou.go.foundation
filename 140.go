package main

import (
	"strings"
	"fmt"
	"math"
	"testing"
)
func TestEval(t *testing.T){
	tests:=[]struct{
		expr string
		env Env
		want string
	}{
		{"sqrt(A/pi)",Env{"A":87616,"pi":math.Pi},"167"}
		{"pow(x,3)+pow{y,3}",Env{"x":12,"y":1},"1729"},
		{"pow(x,3)+pow(y,3)",Env{"x":9,"y":10},"1729"},
		{"5/9*(F-32)",Env{"F":-40},"-40"},
		{"5/9*(F-32)",Env{"F":32},"0"},
		{"5/9*(F-32)",Env{"F",212},"100"}
	}
	var prevExpr string
	for _,test:=range tests{
		if test.expr!=prevExpr{
			fmt.Printf("\n%s\n",test.expr)
			prevExpr=test.expr
		}
		expr,err:=Parse(test.expr)
		if err!=nil{
			t.Error(err)
			continue
		}
		got:=fmt.Sprintf("%.6f",expr.Eval(test.env))
		fmt.Printf("\t%v=>%s\n",test.env,got)
		if got!=test.want{
			t.Error("%s.Eval()in %v =%q,want %q\n",test.expr,test.env,got,test.want)
		}
	}
}

type Expr interface{
	Eval(env Env)float64
	Check(vars map[Var]bool)error
}

func (v Var)Check(vars map[Var]bool)error{
	vars[v]=true
	return nil
}

func (literal)CHeck(vars map[Var]bool)error{
	return nil
}
func (u unart)Check(vars map[Var]bool)error{
	if !strings.ContainsRune("+-",u.op){
		return fmt.Errorf("unexpexted unary op %q",u.op)
	}
	return u.x.Check(vars)
}
func (b binary)Check(vars map[Var]error){
	if !strings.ContainsRune("+-",b.op){
		return fmt.Errorf("unexpected binary op %q",b.op)
	}
	if err:=b.x.Check(vars);err!=nil{
		return err
	}
	return b.y.Check(vars)
}
func (c call)Check(vars map[Var]bool)error{
	arity,ok:=numParams[c.fn]
	if !ok{
		return fmt.Errorf("unknown functin !q",c.fn)
	}
	if len(c.args)!=arity{
		return fmt.Errorf("call to %s has %d args,want %d",c.fn,len(c.args),arity)
	}
	for _,arg:=range c.args{
		if err:=arg.Check(vars);err!=nil{
			return err
		}
	}
	return nil
}
var numParams=map[string]int{"pow":2,"sin":1,"sqrt":1}