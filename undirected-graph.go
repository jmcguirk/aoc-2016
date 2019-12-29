package main


type UndirectedGraph struct {
	LastNodeId 		int;
	LastEdgeId 		int;
	Nodes  map[int]*Node;
	Edges  map[int]*Edge;
	LabelToNode map[string]*Node;
	LabelToEdge map[string]*Edge;
}

func (this *UndirectedGraph) Init() {
	this.LastNodeId = 1;
	this.Nodes = make(map[int]*Node);
	this.Edges = make(map[int]*Edge);
	this.LabelToNode = make(map[string]*Node);
}

func (this *UndirectedGraph) AllNodes()[]*Node {
	res := make([]*Node, 0);
	for _, node := range this.Nodes{
		res = append(res, node);
	}
	return res;
}




func (this *UndirectedGraph) GetOrCreateNode(label string)*Node {
	res, exists := this.LabelToNode[label];
	if(exists){
		return res;
	}
	res = &Node{};
	res.Init(this);
	res.Label = label;
	res.Id = this.LastNodeId;
	this.LastNodeId++;
	this.LabelToNode[res.Label] = res;
	this.Nodes[res.Id] = res;
	return res;
}

func (this *UndirectedGraph) CreateEdge(from *Node, to *Node)*Edge {
	return this.CreateEdgeWithWeight(from, to, 1);
}

func (this *UndirectedGraph) CreateEdgeWithWeight(from *Node, to *Node, weight int)*Edge {
	res := &Edge{};
	res.From = from;
	res.To = to;
	from.Edges = append(from.Edges, res);
	res.Id = this.LastEdgeId;
	res.Weight = weight;
	this.LastEdgeId++;
	this.Edges[res.Id] = res;
	res = &Edge{};
	res.From = to;
	res.To = from;
	to.Edges = append(to.Edges, res);
	res.Id = this.LastEdgeId;
	res.Weight = weight;
	this.LastEdgeId++;
	this.Edges[res.Id] = res;

	return res;
}

type CycleQuery struct{
	StartNode *Node;
	MinQuery bool;
	BestSoFar int;
	BestPath map[string]*Node;
}

func (this *UndirectedGraph) ShortestCycle() *CycleQuery {
	var query *CycleQuery;
	for _, node := range this.Nodes{
		cycle := this.ShortestCycleFrom(node);
		if(cycle.BestPath != nil && query == nil || query.BestSoFar > cycle.BestSoFar){
			query = cycle;
		}
	}
	return query;
}

func (this *UndirectedGraph) ShortestCycleFrom(from *Node) *CycleQuery {
	query := &CycleQuery{};
	query.MinQuery = true;
	query.StartNode = from;
	path := make(map[string]*Node);
	path[from.Label] = from;
	remaining := make(map[string]*Node);
	for _, node := range this.Nodes{
		if(node.Label != from.Label){
			remaining[node.Label] = node;
		}
	}
	this.CycleQueryRecur(from, query, 0, path, remaining);
	return query;
}

func (this *UndirectedGraph) LongestCycle() *CycleQuery {
	var query *CycleQuery;
	for _, node := range this.Nodes{
		cycle := this.LongestCycleFrom(node);
		if(cycle.BestPath != nil && query == nil || query.BestSoFar < cycle.BestSoFar){
			query = cycle;
		}
	}
	return query;
}

func (this *UndirectedGraph) LongestCycleFrom(from *Node) *CycleQuery {
	query := &CycleQuery{};
	query.MinQuery = false;
	query.StartNode = from;
	path := make(map[string]*Node);
	path[from.Label] = from;
	remaining := make(map[string]*Node);
	for _, node := range this.Nodes{
		if(node.Label != from.Label){
			remaining[node.Label] = node;
		}
	}
	this.CycleQueryRecur(from, query, 0, path, remaining);
	return query;
}


func (this *UndirectedGraph) CycleQueryRecur(from *Node, query *CycleQuery, totalSoFar int, pathSoFar map[string]*Node, remaining map[string]*Node) {
	if(len(remaining) == 0){
		if(query.MinQuery){
			if(query.BestPath == nil || totalSoFar < query.BestSoFar){
				query.BestPath = pathSoFar;
				query.BestSoFar = totalSoFar;
			}
		}else {
			if(query.BestPath == nil || totalSoFar > query.BestSoFar){
				query.BestPath = pathSoFar;
				query.BestSoFar = totalSoFar;
			}
		}
		return;
	}
	for _, edge := range from.Edges {
		_, exists := pathSoFar[edge.To.Label];
		if(exists){
			continue;
		}
		neighbor := edge.To;
		cpyPath := make(map[string]*Node);
		for k, v := range pathSoFar{
			cpyPath[k] = v;
		}
		cpyPath[neighbor.Label] = neighbor;
		cpyRemaining := make(map[string]*Node);
		for k, v := range remaining{
			if(k != neighbor.Label) {
				cpyRemaining[k] = v;
			}
		}
		this.CycleQueryRecur(edge.To, query, totalSoFar+edge.Weight, cpyPath, cpyRemaining);
	}
}
