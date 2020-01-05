package main

type Problem11A struct {

}

func (this *Problem11A) Solve() {
	Log.Info("Problem 11A solver beginning!")

	problem := &ElevatorOptimizationProblem{};
	problem.Init(4, 1);
	problem.CaptureImages = false;
	useTestData := false;

	if(useTestData){
		problem.AddPOI(1, TypeMC, 'H')
		problem.AddPOI(1, TypeMC, 'L')

		problem.AddPOI(2, TypeGenerator, 'H')
		problem.AddPOI(3, TypeGenerator, 'L')
	} else{
		problem.AddPOI(1, TypeGenerator, 'P');
		problem.AddPOI(1, TypeMC, 'P');

		problem.AddPOI(2, TypeGenerator, 'C');
		problem.AddPOI(2, TypeGenerator, 'U');
		problem.AddPOI(2, TypeGenerator, 'R');
		problem.AddPOI(2, TypeGenerator, 'T');;

		problem.AddPOI(3, TypeMC, 'C');
		problem.AddPOI(3, TypeMC, 'U');
		problem.AddPOI(3, TypeMC, 'R');
		problem.AddPOI(3, TypeMC, 'T');
	}


	Log.Info("Initial state\n%s", problem.InitialState.Render());

	problem.Search();

}
