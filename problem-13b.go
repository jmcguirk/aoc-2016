package main

type Problem13B struct {

}

func (this *Problem13B) Solve() {
	Log.Info("Problem 13B solver beginning!")

	grid := IntegerGrid2D{}
	grid.Init();

	scanSize := 100;
	//offset := 10;
	offset:= 1362;
	maxSteps := 50;

	from := &IntVec2{};
	from.X = 1;
	from.Y = 1;


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


	Log.Info("Finished setting up grid, starting exploration")
	visitedGrid := grid.ExploreWithDistance(from, int('#'), maxSteps);

	reachable := visitedGrid.CountAll(int('.'));
	Log.Info("Finished exploration, %d reachable nodes", reachable);
}


func (this *Problem13B) IsWall(offset int, x int, y int) bool {
	product := (x * x) + (3*x) + (2*x*y) + y + (y*y);
	product += offset;
	if(HasOddNumberInBinary(product)){
		return true;
	}
	return false;
}