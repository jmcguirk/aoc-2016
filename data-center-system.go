package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DataCenterSystem struct {
	FileName string;
	CanonicalGrid *IntegerGrid2D;
	DataNodes map[int]*DataCenterNode;
}

type DataCenterNode struct {
	Index int;
	X int;
	Y int;
	Avail int;
	Size int;
	InitialData int;
	CurrentData int;
	Name string;
	Neighbors []*DataCenterNode;
}

func (this *DataCenterNode) IsViable(that *DataCenterNode) bool{
	if(this.Index == that.Index){
		return false;
	}
	if(this.CurrentData <= 0){
		return false;
	}
	if(that.Avail >= this.CurrentData){
		return true;
	}
	return false;
}

func (this *DataCenterNode) Parse(lineRaw string) error{
	lineRaw = strings.Replace(lineRaw, "      ", " ", -1);
	lineRaw = strings.Replace(lineRaw, "     ", " ", -1);
	lineRaw = strings.Replace(lineRaw, "    ", " ", -1);
	lineRaw = strings.Replace(lineRaw, "   ", " ", -1);
	lineRaw = strings.Replace(lineRaw, "  ", " ", -1);
	lineParts := strings.Split(lineRaw, " ");
	name := lineParts[0];
	name = strings.Replace(name,"/dev/grid/node-", "", -1);
	this.Name = name;
	xPart := strings.Replace(strings.Split(name, "-")[0], "x","", -1);
	xVal, err := strconv.ParseInt(xPart, 10, 64);
	if(err != nil){
		return err;
	}
	this.X = int(xVal);
	yPart := strings.Replace(strings.Split(name, "-")[1], "y", "", -1);
	yVal, err := strconv.ParseInt(yPart, 10, 64);
	if(err != nil){
		return err;
	}
	this.Y = int(yVal);

	val, err := strconv.ParseInt(strings.Replace(lineParts[1], "T", "", -1), 10, 64);
	if(err != nil){
		return err;
	}
	this.Size = int(val);

	val, err = strconv.ParseInt(strings.Replace(lineParts[2], "T", "", -1), 10, 64);
	if(err != nil){
		return err;
	}
	this.InitialData = int(val);
	this.CurrentData = int(val);

	this.Avail = this.Size - this.CurrentData;

	return nil;
}

func (this *DataCenterSystem) Init(fileName string) error{
	this.CanonicalGrid = &IntegerGrid2D{};
	this.CanonicalGrid.Init();
	this.FileName = fileName;
	this.DataNodes = make(map[int]*DataCenterNode);
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			if(strings.Contains(line, "/dev/grid")){
				node := &DataCenterNode{};
				err = node.Parse(line);
				if(err != nil){
					return err;
				}
				node.Index = this.CanonicalGrid.TileIndex(node.X, node.Y);
				this.CanonicalGrid.SetValue(node.X, node.Y, node.Index);
				this.DataNodes[node.Index] = node;
			}

		}
	}
	Log.Info("Finished parsing data center, %d nodes in system", len(this.DataNodes));
	return nil;
}

func (this *DataCenterSystem) CountViable() int{
	total := 0;
	for _, A := range this.DataNodes{
		for _, B := range this.DataNodes{
			if(A.IsViable(B)){
				total++;
			}
		}
	}
	return total;
}
