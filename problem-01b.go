package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Problem1B struct {

}

func (this *Problem1B) Solve() {
	Log.Info("Problem 1B solver beginning!")

	bytes, err := ioutil.ReadFile("source-data/input-day-01b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	grid := &IntegerGrid2D{};
	grid.Init();
	grid.SetValue(0,0,1);

	contents := string(bytes);
	parts := strings.Split(contents, ",")
	currPos := &IntVec2{};
	currOrientation := OrientationNorth;
	origin := &IntVec2{};
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

			foundPos := false;
			if(mX != 0){
				if(mX > 0){
					for i := 1; i <= mX; i++{
						if(grid.HasValue(currPos.X + i, currPos.Y)){
							currPos.X = currPos.X + i;
							foundPos = true;
							Log.Info("Found dupe location at %d, %d", currPos.X, currPos.Y);
							break;
						}
						grid.SetValue(currPos.X + i, currPos.Y, 1);
					}
				} else{
					flip := mX * -1;
					for i := 1; i <= flip; i++{
						if(grid.HasValue(currPos.X - i, currPos.Y)){
							currPos.X = currPos.X - i;
							foundPos = true;
							Log.Info("Found dupe location at %d, %d", currPos.X, currPos.Y);
							break;
						}
						grid.SetValue(currPos.X - i, currPos.Y, 1);
					}
				}
			} else{
				if(mY > 0){
					for i := 1; i <= mY; i++{
						if(grid.HasValue(currPos.X, currPos.Y + i)){
							currPos.Y = currPos.Y + i;
							foundPos = true;
							Log.Info("Found dupe location at %d, %d", currPos.X, currPos.Y);
							break;
						}
						grid.SetValue(currPos.X, currPos.Y + i, 1);
					}
				} else{
					flip := mY * -1;
					for i := 1; i <= flip; i++{
						if(grid.HasValue(currPos.X, currPos.Y - i)){
							currPos.Y = currPos.Y - i;
							foundPos = true;
							Log.Info("Found dupe location at %d, %d", currPos.X, currPos.Y);
							break;
						}
						grid.SetValue(currPos.X, currPos.Y - i, 1);
					}
				}
			}

			if(foundPos){
				break;
			}

			currPos.X += mX;
			currPos.Y += mY;

			stepCount++;


		}


	}


	Log.Info("After %d steps we are at %d,%d total distance: %d", stepCount, currPos.X, currPos.Y, currPos.ManhattanDistance(origin));

}
