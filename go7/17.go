package main

import (
	"fmt"
)

func main() {
	var x, y int = 6, 9
	sum, div := f2(x, y)
	fmt.Println(sum, div)
}
func f2(a, b int) (int, float32) {
	return a * b, float32(a) / float32(b)
}

/*#include <stdio.h>
int main()
{
	int a[20];
	a[0]=2;
	a[1]=3;
	int b[20];
	b[0]=1;
	b[1]=2;
	int i;
	float p[20];
	int k;
	float r;
	if(1<i<=20)
{
	a[i]=a[i-1]+a[i-2];
	b[i]=b[i-1]+b[i-2];
	p[i]=a[i]/b[i];
}

for(k=0;k<20;k++)
{
	r+=p[k];
}
	//printf("%d\n",a[4]);
	//printf("%d\n",b[4]);
	printf("%f\n",r);
	return 0;
}*/
