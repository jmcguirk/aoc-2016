package main

type Problem13A struct {

}

func (this *Problem13A) Solve() {
	Log.Info("Problem 13A solver beginning!")

	grid := IntegerGrid2D{}
	grid.Init();

	scanSize := 100;
	//offset := 10;
	offset:= 1362

	from := &IntVec2{};
	from.X = 1;
	from.Y = 1;

	to := &IntVec2{};
	to.X = 31;
	to.Y = 39;


	for j := 0; j <= scanSize; j++{
		for i := 0; i <= scanSize; i++{
			isWall := this.IsWall(offset, i, j);
			if(isWall){
				grid.SetValue(i, j, int('#'));
			} else{
				grid.SetValue(i, j, int('.'));
			}
		}
	}


	Log.Info("Finished setting up grid, starting pathing")
	path := grid.ShortestPath(from, to, int('#'));

	if(path == nil){
		Log.Info("Pathing failed");
	} else{
		Log.Info("Found path from %d,%d to %d,%d - contains %d steps", from.X, from.Y, to.X, to.Y, len(path));
	}
	//Log.Info("Initialized grid\n%s", grid.PrintAscii());
}


func (this *Problem13A) IsWall(offset int, x int, y int) bool {
	product := (x * x) + (3*x) + (2*x*y) + y + (y*y);
	product += offset;
	if(HasOddNumberInBinary(product)){
		return true;
	}
	return false;
}