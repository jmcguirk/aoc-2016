package main

import (
	"hash"
)

type Problem16A struct {
	Cache map[int]string;
	Hasher hash.Hash;
	Salt string;
}

func (this *Problem16A) Solve() {
	Log.Info("Problem 16A solver beginning!")

	iv := "01111001100111011";

	targetLen := 272;



	currentVal := make([]int, len(iv));
	for i, c := range iv{
		currentVal[i] = int(c) - int('0');
	}

	for{
		if(len(currentVal) >= targetLen){
			break;
		}

		valA := currentVal;
		valB := make([]int, len(valA));
		for i, v := range valA{
			if(v == 0){
				valB[len(valA) - i - 1] = 1;
			} else{
				valB[len(valA) - i - 1] = 0;
			}
		}
		currentVal = valA;
		currentVal = append(currentVal, 0);
		currentVal = append(currentVal, valB...);
	}


	Log.Info("Finished rewrite for %s of len %d - checksum is %s", iv, targetLen, IntArrayToString(this.CheckSum(currentVal[0:targetLen])));

}

func (this *Problem16A) CheckSum(val []int) []int{
	return this.CheckSumRecur(val, true);
}

func (this *Problem16A) CheckSumRecur(val []int, isInitial bool) []int{
	if(!isInitial && len(val) % 2 == 1){
		return val;
	}
	calc := make([]int, 0);
	for i := 1; i < len(val); i+=2{
		if(val[i] == val[i-1]){
			calc = append(calc, 1);
		} else{
			calc = append(calc, 0);
		}
	}
	return this.CheckSumRecur(calc, false);
}