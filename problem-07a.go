package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem7A struct {

}


func (this *Problem7A) Solve() {
	Log.Info("Problem 7A solver beginning!")
	file, err := os.Open("source-data/input-day-07a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	openBracket := int32('[');
	closeBracket := int32(']');

	supported := 0;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			inBracket := false;
			containsBracketedSequence := false;
			containsUnbracketedSequence := false;
			limit := len(line) - 4;
			for i, c := range line{
				if(c == openBracket){
					inBracket = true;
					continue;
				} else if(c == closeBracket){
					inBracket = false;
					continue;
				}
				if(i <= limit){
					if(line[i] == line[i+3] && line[i+1] == line[i+2] && line[i+1] != line[i]){
						if(inBracket){
							containsBracketedSequence = true;
							break;
						} else{
							containsUnbracketedSequence = true;
						}
					}
				}
			}
			if(containsUnbracketedSequence && !containsBracketedSequence ){
				supported++;
			}
		}
	}
	//114 too low
	Log.Info("Finished file parse! - %d addresses supported", supported);
}
