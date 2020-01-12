package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordSwapSystem struct {
	FileName string;
	Data []int;
	Scratch []int;
	Program []IPasswordSwapInstruction;
	InstructionPointer int;
}

const PasswordSwapRotate = "rotate";
const PasswordSwapReverse = "reverse";
const PasswordSwapMove = "move";
const PasswordSwapSwap = "swap";
const PasswordSwap_LetterDelim = "letter";
const PasswordSwap_RotateBasedDelim = "based";

type IPasswordSwapInstruction interface {
	Execute(data []int, scratch []int) (error);
	Parse(lineParts []string) error;
	Describe() string;
}

func (this *PasswordSwapSystem) PrintProgram() string{
	var builder strings.Builder;
	for _, i := range this.Program {
		builder.WriteString(i.Describe());
		builder.WriteString("\n");
	}
	return builder.String();
}

func (this *PasswordSwapSystem) PrintPassword() string{
	var builder strings.Builder;
	for _, i := range this.Data {
		builder.WriteString(fmt.Sprintf("%c", i));
	}
	return builder.String();
}

func (this *PasswordSwapSystem) RunAgain() string {
	this.InstructionPointer = 0;
	this.Execute();
	return this.PrintPassword();
}

func (this *PasswordSwapSystem) Execute() error {


	instructionsExecuted := 0;
	for{
		if(this.InstructionPointer < 0 || this.InstructionPointer >= len(this.Program)){
			//Log.Info("Breaking after executing %d instructions.", instructionsExecuted);
			break;
		}
		//before := this.PrintPassword();
		next := this.Program[this.InstructionPointer];
		err := next.Execute(this.Data, this.Scratch);
		if(err != nil){
			Log.Info("Encountered error after %d instructions, %s", instructionsExecuted, err.Error());
			return err;
		}
		//after := this.PrintPassword();
		instructionsExecuted++;
		this.InstructionPointer++;
		//Log.Info("%s - %s -> %s", next.Describe(), before, after)
	}
	return nil;
}

func (this *PasswordSwapSystem) Init(state string, fileName string) error{
	this.Data = make([]int, len(state));
	for i, c := range state{
		this.Data[i] = int(c);
	}
	this.Scratch = make([]int, len(state));
	this.FileName = fileName;
	this.Program = make([]IPasswordSwapInstruction, 0);
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
			instruction, err := this.ParseInstruction(line, lineNum);
			if(err != nil){
				return err;
			}
			this.Program = append(this.Program, instruction);
			lineNum++;
		}
	}
	Log.Info("Successfully parsed program from %s - loaded %d instructions", fileName, len(this.Program));
	//Log.Info("\n%s", this.PrintProgram());
	return nil;
}

func (this *PasswordSwapSystem) ParseInstruction(line string, lineNum int) (IPasswordSwapInstruction, error){
	parts := strings.Split(line, " ");
	var res IPasswordSwapInstruction;
	op := strings.TrimSpace(parts[0]);
	switch(op){
		case PasswordSwapSwap:
			if(parts[1] == PasswordSwap_LetterDelim){
				res = &PasswordSwapLetterInstruction{};
			} else{
				res = &PasswordSwapInstruction{};
			}
			break;
		case PasswordSwapReverse:
			res = &PasswordReverseInstruction{};
			break;
		case PasswordSwapRotate:
			if(parts[1] == PasswordSwap_RotateBasedDelim){
				res = &PasswordRotatePivotInstruction{};
			} else{
				res = &PasswordRotateInstruction{};
			}
			break;
		case PasswordSwapMove:
			res = &PasswordMoveInstruction{};
			break;
	}
	if(res == nil){
		return nil, errors.New("Unknown op code " + op);
	}
	err := res.Parse(parts);
	if(err != nil){
		return nil, err;
	}
	return res, nil;
}


type PasswordSwapInstruction struct {
	FromPosition int;
	ToPosition int;
}

func (this *PasswordSwapInstruction) Execute(data []int, scratch []int) error  {
	swap := data[this.ToPosition];
	data[this.ToPosition] = data[this.FromPosition];
	data[this.FromPosition] = swap;
	return nil;
}

func (this *PasswordSwapInstruction) Parse(lineParts []string) error  {
	// swap position 4 with position 0
	val, err := strconv.ParseInt(lineParts[2], 10, 64);
	if(err != nil){
		return err;
	}
	this.FromPosition = int(val);
	val, err = strconv.ParseInt(lineParts[5], 10, 64);
	if(err != nil){
		return err;
	}
	this.ToPosition = int(val);
	return nil;
}

func (this *PasswordSwapInstruction) Describe() string  {
	return fmt.Sprintf("swap position %d with position %d", this.FromPosition, this.ToPosition)
}


type PasswordSwapLetterInstruction struct {
	FromLetter int;
	ToLetter int;
}

func (this *PasswordSwapLetterInstruction) Execute(data []int, scratch []int) error  {
	const tmp = -1000;
	for i, v := range data {
		if(v == this.FromLetter){
			data[i] = tmp;
		}
	}
	for i, v := range data {
		if(v == this.ToLetter){
			data[i] = this.FromLetter;
		}
	}
	for i, v := range data {
		if(v == tmp){
			data[i] = this.ToLetter;
		}
	}

	return nil;
}

func (this *PasswordSwapLetterInstruction) Parse(lineParts []string) error  {
	// swap letter d with letter b
	this.FromLetter = int(lineParts[2][0]);
	this.ToLetter = int(lineParts[5][0]);
	return nil;
}

func (this *PasswordSwapLetterInstruction) Describe() string  {
	return fmt.Sprintf("swap letter %c with letter %c", this.FromLetter, this.ToLetter)
}

type PasswordReverseInstruction struct {
	FromIndex int;
	ToIndex int;
}

func (this *PasswordReverseInstruction) Execute(data []int, scratch []int) error  {


	for i := this.FromIndex; i <= this.ToIndex; i++{
		scratch[i] = data[this.ToIndex - (i - this.FromIndex)];
	}

	for i := this.FromIndex; i <= this.ToIndex; i++{
		data[i] = scratch[i];
	}

	return nil;
}

func (this *PasswordReverseInstruction) Parse(lineParts []string) error  {
	// reverse positions 0 through 4
	val, err := strconv.ParseInt(lineParts[2], 10, 64);
	if(err != nil){
		return err;
	}
	this.FromIndex = int(val);
	val, err = strconv.ParseInt(lineParts[4], 10, 64);
	if(err != nil){
		return err;
	}
	this.ToIndex = int(val);
	return nil;
}

func (this *PasswordReverseInstruction) Describe() string  {
	return fmt.Sprintf("reverse positions %d through %d", this.FromIndex, this.ToIndex)
}



type PasswordRotateInstruction struct {
	Magnitude int;
}

func (this *PasswordRotateInstruction) Execute(data []int, scratch []int) error  {
	for i, v := range data{
		newIndex := i + this.Magnitude;
		newIndex = newIndex % len(data);
		if newIndex < 0 {
			newIndex += len(data);
		}
		scratch[newIndex] = v;
	}
	for i, v := range scratch{
		data[i] = v;
	}

	return nil;
}

func (this *PasswordRotateInstruction) Parse(lineParts []string) error  {
	// rotate left 1 step
	val, err := strconv.ParseInt(lineParts[2], 10, 64);
	if(err != nil){
		return err;
	}
	this.Magnitude = int(val);
	if(lineParts[1] == "left"){
		this.Magnitude = -1 * this.Magnitude;
	}

	return nil;
}

func (this *PasswordRotateInstruction) Describe() string  {
	if(this.Magnitude < 0){
		return fmt.Sprintf("rotate left %d steps", this.Magnitude * -1)
	} else{
		return fmt.Sprintf("rotate right %d steps", this.Magnitude);
	}

}



type PasswordMoveInstruction struct {
	FromPosition int;
	ToPosition int;
}

func (this *PasswordMoveInstruction) Execute(data []int, scratch []int) error  {
	if(this.ToPosition > this.FromPosition){ // Forward swap
		for i, v := range data{
			if(i < this.FromPosition){ // Nothing happens to these values. We are beneath the slice
				scratch[i] = v;
			} else if(i > this.FromPosition && i <= this.ToPosition){ // We are within the slice, everything goes over
				scratch[i - 1] = v;
			} else if(i > this.ToPosition){ // We are above the slice, nothing changes here
				scratch[i] = v;
			}
		}
		scratch[this.ToPosition] = data[this.FromPosition]
	} else if(this.ToPosition < this.FromPosition){ // Reverse swap
		for i, v := range data{
			if(i < this.ToPosition){ // Nothing happens to these values. We are beneath the slice
				scratch[i] = v;
			} else if(i >= this.ToPosition && i < this.FromPosition){ // We are within the slice, everything goes over
				scratch[i + 1] = v;
			} else if(i > this.FromPosition){ // We are above the slice, nothing changes here
				scratch[i] = v;
			}
		}
		scratch[this.ToPosition] = data[this.FromPosition]
	} else{ // Empty swap
		return nil;
	}
	for i, v := range scratch{
		data[i] = v;
	}
	return nil;
}

func (this *PasswordMoveInstruction) Parse(lineParts []string) error  {
	// move position 1 to position 4
	val, err := strconv.ParseInt(lineParts[2], 10, 64);
	if(err != nil){
		return err;
	}
	this.FromPosition = int(val);
	val, err = strconv.ParseInt(lineParts[5], 10, 64);
	if(err != nil){
		return err;
	}
	this.ToPosition = int(val);
	return nil;
}

func (this *PasswordMoveInstruction) Describe() string  {
	return fmt.Sprintf("move position %d to position %d", this.FromPosition, this.ToPosition)
}

type PasswordRotatePivotInstruction struct {
	Pivot int;
}

func (this *PasswordRotatePivotInstruction) Execute(data []int, scratch []int) error  {
	index := 0;
	for i, v := range data{
		if(v == this.Pivot){
			index = i;
			break;
		}
	}
	if(index >= 4){
		index++;
	}
	magnitude := index+1;

	for i, v := range data{
		newIndex := i + magnitude;
		newIndex = newIndex % len(data);
		if newIndex < 0 {
			newIndex += len(data);
		}
		scratch[newIndex] = v;
	}
	for i, v := range scratch{
		data[i] = v;
	}

	return nil;
}

func (this *PasswordRotatePivotInstruction) Parse(lineParts []string) error  {
	//rotate based on position of letter %c
	this.Pivot = int(lineParts[6][0]);

	return nil;
}

func (this *PasswordRotatePivotInstruction) Describe() string  {
	return fmt.Sprintf("rotate based on position of letter %c", this.Pivot);

}