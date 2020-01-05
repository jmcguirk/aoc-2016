package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

type Problem9B struct {

}

type StrDecompressModifier struct {
	Magnitude int;
	RemainingFrames int;
}

func (this *Problem9B) Solve() {
	Log.Info("Problem 9B solver beginning!")


	bytes, err := ioutil.ReadFile("source-data/input-day-09b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	contents := string(bytes);

	len := 0;
	lengthBuff := "";
	magnitudeBuff := "";
	insideParen := false;
	delimEncountered := false;
	copyLength := int64(0);
	copyMagnitude := int64(0);

	openParen := int32('(')
	closeParen := int32(')')
	delim := int32('x');

	activeModifiers := make([]*StrDecompressModifier, 0);

	for _, c := range contents{
		if(unicode.IsSpace(c)){
			continue;
		}
		for _, m := range activeModifiers{
			if(m.RemainingFrames > 0){
				m.RemainingFrames--;
			}
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
				insideParen = false;
				delimEncountered = false;
				coolDown := 1+ int(copyLength);
				mod := &StrDecompressModifier{};
				mod.RemainingFrames = coolDown;
				mod.Magnitude = int(copyMagnitude);
				activeModifiers = append(activeModifiers, mod);
				//len += int(copyLength * copyMagnitude);
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
			prod := 1;
			for _, m := range activeModifiers {
				if(m.RemainingFrames > 0){
					prod *= m.Magnitude;
				}
			}
			len += prod;
		}
	}
	// 112831 Too high - lol whitespace
	Log.Info(contents + " had uncompressed len %d", len);
}
