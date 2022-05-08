package main
import ("fmt"
       )
func Swap(aInNumberToSort[] int , position int){
      temp := aInNumberToSort[position]
	  aInNumberToSort[position]=aInNumberToSort[position+1];
	  aInNumberToSort[position+1]=temp;
}	   
func BubbleSort(aInNumberToSort[] int){
    var isSwapped bool
    for {
	    isSwapped = false
	    for index:=0 ; index < len(aInNumberToSort)-1 ; index++{
		    if (aInNumberToSort[index] > aInNumberToSort[index+1]) {
			   Swap(aInNumberToSort,index)
			   isSwapped = true
			}
		}
	    if (!isSwapped){
		  break;
		}
	}
}	   
func main() {
    // Please free to change the input array 
    var numberToBeSortedSlice= [] int {50,60,-2,11,-1,45,67,89,1,7}
    BubbleSort(numberToBeSortedSlice)
	fmt.Println("Slice After Sortng")
	fmt.Println(numberToBeSortedSlice)
}
