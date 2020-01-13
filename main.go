package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2016");

	solver := Problem22A{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
