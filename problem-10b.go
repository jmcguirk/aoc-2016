package main

type Problem10B struct {

}

func (this *Problem10B) Solve() {
	Log.Info("Problem 10B solver beginning!")

	factory := RobotChipFactory{};
	err := factory.Init("source-data/input-day-10b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	//factory.HighValueOfInterest = 61;
	//factory.LowValueOfInterest = 17;
	//factory.HighValueOfInterest = 5;
	//factory.LowValueOfInterest = 2;
	factory.Simulate();
}
