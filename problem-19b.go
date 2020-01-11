package main

type Problem19B struct {

}



func (this *Problem19B) Solve() {
	Log.Info("Problem 19B solver beginning!");

	elfCount := 3014387;
	elfState := make([]int, elfCount);

	indexMap := make(map[int]int);

	for i, _ := range elfState{
		elfState[i] = 1;
		indexMap[i] = i;
	}
	Log.Info("Initialized elf state array with %d values", elfCount);
	generationCount := 0;
	currMove := 0;
	for {
		newElves := 0;
		genSize := len(elfState);
		Log.Info("After %d generations, %d elves remain", generationCount, genSize);
		for {
			j := currMove % genSize;
			targetIndex := (j+(genSize/2)) % genSize;
			val := elfState[targetIndex];
			curr := elfState[j];
			if(curr > 0 && val > 0){
				if(targetIndex > j){
					currMove++;
				}
				elfState[targetIndex] = 0;
				elfState[j] = curr  + val;
				//mappedStealer, _ := indexMap[j];
				//mappedVictim, _ := indexMap[targetIndex];
				//Log.Info("%d steals %d at generation %d", mappedStealer + 1, mappedVictim + 1, generationCount);
				break;
			} else {
				currMove++;
			}
		}
		for _, v := range elfState{
			if(v > 0){
				newElves++;
			}
		}
		if(newElves == 1){
			Log.Info("Breaking after %d iterations", generationCount);
			for i, v := range elfState{
				if(v > 0){
					Log.Info("Lucky elf is %d", indexMap[i]+1);
				}
			}
			break;
		}
		newState := make([]int, newElves);
		newIndex := 0;
		for j, v := range elfState{
			if(v > 0){
				newState[newIndex] = v;
				oldMapping, _ := indexMap[j];
				indexMap[newIndex] = oldMapping;
				newIndex++;
			}
		}
		generationCount++;
		elfState = newState;
	}
}
