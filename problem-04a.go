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
	sum := 0;
	for _, k := range kiosk.Entries{
		if(!k.IsDecoy){
			sum += k.RoomNumber;
		}
	}
	Log.Info("Sum of real rooms is %d", sum);
}
