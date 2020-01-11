package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DiskSlotSystem struct{
	Slots []*DiskSlot;
	FileName string;
}


type DiskSlot struct{
	Period int;
	InitialPosition int;
	CurrentPosition int;
}

func (this *DiskSlotSystem) Load(fileName string) error {
	this.FileName = fileName;
	this.Slots = make([]*DiskSlot, 0)
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			slot, err := this.Parse(line);
			if(err != nil){
				return err;
			}
			this.Slots = append(this.Slots, slot);
		}
	}
	Log.Info("Finished parsing, contains %d slots", len(this.Slots));
	return nil;
}

func (this *DiskSlotSystem) Simulate(initialTime int) bool {
	for _, slot := range this.Slots{
		slot.CurrentPosition = slot.InitialPosition;
	}
	tick := initialTime + 1;
	for i, slot := range this.Slots{
		slot.CurrentPosition = (slot.CurrentPosition + i + tick) % slot.Period;
		if(slot.CurrentPosition != 0){
			return false;
		}
	}
	return true;
}

func (this *DiskSlotSystem) Parse(line string) (*DiskSlot, error){
	slot := &DiskSlot{};
	line = strings.Replace(line, ".", "", -1);
	lineParts := strings.Split(line, " ");
	val, err := strconv.Atoi(lineParts[3]);
	if(err != nil){
		return nil, err;
	}
	slot.Period = val;

	val, err = strconv.Atoi(lineParts[11]);
	if(err != nil){
		return nil, err;
	}
	slot.InitialPosition = val;
	return slot, nil;
}