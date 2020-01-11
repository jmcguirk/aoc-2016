package main

type Problem12A struct {

}

func (this *Problem12A) Solve() {
	Log.Info("Problem 12A solver beginning!")

	machine := &IntcodeMachine{};
	err := machine.Init("source-data/input-day-12a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	err  = machine.Execute();
	if(err != nil){
		Log.FatalError(err);
	}
	registerOfInterest := int('a');
	val := machine.GetRegisterValue(registerOfInterest);
	Log.Info("Program finished executing - register %c has val %d", registerOfInterest, val);
}
