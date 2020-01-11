package main

type Problem15A struct {

}

func (this *Problem15A) Solve() {
	Log.Info("Problem 15A solver beginning!")

	system := DiskSlotSystem{};
	err := system.Load("source-data/input-day-15a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	time := 0;
	for{
		if(system.Simulate(time)){
			Log.Info("First valid time was %d", time);
			break;
		}
		time++;
	}
}
