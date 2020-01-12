package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Problem20B struct {

}


func (this *Problem20B) Solve() {
	Log.Info("Problem 20A solver beginning!")
	file, err := os.Open("source-data/input-day-20b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	ipList := make([]*IPRange, 0);

	scanner := bufio.NewScanner(file)

	//var currVal int = 0;
	for scanner.Scan() { // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			val := &IPRange{};
			err := val.Parse(line);
			if(err != nil){
				Log.FatalError(err);
			}
			ipList = append(ipList, val);
		}
	}


	sort.SliceStable(ipList, func(i, j int) bool {
		return ipList[j].MaxIp < ipList[i].MaxIp;
	});

	reversed := make([]*IPRange, 0);
	reversed = append(reversed, ipList...);
	sort.SliceStable(reversed, func(i, j int) bool {
		return reversed[i].MinIp < reversed[j].MinIp;
	});
	currGuess := 0;
	total := 0;
	maxVal := 4294967295;
	for {
		found := false;
		for _, v := range ipList{
			if(currGuess >= v.MinIp && currGuess <= v.MaxIp){
				found = true;
				currGuess = v.MaxIp + 1;
				break;
			}
		}
		if(!found){

			found = false;
			for _, v := range reversed{
				if(currGuess < v.MinIp){
					found = true;
					total += v.MinIp - currGuess;
					currGuess = v.MinIp;
					break;
				}
			}
			if(!found){ // We are above max value
				break;
			}
		}
	}
	if(currGuess < maxVal){
		total += maxVal - currGuess;
	}
	// 105 too high
	Log.Info("Total matches %d, last match at %d", total, currGuess);
}
