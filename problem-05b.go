package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Problem5B struct {

}

func (this *Problem5B) Solve() {
	Log.Info("Problem 5B solver beginning!")
	doorId := "abbhdwsy";

	//doorId := "abc";

	zero := uint8(int('0'));
	seven := uint8(int('7'));
	hasher := md5.New()
	index := 0;
	match := true;

	passwordArr := make([]uint8, 8);

	passwordCharsFound := 0;
	for{
		hasher.Reset();
		hasher.Write([]byte(doorId + strconv.Itoa(index)));
		hash := hex.EncodeToString(hasher.Sum(nil));
		if(hash[5] < zero || hash[5] > seven){
			index++;
			continue;
		}
		if(passwordArr[hash[5] - zero] > 0){
			index++;
			continue;
		}
		match = true;
		for i := 0; i < 5; i++{
			if(hash[i] != zero){
				match = false;
				break;
			}
		}
		if(match){
			location, err := strconv.ParseInt(fmt.Sprintf("%c", hash[5]), 10, 64);
			if(err != nil){
				Log.FatalError(err);
				continue;
			}
			if(location < 0 || location > 7){
				continue;
			}
			if(passwordArr[location] > 0){
				continue;
			}
			passwordArr[location] = hash[6];
			passwordCharsFound++;
			Log.Info("Found password char %d %d - hash: %s, %d", passwordCharsFound, location, hash, index);
			if(passwordCharsFound >= 8){
				break;
			}

		}
		index++;
	}
	password := "";
	for _, c := range passwordArr{
		password += fmt.Sprintf("%c", c);
	}
	Log.Info("Generated password %s", password);
}
