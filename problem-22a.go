package main

type Problem22A struct {

}


func (this *Problem22A) Solve() {
	Log.Info("Problem 22A solver beginning!")
	system := &DataCenterSystem{};
	err := system.Init("source-data/input-day-22a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	viable := system.CountViable();
	// too high - 996804
	// too high - 995790
	Log.Info("Found %d viable pairs", viable);
}
