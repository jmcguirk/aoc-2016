package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RobotChipFactory struct{
	Robots map[int]*ChipRobot;
	Output map[int]*ChipOutputBin;
	FileName string;
	HighValueOfInterest int;
	LowValueOfInterest int;
	ValueDiscovered bool;
}

type ChipRobot struct{
	Factory *RobotChipFactory;
	Id int;
	LowTarget int;
	LowIsOutput bool;
	HighTarget int;
	HighIsOutput bool;
	HeldChips []*Microchip;
	Inbox []*Microchip;
}

func (this *ChipRobot) Receive(mc *Microchip) {
	this.Inbox = append(this.Inbox, mc);
}

func (this *ChipOutputBin) Receive(mc *Microchip) {
	this.HeldChips = append(this.HeldChips, mc);
}

func (this *ChipRobot) ProcessInbox() {
	for _, m := range this.Inbox{
		this.HeldChips = append(this.HeldChips, m);
	}
	this.Inbox = nil;
}



func (this *ChipRobot) Send() int {

	if(len(this.HeldChips) < 2){ // Not enough to send per rules
		return 0;
	}

	lowChip := this.HeldChips[0];
	highChip := this.HeldChips[1];
	if(lowChip.Value > highChip.Value){
		swp := lowChip;
		lowChip = highChip;
		highChip = swp;
	}

	if(this.Factory.HighValueOfInterest == highChip.Value && this.Factory.LowValueOfInterest == lowChip.Value){
		Log.Info("Bot %d is comparing %d and %d", this.Id, this.Factory.HighValueOfInterest, this.Factory.LowValueOfInterest);
		this.Factory.ValueDiscovered = true;
	}
	if(this.HighIsOutput){
		receiver, _ := this.Factory.Output[this.HighTarget];
		receiver.Receive(highChip);
	} else{
		receiver, _ := this.Factory.Robots[this.HighTarget];
		receiver.Receive(highChip);
	}
	if(this.LowIsOutput){
		receiver, _ := this.Factory.Output[this.LowTarget];
		receiver.Receive(lowChip);
	} else{
		receiver, _ := this.Factory.Robots[this.LowTarget];
		receiver.Receive(lowChip);
	}
	this.HeldChips = nil;
	return 2;
}

func (this *RobotChipFactory) Simulate() int {
	steps := 0;
	for{
		for _, r := range this.Robots{
			r.ProcessInbox();
		}
		chipsExchanged := 0;
		for _, r := range this.Robots{
			chipsExchanged += r.Send();
		}
		steps++;
		if(chipsExchanged == 0){
			break;
		}
		if(len(this.Output[0].HeldChips) > 0 && len(this.Output[1].HeldChips) > 0 && len(this.Output[2].HeldChips) > 0){
			break;
		}
	}
	Log.Info("Output bin 0 contents %d", this.Output[0].HeldChips[0].Value);
	Log.Info("Output bin 1 contents %d", this.Output[1].HeldChips[0].Value);
	Log.Info("Output bin 2 contents %d", this.Output[2].HeldChips[0].Value);
	Log.Info("Simulation completed after %d steps", steps);
	return steps;
}

func (this *ChipRobot) Configure(lineParts []string) error {
	// 0  1  2     3  4   5  6  7  8     9  10  11
	//bot 2 gives low to bot 1 and high to bot 0

	this.LowIsOutput = lineParts[5] == "output";
	this.HighIsOutput = lineParts[10] == "output";
	val, err := strconv.ParseInt(lineParts[6], 10, 64);
	if(err != nil){
		return err;
	}
	this.LowTarget = int(val);


	val, err = strconv.ParseInt(lineParts[11], 10, 64);
	if(err != nil){
		return err;
	}
	this.HighTarget = int(val);

	if(this.LowIsOutput){ // Force creation of receipients
		this.Factory.GetOrCreateOutput(this.LowTarget);
	} else{
		this.Factory.GetOrCreateRobot(this.LowTarget);
	}
	if(this.HighIsOutput){ // Force creation of receipients
		this.Factory.GetOrCreateOutput(this.HighTarget);
	} else{
		this.Factory.GetOrCreateRobot(this.HighTarget);
	}

	return nil;
}

type Microchip struct{
	Value int;
}

type ChipOutputBin struct{
	HeldChips []*Microchip;
}

func (this *RobotChipFactory) GetOrCreateRobot(robotId int) *ChipRobot {
	_, exists := this.Robots[robotId];
	if(!exists){
		robo := &ChipRobot{};
		robo.Id = robotId;
		robo.Factory = this;
		this.Robots[robotId] = robo;
	}
	robo, _ := this.Robots[robotId];
	return robo;
}

func (this *RobotChipFactory) GetOrCreateOutput(outputId int) *ChipOutputBin {
	_, exists := this.Output[outputId];
	if(!exists){
		bin := &ChipOutputBin{};
		this.Output[outputId] = bin;
	}
	bin, _ := this.Output[outputId];
	return bin;
}

func (this *RobotChipFactory) Init(fileName string) error {
	this.FileName = fileName;
	this.Robots = make(map[int]*ChipRobot);
	this.Output = make(map[int]*ChipOutputBin);
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			lineParts := strings.Split(line, " ");
			if(lineParts[0] == "value"){
				val, err := strconv.ParseInt(lineParts[1], 10, 64);
				if(err != nil){
					return err;
				}
				botId64, err := strconv.ParseInt(lineParts[5], 10, 64);
				if(err != nil){
					return err;
				}
				botId := int(botId64);
				chip := &Microchip{};
				chip.Value = int(val);
				_, exists := this.Robots[botId];
				if(!exists){
					robo := &ChipRobot{};
					robo.Id = botId;
					robo.Factory = this;
					this.Robots[botId] = robo;
				}
				robo, _ := this.Robots[botId];
				robo.Receive(chip);
			} else if(lineParts[0] == "bot"){
				botId64, err := strconv.ParseInt(lineParts[1], 10, 64);
				if(err != nil){
					return err;
				}
				botId := int(botId64);
				_, exists := this.Robots[botId];
				if(!exists){
					robo := &ChipRobot{};
					robo.Id = botId;
					robo.Factory = this;
					this.Robots[botId] = robo;
				}
				robo, _ := this.Robots[botId];
				err = robo.Configure(lineParts);
				if(err != nil){
					return err;
				}
			}
		}
	}
	return nil;
}