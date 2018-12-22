package bank

var deposits = make(chan int)
var balance = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func teller() {

	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func init() {
	go teller()
}
type Cake struct{state string}
func baker(cooked chan <-*Cake){
	for{
		cake:=new(Cake){
			for{
				cake:=new(Cake)
				cake.state="cooked"
			cooked<-cake
			}
}
func icer(iced chan<-*Cake,cooked<-chan *Cake){
	for cake:=range cooked{
		cake.state="iced"
		iced<-cake
	}
}