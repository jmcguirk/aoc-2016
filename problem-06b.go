package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Problem6B struct {
	Histograms []*LetterHistogram;
}


func (this *Problem6B) Solve() {
	Log.Info("Problem 6B solver beginning!")
	this.Histograms = make([]*LetterHistogram, 0);
	file, err := os.Open("source-data/input-day-06b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text());
		if (line != "") {
			for i, c := range line{
				if(i >= len(this.Histograms)){
					hist := &LetterHistogram{};
					hist.LetterCounts = make(map[uint8]int);
					this.Histograms  = append(this.Histograms, hist);
				}
				hist := this.Histograms[i];
				_, exists := hist.LetterCounts[uint8(c)];
				if(!exists){
					hist.LetterCounts[uint8(c)] = 0;
				}
				v, _ := hist.LetterCounts[uint8(c)];
				v++;
				hist.LetterCounts[uint8(c)] = v;
			}
		}
	}
	password := "";
	for _, hist := range this.Histograms{
		peakC := uint8(' ');
		peakV := int(math.MaxInt64);
		for k, v := range hist.LetterCounts{
			if(v < peakV){
				peakV = v;
				peakC = k;
			}
		}
		password += fmt.Sprintf("%c", peakC);
	}
	Log.Info("Password is %s", password);
}
