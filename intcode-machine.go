package main

import (
	"bufio"
	"os"
	"strings"
)

type IntcodeMachine struct {
	FileName string;
	InstructionPointer int;
	Registers map[int]int;
	Program []IntcodeInstruction;
}

type IntcodeInstruction interface {
	Init(lineNum int);
	Execute(machine *IntcodeMachine) (bool, error);
	Describe() string;
	Parse(lineParts []string) error;
}

func (this *IntcodeMachine) GetRegisterValue(registerNum int) int{
	val, exists := this.Registers[registerNum];
	if(!exists){
		return 0;
	}
	return val;
}

func (this *IntcodeMachine) AdvanceInstructionPointer(){
	this.AdvanceInstructionPointerByAmount(1);
}

func (this *IntcodeMachine) AdvanceInstructionPointerByAmount(val int){
	this.InstructionPointer += val;
}


func (this *IntcodeMachine) SetRegisterValue(registerNum int, val int){
	this.Registers[registerNum] = val;
}

func (this *IntcodeMachine) PrintProgram() string{
	var builder strings.Builder;
	for _, i := range this.Program {
		builder.WriteString(i.Describe());
		builder.WriteString("\n");
	}
	return builder.String();
}

func (this *IntcodeMachine) Execute() error {

	instructionsExecuted := 0;
	for{
		if(this.InstructionPointer < 0 || this.InstructionPointer >= len(this.Program)){
			Log.Info("Breaking after executing %d instructions.", instructionsExecuted);
			break;
		}
		next := this.Program[this.InstructionPointer];
		halted, err := next.Execute(this);
		if(err != nil){
			Log.Info("Encountered error after %d instructions, %s", instructionsExecuted, err.Error());
			return err;
		}
		instructionsExecuted++;
		if(halted){
			Log.Info("Instruction halted after %d instructions", instructionsExecuted);
			break;
		}
	}
	return nil;
}

func (this *IntcodeMachine) Init(fileName string) error{
	this.FileName = fileName;
	this.Program = make([]IntcodeInstruction, 0);
	this.Registers = make(map[int]int);
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
			instruction, err := ParseIntcodeInstruction(line, lineNum);
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


