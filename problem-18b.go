package main

type Problem18B struct {

}

func (this *Problem18B) Solve() {
	Log.Info("Problem 18B solver beginning!");

	// Special note - A much cheaper way to do this is to represent just the previous generation rather than tracking the entirety generated so far
	// However, the naive solution is more than fast enough

	grid := &IntegerGrid2D{};
	grid.Init();
	initialState := ".^^^^^.^^^..^^^^^...^.^..^^^.^^....^.^...^^^...^^^^..^...^...^^.^.^.......^..^^...^.^.^^..^^^^^...^.";
	width := len(initialState);
	for i, c := range initialState {
		grid.SetValue(i, 0, int(c));
	}

	generations := 400000;
	for n := 1; n < generations; n++{

		for i := 0; i < width; i++{
			isTrap := false;

			left := grid.CountIf(i-1, n-1, TrapGridChar) > 0;
			center := grid.CountIf(i, n-1, TrapGridChar) > 0;
			right := grid.CountIf(i+1, n-1, TrapGridChar) > 0;

			if((left && center) && !right){
				isTrap = true;
			} else if ((right && center) && !left){
				isTrap = true;
			} else if(left && !right && !center){
				isTrap = true;
			} else if(right && !left && !center){
				isTrap = true;
			}


			if(isTrap){
				grid.SetValue(i, n, TrapGridChar);
			} else{
				grid.SetValue(i, n, SafeGridChar);
			}

		}
		if(n % 10000 == 0){
			Log.Info("Completed %d generations", n);
		}
	}

	Log.Info("Grid after %d generations with %d safe tiles", generations, grid.CountAll(SafeGridChar));
}
