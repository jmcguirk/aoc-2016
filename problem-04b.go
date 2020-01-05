package main

type Problem4B struct {

}

func (this *Problem4B) Solve() {
	Log.Info("Problem 4B solver beginning!")

	kiosk := EncryptedKiosk{};
	err := kiosk.Load("source-data/input-day-04b.txt");
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
