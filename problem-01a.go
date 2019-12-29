package main

import (

)

type Problem1A struct {

}

func (this *Problem1A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	/*
	file, err := os.Open("source-data/input-day-01a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currVal int = 0;
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		for _, c := range line{
			if(int(c)) == int('('){
				currVal++
			}else if(int(c)) == int(')'){
				currVal--
			}
		}
	}*/
}
