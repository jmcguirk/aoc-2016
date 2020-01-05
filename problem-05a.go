package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Problem5A struct {

}

func (this *Problem5A) Solve() {
	Log.Info("Problem 5A solver beginning!")

	doorId := "abbhdwsy";

	zero := uint8(int('0'));
	hasher := md5.New()
	index := 0;
	match := true;
	password := "";
	for{
		hasher.Reset();
		hasher.Write([]byte(doorId + strconv.Itoa(index)));
		hash := hex.EncodeToString(hasher.Sum(nil));
		match = true;
		for i := 0; i < 5; i++{
			if(hash[i] != zero){
				match = false;
				break;
			}
		}
		if(match){
			Log.Info("Found match %d", index);
			password += fmt.Sprintf("%c", hash[5]);
			if(len(password) >= 8){
				break;
			}

		}
		index++;
	}
	Log.Info("Generated password %s", password);
}
