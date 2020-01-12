package main

import "fmt"

type Problem21B struct {
	Log bool;
}


func (this *Problem21B) Solve() {
	Log.Info("Problem 21B solver beginning!");
	this.Log = false;
	system := &PasswordSwapSystem{};

	// Note - the naive search approach is maybe not the intended way to solve this problem, but its fast enough < 10s to be viable
	// In addition, this loop considers obviously invalid patterns (with repeat letters). Ignoring these probably brings the total search to sub second
	// However, I suspect there's a relatively closed form solution using inverting the individual operations

	targetPW := "fbgdceah";
	maxGen := -1;
	err := system.Init(targetPW, "source-data/input-day-21b.txt");
	if err != nil {
		Log.FatalError(err);
	}

	maxOdometerReading := int('h') - int('a') + 1;
	indexArr := make([]int, len(targetPW));
	for i := 0; i < len(indexArr); i++ {
		indexArr[i] = 0;
	}
	genCount := 0;
	for{

		for i, v := range indexArr{
			system.Data[i] = v + int('a');
		}

		password := system.RunAgain();
		if(password == targetPW){
			sourcePass := "";
			for _, v := range indexArr{
				sourcePass += fmt.Sprintf("%c", v + int('a'));
			}
			Log.Info("Found target password %s using source pass %s", targetPW, sourcePass);
			break;
		} else{
			if(this.Log){
				sourcePass := "";
				for _, v := range indexArr{
					sourcePass += fmt.Sprintf("%c", v + int('a'));
				}
				Log.Info("%s = %s", sourcePass, password);
			}
		}

		atLim := false;
		for j := len(indexArr) - 1; j >= 0; j--{
			if(indexArr[j] + 1 < maxOdometerReading){
				indexArr[j]++;
				break;
			} else{
				if(j == 0){
					atLim = true;
					break;
				}
				indexArr[j] = 0;
			}
		}
		if(atLim){
			Log.Info("Odomoter hit max lim");
			break;
		}
		genCount++;
		if(maxGen > 0 && genCount > maxGen) {
			break;
		}
	}


	err = system.Execute();
	if err != nil {
		Log.FatalError(err);
	}
	//Log.Info(system.PrintProgram());
	pw := system.PrintPassword();
	Log.Info("Final password is %s", pw);

}
