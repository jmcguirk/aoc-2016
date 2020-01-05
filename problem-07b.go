package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem7B struct {

}


type IPABA struct {
	A uint8;
	B uint8;
}


func (this *Problem7B) Solve() {
	Log.Info("Problem 7B solver beginning!")
	file, err := os.Open("source-data/input-day-07b.txt");
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
			limit := len(line) - 3;
			pairs := make([]*IPABA, 0);
			for i, c := range line{
				if(c == openBracket){
					inBracket = true;
					continue;
				} else if(c == closeBracket){
					inBracket = false;
					continue;
				}
				if(i <= limit && !inBracket){
					if(line[i] == line[i+2] && line[i+1] != line[i]){
						pair := &IPABA{};
						pair.A = line[i];
						pair.B = line[i+1];
						pairs = append(pairs, pair);
					}
				}
			}
			containsMatch := false;
			for i, c := range line{
				if(c == openBracket){
					inBracket = true;
					continue;
				} else if(c == closeBracket){
					inBracket = false;
					continue;
				}
				if(i <= limit && inBracket){
					if(line[i] == line[i+2] && line[i+1] != line[i]){
						for _, pair := range pairs{
							if(pair.B == line[i] && pair.A == line[i+1]){
								containsMatch = true;
								break;
							}
						}
					}
				}
				if(containsMatch){
					break;
				}
			}
			if(containsMatch){
				supported++;
			}
		}
	}
	//114 too low
	Log.Info("Finished file parse! - %d addresses supported", supported);
}
