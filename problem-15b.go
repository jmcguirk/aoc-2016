package main

type Problem15B struct {

}

func (this *Problem15B) Solve() {
	Log.Info("Problem 15B solver beginning!")

	system := DiskSlotSystem{};
	err := system.Load("source-data/input-day-15b.txt");
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
