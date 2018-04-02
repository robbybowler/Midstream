package main

import (
	"math/rand"
	"fmt"
	"time"
)


type person struct{
	fName string
	lName string
	age int
	SSN int
}
func main(){
	FN := [5]string{"john", "bob", "joe", "harry","andy"}
	LN := [5]string{"hansen","bowler","garcia","yoho","danger"}
	rand.Seed(time.Now().UTC().UnixNano())
	var people []person

	for i := 0; i < 10; i++ {
		randFN := FN[rand.Intn(len(FN))]
		randLN := LN[rand.Intn(len(LN))]
		randAge :=rand.Intn(90)+10
		randSSN :=rand.Intn(899999999)+1000000000

		p := person{
			fName:randFN,
			lName:randLN,
			age:randAge,
			SSN:randSSN,
		}
		people = append(people, p)
	}
	for _,p := range people{
		fmt.Printf("%[3]v %[2]v %[1]d \n",p.age,p.SSN,p.lName)
	}
	fmt.Printf("\n")
	fmt.Print("\n")

	sortedSlice := MergeSort(people)
	for _,p := range sortedSlice{
		fmt.Printf("%[3]v, %[4]v SSN: %[2]v Age: %[1]d \n",p.age,p.SSN,p.lName,p.fName)
	}
}



func MergeSort(people []person) []person {

	if len(people) < 2 {
		return people
	}
	mid := (len(people)) / 2
	return MergeSortR(MergeSort(people[:mid]), MergeSort(people[mid:]))
}


//func MergeSortR(L, R []person) []person {
//
//	size := len(L)+len(R)
//	i := 0
//	j := 0
//	mergeSlice := make([]person, size)
//
//	for k := 0; k < size; k++ {
//		if i > len(L)-1 && j <= len(R)-1 {
//			mergeSlice[k] = R[j]
//			j++
//		} else if i <= len(L)-1 && j > len(R)-1{
//			mergeSlice[k] = L[i]
//			i++
//		} else if L[i].lName <= R[j].lName{
//			if L[i].lName == R[j].lName {
//				if L[i].fName <= R[j].fName{
//					if L[i].fName == R[j].fName{
//						if L[i].age > R[j].age {
//							mergeSlice[k] = L[i]
//							i++
//						}else if(L[i].age == R[j].age) {
//							L[i], R[j] = R[j],L[i]
//							mergeSlice[k] = L[i]
//							i++
//						}else{
//							mergeSlice[k] = R[j]
//							j++
//						}
//					}else{
//						mergeSlice[k] = L[i]
//						i++
//					}
//
//				}else {
//					mergeSlice[k] = R[j]
//					j++
//				}
//
//			}else{
//				mergeSlice[k] = L[i]
//				i++
//			}
//		}  else {
//			mergeSlice[k] = R[j]
//			j++
//		}
//	}
//	return mergeSlice
//}

func MergeSortR(L, R []person) []person {

	size := len(L)+len(R)
	i := 0
	j := 0
	mergeSlice := make([]person, size)

	for k := 0; k < size; k++ {
		switch {
		case i > len(L)-1 && j <= len(R)-1:
			mergeSlice[k] = R[j]
			j++
		case i <= len(L)-1 && j > len(R)-1:
			mergeSlice[k] = L[i]
			i++
		case L[i].lName > R[j].lName:
			mergeSlice[k] = R[j]
			j++
		case L[i].lName < R[j].lName:
			mergeSlice[k] = L[i]
			i++
		case L[i].lName == R[j].lName && L[i].fName != R[j].fName:
			//same last name
			switch{
			case L[i].fName < R[j].fName:
				mergeSlice[k] = L[i]
				i++
			case L[i].fName > R[j].fName:
				mergeSlice[k] = R[j]
				j++
			}
		case L[i].lName == R[j].lName && L[i].fName == R[j].fName:
			//same last and first name
			switch{
			case L[i].age > R[j].age:
				//order from oldest to youngest
				mergeSlice[k] = L[i]
				i++
			case L[i].age == R[j].age:
				//same name and age results reverses order from how they were received.
				L[i], R[j] = R[j],L[i]
				mergeSlice[k] = L[i]
				i++
			case L[i].age < R[j].age:
				mergeSlice[k] = R[j]
				j++
			}

		}
	}
	return mergeSlice
}

