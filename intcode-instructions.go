package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This file is the most likely to change between years and contains specific intcode instruction implementations

const IntCodeCpy = "cpy";
const IntCodeInc = "inc";
const IntCodeDec = "dec";
const IntCodeJump = "jnz";

func ParseIntcodeInstruction(line string, lineNum int) (IntcodeInstruction, error){
	parts := strings.Split(line, " ");
	var res IntcodeInstruction;
	op := strings.TrimSpace(parts[0]);
	switch(op){
		case IntCodeCpy:
			res = &IntCodeInstructionCopy{};
		case IntCodeInc:
			res = &IntCodeInstructionInc{};
		case IntCodeDec:
			res = &IntCodeInstructionDec{};
		case IntCodeJump:
			res = &IntCodeInstructionJump{};
			break;
	}
	if(res == nil){
		return nil, errors.New("Unknown op code " + op);
	}
	res.Init(lineNum);
	err := res.Parse(parts);
	if(err != nil){
		return nil, err;
	}
	return res, nil;
}

type IntCodeInstructionCopy struct{
	RegisterSource int;
	HasLiteralSource bool;
	RegisterTarget int;
	LineNum int;
}

func (this *IntCodeInstructionCopy) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionCopy) Execute(machine *IntcodeMachine) (bool, error) {

	val := this.RegisterSource;
	if(!this.HasLiteralSource){
		val = machine.GetRegisterValue(val);
	}

	machine.SetRegisterValue(this.RegisterTarget, val);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionCopy) Describe() string {
	if(this.HasLiteralSource){
		return fmt.Sprintf("%d.) %s %d %c", this.LineNum, IntCodeCpy, this.RegisterSource, this.RegisterTarget);
	} else{
		return fmt.Sprintf("%d.) %s %c %c", this.LineNum, IntCodeCpy, this.RegisterSource, this.RegisterTarget);
	}
}

func (this *IntCodeInstructionCopy) Parse(lineParts []string) error {
	if(len(lineParts) != 3){
		return errors.New("Incorrect number of args to copy instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		this.RegisterSource = int(lineParts[1][0]); // Interpret this as an ascii value instead
	} else{
		this.RegisterSource = val;
		this.HasLiteralSource = true;
	}

	regTarget := strings.TrimSpace(lineParts[2]);
	val, err = strconv.Atoi(regTarget);
	if(err != nil){
		this.RegisterTarget = int(lineParts[2][0]); // Interpret this as an ascii value instead
	} else{
		this.RegisterTarget = val;
	}

	return nil;
}


type IntCodeInstructionJump struct{
	RegisterSource int;
	HasLiteralSource bool;
	JumpAmount int;
	LineNum int;
}

func (this *IntCodeInstructionJump) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionJump) Execute(machine *IntcodeMachine) (bool, error) {

	val := this.RegisterSource;
	if(!this.HasLiteralSource){
		val = machine.GetRegisterValue(val);
	}

	if(val != 0){
		machine.AdvanceInstructionPointerByAmount(this.JumpAmount);
	} else{
		machine.AdvanceInstructionPointer();
	}

	return false, nil;
}

func (this *IntCodeInstructionJump) Describe() string {
	if(this.HasLiteralSource){
		return fmt.Sprintf("%d.) %s %d %c", this.LineNum, IntCodeJump, this.RegisterSource, this.JumpAmount);
	} else{
		return fmt.Sprintf("%d.) %s %c %c", this.LineNum, IntCodeJump, this.RegisterSource, this.JumpAmount);
	}
}

func (this *IntCodeInstructionJump) Parse(lineParts []string) error {
	if(len(lineParts) != 3){
		return errors.New("Incorrect number of args to jump instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		this.RegisterSource = int(lineParts[1][0]); // Interpret this as an ascii value instead
	} else{
		this.RegisterSource = val;
		this.HasLiteralSource = true;
	}

	regTarget := strings.TrimSpace(lineParts[2]);
	val, err = strconv.Atoi(regTarget);
	if(err != nil){
		return err;
	} else{
		this.JumpAmount = val;
	}

	return nil;
}

type IntCodeInstructionInc struct{
	RegisterTarget int;
	LineNum int;
}

func (this *IntCodeInstructionInc) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionInc) Execute(machine *IntcodeMachine) (bool, error) {

	val := this.RegisterTarget;
	val = machine.GetRegisterValue(val);
	val++;

	machine.SetRegisterValue(this.RegisterTarget, val);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionInc) Describe() string {
	return fmt.Sprintf("%d.) %s %c", this.LineNum, IntCodeInc, this.RegisterTarget);
}

func (this *IntCodeInstructionInc) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to inc instruction");
	}
	this.RegisterTarget = int(lineParts[1][0]); // Interpret this as an ascii value
	return nil;
}

type IntCodeInstructionDec struct{
	RegisterTarget int;
	LineNum int;
}

func (this *IntCodeInstructionDec) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionDec) Execute(machine *IntcodeMachine) (bool, error) {

	val := this.RegisterTarget;
	val = machine.GetRegisterValue(val);
	val--;

	machine.SetRegisterValue(this.RegisterTarget, val);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionDec) Describe() string {
	return fmt.Sprintf("%d.) %s %c", this.LineNum, IntCodeDec, this.RegisterTarget);
}

func (this *IntCodeInstructionDec) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to dec instruction");
	}
	this.RegisterTarget = int(lineParts[1][0]); // Interpret this as an ascii value
	return nil;
}
