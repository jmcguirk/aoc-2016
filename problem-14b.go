package main

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"strconv"
)

type Problem14B struct {
	Cache map[int]string;
	Hasher hash.Hash;
	Salt string;
	HashRepeats int;
}

func (this *Problem14B) Solve() {
	Log.Info("Problem 14B solver beginning!")
	this.Salt = "zpqevtbw";
	//this.Salt = "abc";
	this.HashRepeats = 2016;
	this.Hasher = md5.New();
	this.Cache = make(map[int]string);
	keysDesired := 64;

	index := 0;
	keysGenerated := 0;
	for{
		tripletFound := false;
		tripleValue := int32(0);
		fiveletFound := false;
		hash := this.GetHash(index);

		for i, c := range hash{
			if(i >= 2){
				if(hash[i-1] == hash[i-2] && hash[i-1] == uint8(c)){
					tripletFound = true;
					tripleValue = c;
					break;
				}
			}
		}
		if(tripletFound){
			for j := index+1; j <= index+1000; j++{
				forwardHash := this.GetHash(j);
				accum := 0;
				for _, c := range forwardHash{
					if(c == tripleValue){
						accum++;
						if(accum >= 5){
							fiveletFound = true;
							break;
						}
					} else{
						accum = 0;
					}
				}
				if(fiveletFound){
					break;
				}
			}

		}
		if(fiveletFound){
			Log.Info("Found password %s at index %d", hash, index);
			keysGenerated++;
			if(keysGenerated >= keysDesired){
				break;
			}

		}
		this.DeleteCacheValue(index);
		index++;
	}

}

func (this *Problem14B) GetHash(index int)string{
	_, exists := this.Cache[index];
	if(!exists){
		this.Hasher.Reset();
		this.Hasher.Write([]byte(this.Salt + strconv.Itoa(index)));
		iv := hex.EncodeToString(this.Hasher.Sum(nil));
		for i := 0; i < this.HashRepeats; i++{
			this.Hasher.Reset();
			this.Hasher.Write([]byte(iv));
			iv = hex.EncodeToString(this.Hasher.Sum(nil));
		}
		this.Cache[index] = iv;
	}
	val, _ := this.Cache[index];
	return val;
}

func (this *Problem14B) DeleteCacheValue(index int) {
	delete(this.Cache, index);
}
