package main

type Problem12B struct {

}

func (this *Problem12B) Solve() {
	Log.Info("Problem 12A solver beginning!")

	machine := &IntcodeMachine{};
	err := machine.Init("source-data/input-day-12b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	machine.SetRegisterValue(int('c'), 1);
	err  = machine.Execute();
	if(err != nil){
		Log.FatalError(err);
	}
	registerOfInterest := int('a');
	val := machine.GetRegisterValue(registerOfInterest);
	Log.Info("Program finished executing - register %c has val %d", registerOfInterest, val);
}
