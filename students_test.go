package coverage

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW


func TestPersonLen (t *testing.T) {
	cases := map[string]struct {
		people     People
		needResult int
	}{
		"zeroValue": {
			people:     []Person{},
			needResult: 0,
		},
		"twoPerson": {
			people: []Person{{"Ivan", "Ivanin", time.Now().Add(time.Second * 2)},
				{"Vasia", "Vasin", time.Now().Add(time.Second * 3)},
			},
			needResult: 2,
		},
		"FivePerson": {
			people: []Person{{"Ivan", "Ivanin", time.Now().Add(time.Second * 1)},
				{"Vasia", "Vasin", time.Now().Add(time.Second * 2)},
				{"Ivan2", "Ivanin2", time.Now().Add(time.Second * 3)},
				{"Vasia2", "Vasin2", time.Now().Add(time.Second * 4)},
				{"Vasia3", "Vasin3", time.Now().Add(time.Second * 5)},
			},
			needResult: 5,
		},
	}

	for name,testedValue:= range cases{
		if k:=testedValue.people.Len() ;k!=testedValue.needResult{
			t.Errorf("Bad result in %s test, put result - %d, need result %d",name, k, testedValue.needResult)
		}else {
			t.Logf("Test %s pass good",name)
		}
	}

}

func TestPersonLess (t *testing.T) {
	now := time.Now()

	people := People{
		{"AA", "BB", now},
		{"AA", "BB", now},
		{"AA", "CC", now},
		{"Vasia2", "Vasin2", now.Add(time.Second * 4)},
		{"Vasia3", "Vasin3", now.Add(time.Second * 5)},
		{"Vasia4", "Vasin3", now.Add(time.Second * 5)},
	}
	if people.Less(0, 1) {
		t.Errorf("Equal values")
	}
	t.Logf("Values %v, %v .Answer %t:",people[0],people[1],people.Less(0, 1))
	if people.Less(0, 3) {
		t.Errorf("Equal by lastname")
	}
	t.Logf("Values %v, %v.Answer: %t",people[0],people[3],people.Less(0, 3))
	if people.Less(0, 4) {
		t.Errorf("Less by birthday")
	}
	t.Logf("Values %v, %v.Answer: %t",people[0],people[4],people.Less(0, 4))
	if people.Less(5, 4) {
		t.Errorf("Less by FirstName")
	}
	t.Logf("Values %v, %v.Answer: %t",people[4],people[5],people.Less(0, 4))
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

func TestNew(t *testing.T) {
	matrix,err:=New("Вышел Заяц На Крыльцо")
	if matrix!=nil ||  !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Error in string type")
	} else{
		t.Log("Test with incorrect matrix pass good")
	}

	matrix,err=New("1 1 \n 2 2")
	needMartix:=Matrix{rows: 2,
		cols: 2,
		data: []int{1,1,2,2},
	}
	if  matrix.cols!=needMartix.cols ||  matrix.rows!=needMartix.rows || !reflect.DeepEqual(matrix.data,needMartix.data) {
		t.Errorf("Error in creating matrix 2X2")
		t.Logf("Equal cols => %t",matrix.cols!=needMartix.cols)
		t.Logf("Equal rows => %t",matrix.rows!=needMartix.rows)
		t.Logf("need value %v, get value %v",matrix.data,needMartix.data)
	}else {
		t.Log("Test with correct matrix pass good")
	}

	matrix,err=New("1 1 \n 2")
	if err.Error()!=("Rows need to be the same length"){
			t.Errorf("Error in creating matrix with not same length")
		}else{
			t.Log("Test with not the same length pass good")
}
	matrix,err=New("")
	if matrix!=nil ||  !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Error in string type")
	} else{
		t.Log("Test with empty matrix pass good")
	}
}

func TestMatrix_Rows(t *testing.T) {
cases:=map[string]struct{
	m Matrix
	needValue [][]int
}{
	"Matrix2x2":{
	m: Matrix{2,2,[]int{1,2,3,4}},
	needValue : [][]int{{1,2},{3,4}},
},
	"Matrix0x0":{
		m: Matrix{0,0,[]int{1,2,3,4}},
		needValue : [][]int{},
	},
}

	for k,_:= range cases{
		if !reflect.DeepEqual(cases[k].m.Rows(),cases[k].needValue){
			t.Errorf("Error in %v test. Need result %v, get %v",k,cases[k].needValue,cases[k].m.Rows())
		}


	}

}


func TestMatrix_Cols(t *testing.T) {
	cases:=map[string]struct{
		m Matrix
		needValue [][]int
	}{
		"Matrix2x2":{
			m: Matrix{2,2,[]int{1,2,3,4}},
			needValue : [][]int{{1,3},{2,4}},
		},
		"Matrix0x0":{
			m: Matrix{0,0,[]int{1,2,3,4}},
			needValue : [][]int{},
		},
	}

	for k,_:= range cases{
		if !reflect.DeepEqual(cases[k].m.Cols(),cases[k].needValue){
			t.Errorf("Error in %v test. Need result %v, get %v",k,cases[k].needValue,cases[k].m.Cols())
		}


	}

}


func TestMatrix_Set(t *testing.T) {
	cases := map[string]struct {
		m      *Matrix
		setRow int
		setCol int
		result  bool
	}{
		"MatrixSet2x2": {
			m:      &Matrix{2, 2, []int{1, 2, 3, 4}},
			setRow: 1,
			setCol: 1,
			result:  true,
		},
		"MatrixSet3x3": {
			m:     &Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			setRow: 2,
			setCol: 2,
			result:  true,
		},
		"MatrixSetIncorrect": {
			m:     &Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			setRow: 3,
			setCol: 4,
			result:  false,
		},
		"MatrixSetMinusValue": {
			m:     &Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			setRow: -1,
			setCol: 1,
			result:  false,
		},
	}
for k,_ := range cases{
		if result:= cases[k].m.Set(cases[k].setCol, cases[k].setRow, 111);result!=cases[k].result{
			t.Errorf("Error in %s test",k)
		}
	}
}