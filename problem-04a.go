package main

type Problem4A struct {

}

func (this *Problem4A) Solve() {
	Log.Info("Problem 4A solver beginning!")

	kiosk := EncryptedKiosk{};
	err := kiosk.Load("source-data/input-day-04a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
}
