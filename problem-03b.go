package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem3B struct {

}

func (this *Problem3B) Solve() {
	Log.Info("Problem 3B solver beginning!")

	file, err := os.Open("source-data/input-day-03b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	grid := IntegerGrid2D{};
	grid.Init();


	scanner := bufio.NewScanner(file)

	validTriangles := 0;
	row := 0;
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

			grid.SetValue(0, row, int(v1));
			grid.SetValue(1, row, int(v2));
			grid.SetValue(2, row, int(v3));
			row++;
		}

	}

	for i := 0; i <= row; i+=3{
		for j := 0; j <= 3; j++{
			v1 := grid.GetValue(j, i);
			v2 := grid.GetValue(j, i+1);
			v3 := grid.GetValue(j, i+2);
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
