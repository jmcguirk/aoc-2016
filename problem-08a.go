package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem8A struct {

}


func (this *Problem8A) Solve() {
	Log.Info("Problem 8A solver beginning!");

	gridHeight := 6;
	gridWidth := 50;

	//gridHeight := 3;
	//gridWidth := 7;

	grid := &IntegerGrid2D{};
	grid.Init();

	cellOff := int(' ');
	cellOn := int('#');

	for j := 0; j < gridHeight ;j++{
		for i := 0; i < gridWidth ;i++{
			grid.SetValue(i, j, cellOff);
		}
	}

	file, err := os.Open("source-data/input-day-08a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			lineParts := strings.Split(line, " ");
			op := lineParts[0];
			switch(op){
				case "rect":
					dims := strings.Split(lineParts[1], "x");
					width, err := strconv.ParseInt(dims[0], 10, 64);
					if(err != nil){
						Log.FatalError(err);
					}
					height, err := strconv.ParseInt(dims[1], 10, 64);
					if(err != nil){
						Log.FatalError(err);
					}
					for j:=0; j < int(height); j++{
						for i:=0; i < int(width); i++{
							grid.SetValue(i, j, cellOn);
						}
					}
					break;
				case "rotate":
					dir := lineParts[1];
					val64, err := strconv.ParseInt(strings.TrimSpace(strings.Split(lineParts[2], "=")[1]), 10, 64);
					if(err != nil){
						Log.FatalError(err);
					}
					mag64, err := strconv.ParseInt(strings.TrimSpace(lineParts[4]), 10, 64);
					if(err != nil){
						Log.FatalError(err);
					}
					val := int(val64);
					mag := int(mag64)
					if(dir == "column"){;
						cache := make([]int, 0);
						for j:=0; j < gridHeight; j++{
							cache = append(cache, grid.GetValue(val, j));
						}
						for i, v := range cache {
							grid.SetValue(val, (i + mag) % gridHeight, v);
						}
					} else if (dir == "row"){
						cache := make([]int, 0);
						for i:=0; i < gridWidth; i++{
							cache = append(cache, grid.GetValue(i, val));
						}
						for j, v := range cache {
							grid.SetValue((j + mag) % gridWidth, val, v);
						}
					}
			}
		}
	}



	Log.Info("Lights On %d \n%s", grid.CountAll(cellOn), grid.PrintAscii());
}
