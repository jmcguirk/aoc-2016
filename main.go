package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2016");

	solver := Problem19B{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
