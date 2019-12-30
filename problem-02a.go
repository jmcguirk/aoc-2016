package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem2A struct {

}

func (this *Problem2A) Solve() {
	Log.Info("Problem 2A solver beginning!")

	grid := &IntegerGrid2D{};
	grid.Init();
	// Populate the grid
	grid.SetValue(-1,-1,1);
	grid.SetValue(0,-1,2);
	grid.SetValue(1,-1,3);
	grid.SetValue(-1,0,4);
	grid.SetValue(0,0,5);
	grid.SetValue(1,0,6);
	grid.SetValue(-1,1,7);
	grid.SetValue(0,1,8);
	grid.SetValue(1,1,9);

	password := "";
	currPos := &IntVec2{};

	upChar := int('U');
	leftChar := int('L');
	rightChar := int('R');
	downChar := int('D');

	file, err := os.Open("source-data/input-day-02a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			for _, c := range line{
				mX := 0;
				mY := 0;
				switch(int(c)){
					case upChar:
						mY = -1;
						break;
					case downChar:
						mY = 1;
						break;
					case leftChar:
						mX = -1;
						break;
					case rightChar:
						mX = 1;
						break;
				}
				if(grid.HasValue(currPos.X + mX, currPos.Y + mY)){
					currPos.X += mX;
					currPos.Y += mY;
				}

			}
		}
		if(grid.HasValue(currPos.X, currPos.Y)){
			password += strconv.Itoa(grid.GetValue(currPos.X,currPos.Y));
		}
	}


	Log.Info("Calculated password: %s", password);
}
