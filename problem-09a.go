package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

type Problem9A struct {

}

func (this *Problem9A) Solve() {
	Log.Info("Problem 9A solver beginning!")


	bytes, err := ioutil.ReadFile("source-data/input-day-09a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	contents := string(bytes);

	len := 0;
	lengthBuff := "";
	magnitudeBuff := "";
	insideParen := false;
	delimEncountered := false;
	coolDown := 0;
	copyLength := int64(0);
	copyMagnitude := int64(0);

	openParen := int32('(')
	closeParen := int32(')')
	delim := int32('x');

	for _, c := range contents{
		if(unicode.IsSpace(c)){
			continue;
		}
		if(coolDown > 0){
			coolDown--;
			len++;
			continue;
		}
		if(c == openParen){
			magnitudeBuff = "";
			lengthBuff = "";
			insideParen = true;
			delimEncountered = false;
			continue;
		}
		if(insideParen){
			if(c == closeParen){
				copyLength, err = strconv.ParseInt(lengthBuff, 10, 64);
				if(err != nil){
					Log.Info("Failed to parse copy len");
					Log.FatalError(err);
				}

				copyMagnitude, err = strconv.ParseInt(magnitudeBuff, 10, 64);
				if(err != nil){
					Log.Info("Failed to parse copy len");
					Log.FatalError(err);
				}
				copyMagnitude = copyMagnitude - 1;
				insideParen = false;
				delimEncountered = false;
				coolDown = int(copyLength);
				len += int(copyLength * copyMagnitude);
				continue;
			} else if(c == delim){
				delimEncountered = true;
				continue;
			} else{
				if(delimEncountered){
					magnitudeBuff += fmt.Sprintf("%c", c);
				} else{
					lengthBuff += fmt.Sprintf("%c", c);
				}
			}
		} else{
			len++;
		}
	}
	// 112831 Too high - lol whitespace
	Log.Info(contents + " had uncompressed len %d", len);
}
