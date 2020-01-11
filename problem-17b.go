package main

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"math"
	"sort"
	"strings"
)

type Problem17B struct {
	Hasher hash.Hash;
	InitialState string;
	CanonicalGrid *IntegerGrid2D;
}

func (this *Problem17B) GenerateEdges(state *HashMazeState) []*HashMazeState {

	this.Hasher.Reset();
	this.Hasher.Write([]byte(state.PathHashSoFar));
	hash := hex.EncodeToString(this.Hasher.Sum(nil));
	var node *HashMazeState;
	x := state.PositionX;
	y := state.PositionY;
	res := make([]*HashMazeState,0);
	node = this.GenerateEdgeIfRelevant(hash, x,y-1, 0, "U", state);
	if(node != nil){
		res = append(res, node);
	}
	node = this.GenerateEdgeIfRelevant(hash, x,y+1, 1, "D", state);
	if(node != nil){
		res = append(res, node);
	}
	node = this.GenerateEdgeIfRelevant(hash, x-1,y, 2, "L", state);
	if(node != nil){
		res = append(res, node);
	}
	node = this.GenerateEdgeIfRelevant(hash, x+1,y, 3, "R", state);
	if(node != nil){
		res = append(res, node);
	}
	return res;
}

func (this *Problem17B) GenerateEdgeIfRelevant(hash string, posX int, posY int, index int, dir string, oldState *HashMazeState) *HashMazeState {
	if(this.IsFinishState(oldState)){
		return nil;
	}
	if(!this.CanonicalGrid.HasValue(posX, posY)){
		return nil;
	}
	char := hash[index];
	if(char >= 'b'){
		newState := &HashMazeState{};
		newState.PathHashSoFar = oldState.PathHashSoFar + dir;
		newState.PositionY = posY;
		newState.PositionX = posX;
		newState.AssignId();
		return newState;
	}
	return nil;
}

func (this *Problem17B) IsFinishState(state *HashMazeState) bool {
	return state.PositionX == 3 && state.PositionY == 3;
}

func (this *Problem17B) Solve() {
	Log.Info("Starting Problem 17B");
	this.CanonicalGrid = &IntegerGrid2D{};
	this.CanonicalGrid.Init();
	this.Hasher = md5.New();
	this.InitialState = "hhhxzeay";

	res := make([]*HashMazeState, 0);
	GridSize := 4;


	for i := 0; i < GridSize; i++{
		for j := 0; j < GridSize; j++{
			this.CanonicalGrid.SetValue(i, j, 1);
		}
	}

	var maxPath *HashMazeState;

	start := &HashMazeState{};
	start.PositionX = 0;
	start.PositionY = 0;
	start.PathHashSoFar = this.InitialState;
	start.AssignId();

	frontier := make([]*HashMazeState, 0);
	frontier = append(frontier, start);
	frontierMap := make(map[string]*HashMazeState);
	frontierMap[start.Id] = start;

	visitedNodes := make(map[string]*HashMazeState);
	minCostToStart := make(map[string]int);
	nearestToStart := make(map[string]*HashMazeState);


	minCostToStart[start.Id] = 0;
	var finalState *HashMazeState;
	for {
		if (len(frontier) <= 0) {
			break;
		}
		sort.SliceStable(frontier, func(i, j int) bool {
			return minCostToStart[frontier[i].Id] < minCostToStart[frontier[j].Id];
		});

		next := frontier[0];
		//Log.Info("Exploring %s - %d, %d", next.Id, next.PositionX, next.PositionY);
		frontier = frontier[1:];
		delete(frontierMap, next.Id);
		costToHere := minCostToStart[next.Id];
		edges := this.GenerateEdges(next);
		for _, neighbor := range edges{
			_, visited := visitedNodes[neighbor.Id];
			if(visited){
				continue;
			}

			bestToHere, bestCostExists := minCostToStart[neighbor.Id];
			if(!bestCostExists){
				bestToHere = int(math.MaxInt32);
			}

			if(costToHere + 1 < bestToHere){
				minCostToStart[neighbor.Id] = costToHere + 1;
				nearestToStart[neighbor.Id] = next;
				_, alreadyEnqueued := frontierMap[neighbor.Id];
				if(!alreadyEnqueued){
					frontierMap[neighbor.Id] = neighbor;
					frontier = append(frontier, neighbor);
				}
			}

		}
		visitedNodes[next.Id] = next;
		if(this.IsFinishState(next)){
			if(maxPath == nil || len(next.PathHashSoFar) > len(maxPath.PathHashSoFar)){
				maxPath = next;
				finalState = maxPath;
			}
		}
	}

	_, exists := minCostToStart[finalState.Id];
	if(!exists){
		Log.Fatal("Failed to find a path to destination")
	}

	nextPathStep := finalState.Id;

	for {
		next := nearestToStart[nextPathStep];
		if(next == start){
			break;
		}
		nextPathStep = next.Id;
		res = append(res, next);
	}

	ReverseSlice(res);

	Log.Info("Found max path %d steps" , len(strings.ReplaceAll(finalState.PathHashSoFar, this.InitialState, "")));
}
