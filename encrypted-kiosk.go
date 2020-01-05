package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EncryptedKiosk struct{
	Entries []*EncryptedKioskEntry;
	FileName string;
}


type EncryptedKioskEntry struct{
	RawString string;
	LineNum int;
	IsDecoy bool;
	RoomNumber int;
	CheckSum string;
	EncryptedValue string;
	DecryptedValue string;
	CalculatedCheckSum string;
	Histogram map[uint8]*EncryptedKioskHist;
	HistFlat []*EncryptedKioskHist;
}

type EncryptedKioskHist struct{
	Letter uint8;
	Count int;
}

func (this *EncryptedKiosk) Load(fileName string) error {
	this.FileName = fileName;
	this.Entries = make([]*EncryptedKioskEntry, 0)
	file, err := os.Open(fileName);
	if err != nil {
		return err;
	}
	defer file.Close()
	lineNum := 1;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			entry := &EncryptedKioskEntry{};
			err = entry.Parse(line, lineNum);
			if(err != nil){
				return err;
			}
			//Log.Info("%s - decoy? %t - %s vs %s", entry.RawString, entry.IsDecoy, entry.CheckSum, entry.CalculatedCheckSum)
			this.Entries = append(this.Entries, entry);
		}
	}
	return nil;
}

func (this *EncryptedKioskEntry) Parse(line string, lineNum int) error {
	this.RawString = line;

	lineParts := strings.Split(line, "[");
	checkSum := strings.TrimSpace(lineParts[1]);
	this.CheckSum = strings.Replace(checkSum, "]", "", -1);

	this.HistFlat = make([]*EncryptedKioskHist, 0);
	this.Histogram = make(map[uint8]*EncryptedKioskHist);

	signatureParts := strings.Split(lineParts[0], "-");
	this.EncryptedValue = "";
	for i := 0; i < len(signatureParts) - 1; i++{
		if(i > 0){
			this.EncryptedValue += " ";
		}
		segment := strings.TrimSpace(signatureParts[i]);
		this.EncryptedValue += segment;
		for _, c := range segment {
			_, exists := this.Histogram[uint8(c)];
			if(!exists){
				hist := &EncryptedKioskHist{};
				hist.Letter = uint8(c);
				this.Histogram[hist.Letter] = hist;
				this.HistFlat = append(this.HistFlat, hist);
			}
			hist, _ := this.Histogram[uint8(c)];
			hist.Count++;
		}

	}
	parsed, err := strconv.Atoi(strings.TrimSpace(signatureParts[len(signatureParts) - 1]));
	if(err != nil){
		return err;
	}
	this.RoomNumber = parsed;
	sort.SliceStable(this.HistFlat, func(i, j int) bool {
		vI := this.HistFlat[i];
		vJ := this.HistFlat[j];
		if(vI.Count != vJ.Count){
			return vJ.Count < vI.Count;
		}
		return vI.Letter < vJ.Letter;
	});
	this.CalculatedCheckSum = "";
	for i, c := range this.HistFlat{
		if(i >= 5){
			break;
		}
		this.CalculatedCheckSum += fmt.Sprintf("%c", c.Letter);
	}
	this.IsDecoy = this.CalculatedCheckSum != this.CheckSum;

	if(!this.IsDecoy){
		this.DecryptedValue = "";
		for _, c := range this.EncryptedValue{
			//shift := int(' ');
			if(c != ' '){
				shift := int(c) - int('a');
				shift += this.RoomNumber;
				shift = shift % 26;
				shift += int('a');
				this.DecryptedValue += fmt.Sprintf("%c", shift);
			} else{
				this.DecryptedValue += " ";
			}


		}
		Log.Info("%d - %s", this.RoomNumber, this.DecryptedValue)
	}


	return nil;
}


