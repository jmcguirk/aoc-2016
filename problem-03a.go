package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem3A struct {

}

func (this *Problem3A) Solve() {
	Log.Info("Problem 3A solver beginning!")

	file, err := os.Open("source-data/input-day-03a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()



	scanner := bufio.NewScanner(file)

	validTriangles := 0;
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != "") {
			for j := 0; j < 5; j++{ // Replace ridiculous formatting
				line = strings.Replace(line,"  ", " ", -1);
			}
			lineParts := strings.Split(line, " ");

			v1, err := strconv.ParseInt(strings.TrimSpace(lineParts[0]), 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			v2, err := strconv.ParseInt(strings.TrimSpace(lineParts[1]), 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			v3, err := strconv.ParseInt(strings.TrimSpace(lineParts[2]), 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			if(v1 + v2 <= v3){
				continue;
			}
			if(v2 + v3 <= v1){
				continue;
			}
			if(v1 + v3 <= v2){
				continue;
			}
			validTriangles++;
		}
	}


	Log.Info("Checked triangles - %d are valid", validTriangles);
}
