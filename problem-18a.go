package main

type Problem18A struct {

}

const TrapGridChar = '^';
const SafeGridChar = '.';

func (this *Problem18A) Solve() {
	Log.Info("Problem 18A solver beginning!");

	grid := &IntegerGrid2D{};
	grid.Init();
	initialState := ".^^^^^.^^^..^^^^^...^.^..^^^.^^....^.^...^^^...^^^^..^...^...^^.^.^.......^..^^...^.^.^^..^^^^^...^.";
	width := len(initialState);
	for i, c := range initialState {
		grid.SetValue(i, 0, int(c));
	}

	generations := 40;
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

	}

	Log.Info("Grid after %d generations with %d safe tiles:\n%s", generations, grid.CountAll(SafeGridChar), grid.PrintAscii());
}
