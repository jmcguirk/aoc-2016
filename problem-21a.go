package main

type Problem21A struct {

}


func (this *Problem21A) Solve() {
	Log.Info("Problem 21A solver beginning!")
	system := &PasswordSwapSystem{};
	iv := "abcdefgh";
	err := system.Init(iv, "source-data/input-day-21a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	err = system.Execute();
	if err != nil {
		Log.FatalError(err);
	}
	//Log.Info(system.PrintProgram());
	pw := system.PrintPassword();
	Log.Info("Final password is %s", pw);

}
