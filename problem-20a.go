package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Problem20A struct {

}

type IPRange struct{
	MinIp int;
	MaxIp int;
}

func (this *IPRange) ToString() string {
	return fmt.Sprintf("%d-%d", this.MinIp, this.MaxIp);
}

func (this *IPRange) Parse(line string) error {
	parts := strings.Split(line, "-");
	val, err := strconv.ParseInt(parts[0], 10, 64);
	if(err != nil){
		return err;
	}
	this.MinIp = int(val);
	val, err = strconv.ParseInt(parts[1], 10, 64);
	if(err != nil){
		return err;
	}
	this.MaxIp = int(val);
	return nil;
}

func (this *Problem20A) Solve() {
	Log.Info("Problem 20A solver beginning!")
	file, err := os.Open("source-data/input-day-20a.txt");
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

	/*
	for _, v := range ipList{
		Log.Info(v.ToString());
	}*/

	currGuess := 0;

	for {
		found := false;
		for _, v := range ipList{
			if(currGuess >= v.MinIp && currGuess <= v.MaxIp){
				found = true;
				currGuess = v.MaxIp + 1;
			}
		}
		if(!found){
			break;
		}
	}
	Log.Info("Found match at %d", currGuess);
}
