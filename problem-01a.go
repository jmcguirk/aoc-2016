package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Problem1A struct {

}

func (this *Problem1A) Solve() {
	Log.Info("Problem 1A solver beginning!")


	bytes, err := ioutil.ReadFile("source-data/input-day-01a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	contents := string(bytes);
	parts := strings.Split(contents, ",")
	currPos := &IntVec2{};
	currOrientation := OrientationNorth;
	stepCount := 0;
	for _, step := range parts{
		trimmed := strings.TrimSpace(step);
		if(trimmed != ""){
			dir := trimmed[0:1];
			mag64, err := strconv.ParseInt(trimmed[1:], 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			mag := int(mag64);
			isLeft := dir == "L";

			mX := 0;
			mY := 0;
			switch(currOrientation){
				case OrientationNorth:
					if(isLeft){
						mX = -mag;
						currOrientation = OrientationWest;
					} else{
						mX = +mag;
						currOrientation = OrientationEast
					}
					break;
				case OrientationWest:
					if(isLeft){
						mY = +mag;
						currOrientation = OrientationSouth;
					} else{
						mY = -mag;
						currOrientation = OrientationNorth
					}
				break;
				case OrientationEast:
					if(isLeft){
						mY = -mag;
						currOrientation = OrientationNorth
					} else{
						mY = +mag;
						currOrientation = OrientationSouth;
					}
				break;
			case OrientationSouth:
				if(isLeft){
					mX = mag;
					currOrientation = OrientationEast
				} else{
					mX = -mag;
					currOrientation = OrientationWest;
				}
				break;
			}
			currPos.X += mX;
			currPos.Y += mY;
			stepCount++;
		}


	}

	origin := &IntVec2{};
	Log.Info("After %d steps we are at %d,%d total distance: %d", stepCount, currPos.X, currPos.Y, currPos.ManhattanDistance(origin));

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
