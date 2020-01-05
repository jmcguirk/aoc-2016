package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type ElevatorOptimizationProblem struct {
	InitialState *LabSearchState;
	Visited map[uint64]int;
	VisitedStr map[string]int;
	FewestStepsDiscovered int;
	LastPrimeIndex int;
	CaptureImages bool;
	Primes []int;
}

const TypeGenerator = 'G';
const TypeMC = 'M';

func (this *ElevatorOptimizationProblem) Init(numFloors int, elevatorPos int) *LabSearchState {
	this.Primes = PrimesLessThan(100);
	res := &LabSearchState{};
	res.Depth = 0;
	res.Elevator = elevatorPos - 1;
	res.Floors = make([]*LabFloorState, numFloors);

	for i, _ := range res.Floors {
		state := &LabFloorState{};
		state.FloorNumber = i+1;
		state.UniqueId = this.Primes[this.LastPrimeIndex];
		this.LastPrimeIndex++;
		res.Floors[i] = state;

	}
	this.InitialState = res;
	return res;
}

func (this *ElevatorOptimizationProblem) Search(){
	start := this.InitialState.Clone();
	start.Depth = 0;
	start.Signature = 0;
	if(this.CaptureImages){
		start.CaptureScreenshot();
	}
	this.Visited = make(map[uint64]int);
	this.VisitedStr = make(map[string]int);
	this.FewestStepsDiscovered = int(math.MaxInt64);
	this.SearchIter(start);
}


func (this *ElevatorOptimizationProblem) MarkVisited(state *LabSearchState) {
	if(state.StringSignature == ""){
		state.CalculateSignatureString();
	}
	this.VisitedStr[state.StringSignature] = state.Depth;
}

func (this *ElevatorOptimizationProblem) IsVisited(state *LabSearchState) bool {
	if(state.StringSignature == ""){
		state.CalculateSignatureString();
	}
	v, exists := this.VisitedStr[state.StringSignature];
	if(!exists){
		return false;
	}
	return state.Depth >= v;
}

/*
func (this *ElevatorOptimizationProblem) MarkVisited(state *LabSearchState) {
	if(state.Signature == 0){
		state.CalculateSignature();
	}
	this.Visited[state.Signature] = state.Depth;
}

func (this *ElevatorOptimizationProblem) IsVisited(state *LabSearchState) bool {
	if(state.Signature == 0){
		state.CalculateSignature();
	}
	v, exists := this.Visited[state.Signature];
	if(!exists){
		return false;
	}
	return state.Depth >= v;
}*/

func (this *ElevatorOptimizationProblem) SearchIter(iv *LabSearchState) {


	iv.AssignScore();
	frontier := make([]*LabSearchState, 0);
	frontier = append(frontier, iv);

	allPowerSets := make(map[int]*[][]int);
	allFloorOptions := make(map[int]*[]int);

	for{
		if(len(frontier) <= 0){
			break;
		}


		sort.SliceStable(frontier, func(i, j int) bool {
			return frontier[i].Score > frontier[j].Score;
		});

		/*
		sort.SliceStable(frontier, func(i, j int) bool {
			return frontier[i].Depth < frontier[j].Depth;
		});*/

		state := frontier[0];
		frontier = frontier[1:];
		if(this.IsVisited(state)){
			continue;
		}
		this.MarkVisited(state);
		if(state.IsFinishCriteria()){
			if(state.Depth < this.FewestStepsDiscovered){
				this.FewestStepsDiscovered = state.Depth;
				Log.Info("Found new fewest steps %d", this.FewestStepsDiscovered);
				for i, v := range state.Path{
					Log.Info("Move %d\n%s\n", i, v);
				}
			}
			continue;
		}
		if(state.IsFailCriteria()){
			continue;
		}
		if(state.Depth >= this.FewestStepsDiscovered){
			continue;
		}

		currItems := state.Floors[state.Elevator].PointsOfInterest;
		if(len(currItems) == 0){
			Log.Info("Arrived on a floor with no items. This shouldn't be possible");
			continue;
		}


		_, exists := allFloorOptions[state.Elevator];
		if(!exists){
			floorOptions := make([]int, 0);
			if(state.Elevator > 0){
				floorOptions = append(floorOptions, state.Elevator-1);
			}
			if(state.Elevator < len(state.Floors) - 1){
				floorOptions = append(floorOptions, state.Elevator+1);
			}
			allFloorOptions[state.Elevator] = &floorOptions;
		}

		floorOptionsP, _ := allFloorOptions[state.Elevator];
		floorOptions := *floorOptionsP;

		_, exists = allPowerSets[len(currItems)];
		if(!exists){
			indices := make([]int, len(currItems));
			for i, _ := range currItems{
				indices[i] = i;
			}
			combinations := PowerSetInt(indices);
			filtered := make([][]int, 0);
			for _, ints := range combinations {
				if (len(ints) == 0 || len(ints) > 2) {
					continue;
				}
				filtered = append(filtered, ints);
				ReverseSlice(filtered);
			}
			allPowerSets[len(currItems)] = &filtered;
		}
		combinationsP, _ := allPowerSets[len(currItems)];
		combinations := *combinationsP;


		for _, ints := range combinations {
			if(len(ints) == 0 || len(ints) > 2) {
				continue;
			}
			for _, newFloor := range floorOptions{
				newState := state.Clone();
				newState.Depth++;
				oldFloor := state.Elevator;
				newState.Elevator = newFloor;
				item1 := newState.Floors[oldFloor].PointsOfInterest[ints[0]];
				var item2 *LabPointOfInterest;
				if(len(ints) == 2){
					item2 = newState.Floors[oldFloor].PointsOfInterest[ints[1]];
				}
				if(item2 != nil){
					newState.Floors[oldFloor].RemoveItemPair(item1, item2);
					newState.Floors[newFloor].AddItem(item1);
					newState.Floors[newFloor].AddItem(item2);
				} else{
					newState.Floors[oldFloor].RemoveItem(item1);
					newState.Floors[newFloor].AddItem(item1);
				}
				//newState.CalculateSignature();
				if(!this.IsVisited(newState)){
					if(this.CaptureImages){
						newState.CaptureScreenshot();
					}
					newState.AssignScore();
					frontier = append(frontier, newState);
				}
			}

		};
	}
}

func (this *ElevatorOptimizationProblem) AddPOI(floorNum int, poiType uint8, poiLetter uint8) {
	poi := &LabPointOfInterest{};
	poi.PointType = poiType;
	poi.UniqueId = this.Primes[this.LastPrimeIndex];
	this.LastPrimeIndex++;
	poi.PointLetter = poiLetter;
	floor := this.InitialState.Floors[floorNum-1];
	floor.PointsOfInterest = append(floor.PointsOfInterest, poi);
}

type LabSearchState struct {
	Floors []*LabFloorState;
	Depth int;
	Signature uint64;
	StringSignature string;
	Path []string;
	Elevator int;
	Score int;
}

func (this *LabSearchState) AssignScore() {
	sum := 0;
	for i, v := range this.Floors{
		sum += i * len(v.PointsOfInterest); // Favor things that move us higher
	}
	this.Score = sum;
}

func (this *LabSearchState) Clone() *LabSearchState {
	res := &LabSearchState{};
	res.Depth = this.Depth;
	res.Signature = 0;
	res.Elevator = this.Elevator;
	res.Floors = make([]*LabFloorState, len(this.Floors));
	for i, v := range this.Floors{
		res.Floors[i] = v.Clone();
	}
	res.Path = make([]string, len(this.Path));
	for i, v := range this.Path{
		res.Path[i] = v;
	}

	return res;
}

func (this *LabSearchState) CaptureScreenshot() {
	this.Path = append(this.Path, this.Render());
}

func (this *LabSearchState) CalculateSignature() {

	sig := uint64(0);
	for i, v := range this.Floors{
		floorSig := uint64(1);
		for _, poi := range v.PointsOfInterest{
			floorSig *= uint64(poi.UniqueId);
		}
		//Log.Info("Sig floor %d", floorSig);
		sig += (floorSig + uint64(math.Pow(10, float64(i*3))));
	}
	//Log.Info("Sig %s", strconv.FormatUint(sig, 10));
	this.Signature = sig;
}

func (this *LabSearchState) CalculateSignatureString() {
	this.StringSignature = this.RenderStringSignature();
}

func (this *LabSearchState) Render() string {
	var buff strings.Builder;
	for i := 3; i >= 0; i--{
		buff.WriteString(fmt.Sprintf("F%d", i+1));
		if(this.Elevator == i){
			buff.WriteString(" E ");
		} else{
			buff.WriteString(" . ");
		}
		buff.WriteString(this.Floors[i].Render());
		buff.WriteString("\n");
	}
	return buff.String();
}

func (this *LabSearchState) RenderStringSignature() string {
	var buff strings.Builder;
	buff.WriteString(fmt.Sprintf("E%d", this.Elevator));
	for i := 3; i >= 0; i--{
		pois := this.Floors[i].PointsOfInterest;
		sort.SliceStable(pois, func(i, j int) bool {
			return pois[i].UniqueId < pois[j].UniqueId;
		});
		buff.WriteString(fmt.Sprintf("F%d", i+1));
		for _, v := range pois{
			buff.WriteByte(byte(v.UniqueId));
		}

		buff.WriteString("\n");
	}
	return buff.String();
}

func (this *LabSearchState) IsFinishCriteria() bool {
	for i := 0; i < len(this.Floors) - 1; i++{
		if(len(this.Floors[i].PointsOfInterest) > 0){
			return false;
		}
	}
	return true;
}

func (this *LabSearchState) IsFailCriteria() bool {
	for i := 0; i < len(this.Floors); i++{
		if(this.Floors[i].IsFailCriteria()){
			return true;
		}
	}
	return false;
}

func (this *LabFloorState) IsFailCriteria() bool {
	hasGen := false;
	for _, poi := range this.PointsOfInterest{
		if(poi.PointType == TypeGenerator){
			hasGen = true;
			break;
		}
	}
	if(!hasGen){
		return false;
	}
	for _, poi := range this.PointsOfInterest{
		if(poi.PointType == TypeMC){
			match := false;
			for _, poi2 := range this.PointsOfInterest{
				if(poi2.PointType == TypeGenerator && poi2.PointLetter == poi.PointLetter){
					match = true;
					break;
				}
			}
			if(!match){
				return true;
			}
		}
	}
	return false;
}


func (this *LabFloorState) Render() string {
	var buff strings.Builder;
	for i := 0; i < 6; i++{
		buff.WriteString(" ");
		if(i < len(this.PointsOfInterest)){
			buff.WriteString(fmt.Sprintf("%c%c", this.PointsOfInterest[i].PointLetter, this.PointsOfInterest[i].PointType));
		} else{
			buff.WriteString(".");
		}
		buff.WriteString(" ");
	}
	return buff.String();
}


func (this *LabFloorState) RemoveItem(item *LabPointOfInterest) {
	filtered := make([]*LabPointOfInterest, 0);
	for _, v := range this.PointsOfInterest{
		if(v != item){
			filtered = append(filtered, v);
		}
	}
	this.PointsOfInterest = filtered;
}

func (this *LabFloorState) AddItem(item *LabPointOfInterest) {
	this.PointsOfInterest = append(this.PointsOfInterest, item);
}

func (this *LabFloorState) RemoveItemPair(item *LabPointOfInterest, item2 *LabPointOfInterest) {
	filtered := make([]*LabPointOfInterest, 0);
	for _, v := range this.PointsOfInterest{
		if(v != item && v != item2){
			filtered = append(filtered, v);
		}
	}
	this.PointsOfInterest = filtered;
}


func (this *LabFloorState) Clone() *LabFloorState {
	res := &LabFloorState{};
	res.FloorNumber = this.FloorNumber;
	res.UniqueId = this.UniqueId;
	res.PointsOfInterest = make([]*LabPointOfInterest, len(this.PointsOfInterest));
	for i, v := range this.PointsOfInterest{
		res.PointsOfInterest[i] = v.Clone();
	}
	return res;
}


func (this *LabPointOfInterest) Clone() *LabPointOfInterest {
	res := &LabPointOfInterest{};
	res.PointLetter = this.PointLetter;
	res.PointType = this.PointType;
	res.UniqueId = this.UniqueId;
	return res;
}

type LabFloorState struct {
	FloorNumber int;
	UniqueId int;
	PointsOfInterest []*LabPointOfInterest;
}

type LabPointOfInterest struct {
	PointType uint8;
	PointLetter uint8;
	UniqueId int;
}