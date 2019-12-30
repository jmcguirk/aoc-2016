package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem2B struct {

}

func (this *Problem2B) Solve() {
	Log.Info("Problem 2B solver beginning!")

	grid := &IntegerGrid2D{};
	grid.Init();
	// Populate the grid
	grid.SetValue(0,0,5);

	grid.SetValue(1,-1,2);
	grid.SetValue(1,0,6);
	grid.SetValue(1,1, int('A'));

	grid.SetValue(2,-2,1);
	grid.SetValue(2,-1,3);
	grid.SetValue(2,0,7);
	grid.SetValue(2,1, int('B'));
	grid.SetValue(2,2, int('D'));

	grid.SetValue(3,-1,4);
	grid.SetValue(3,0,8);
	grid.SetValue(3,1, int('C'));

	grid.SetValue(4,0,9);


	password := "";
	currPos := &IntVec2{};

	upChar := int('U');
	leftChar := int('L');
	rightChar := int('R');
	downChar := int('D');

	file, err := os.Open("source-data/input-day-02b.txt");
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
			val := grid.GetValue(currPos.X,currPos.Y);
			if(val < 10){
				password += strconv.Itoa(val);
			} else{
				password += fmt.Sprintf("%c", val);
			}

		}
	}


	Log.Info("Calculated password: %s", password);
}
