package main

type Problem10A struct {

}

func (this *Problem10A) Solve() {
	Log.Info("Problem 10A solver beginning!")

	factory := RobotChipFactory{};
	err := factory.Init("source-data/input-day-10a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	factory.HighValueOfInterest = 61;
	factory.LowValueOfInterest = 17;
	//factory.HighValueOfInterest = 5;
	//factory.LowValueOfInterest = 2;
	factory.Simulate();
}
