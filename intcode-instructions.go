package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This file is the most likely to change between years and contains specific intcode instruction implementations

const IntCodeHalf = "hlf";
const IntCodeTriple = "tpl";
const IntCodeIncrement = "inc";
const IntCodeJump = "jmp";
const IntCodeJumpEven = "jie";
const IntCodeJumpOne = "jio";

func ParseIntcodeInstruction(line string, lineNum int) (IntcodeInstruction, error){
	parts := strings.Split(line, " ");
	var res IntcodeInstruction;
	op := strings.TrimSpace(parts[0]);
	switch(op){
		case IntCodeHalf:
			res = &IntCodeInstructionHalf{};
			break;
		case IntCodeTriple:
			res = &IntCodeInstructionTriple{};
			break;
		case IntCodeIncrement:
			res = &IntCodeInstructionIncrement{};
			break;
		case IntCodeJump:
			res = &IntCodeInstructionJump{};
			break;
		case IntCodeJumpEven:
			res = &IntCodeInstructionJumpEven{};
			break;
		case IntCodeJumpOne:
			res = &IntCodeInstructionJumpOne{};
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

type IntCodeInstructionHalf struct{
	Register int;
	LineNum int;
}

func (this *IntCodeInstructionHalf) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionHalf) Execute(machine *IntcodeMachine) (bool, error) {
	val := machine.GetRegisterValue(this.Register);
	machine.SetRegisterValue(this.Register, val / 2);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionHalf) Describe() string {
	return fmt.Sprintf("%d.) %s %c", this.LineNum, IntCodeHalf, this.Register);
}

func (this *IntCodeInstructionHalf) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to half instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		this.Register = int(lineParts[1][0]); // Interpret this as an ascii value instead
	} else{
		this.Register = val;
	}

	return nil;
}




type IntCodeInstructionTriple struct{
	Register int;
	LineNum int;
}

func (this *IntCodeInstructionTriple) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionTriple) Execute(machine *IntcodeMachine) (bool, error) {
	val := machine.GetRegisterValue(this.Register);
	machine.SetRegisterValue(this.Register, val * 3);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionTriple) Describe() string {
	return fmt.Sprintf("%d.) %s %c", this.LineNum, IntCodeTriple, this.Register);
}

func (this *IntCodeInstructionTriple) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to triple instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		this.Register = int(lineParts[1][0]); // Interpret this as an ascii value instead
	} else{
		this.Register = val;
	}

	return nil;
}



type IntCodeInstructionIncrement struct{
	Register int;
	LineNum int;
}

func (this *IntCodeInstructionIncrement) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionIncrement) Execute(machine *IntcodeMachine) (bool, error) {
	val := machine.GetRegisterValue(this.Register);
	machine.SetRegisterValue(this.Register, val + 1);
	machine.AdvanceInstructionPointer();
	return false, nil;
}

func (this *IntCodeInstructionIncrement) Describe() string {
	return fmt.Sprintf("%d.) %s %c", this.LineNum, IntCodeIncrement, this.Register);
}

func (this *IntCodeInstructionIncrement) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to increment instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		this.Register = int(lineParts[1][0]); // Interpret this as an ascii value instead
	} else{
		this.Register = val;
	}

	return nil;
}


type IntCodeInstructionJump struct{
	JumpAmount int;
	LineNum int;
}

func (this *IntCodeInstructionJump) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionJump) Execute(machine *IntcodeMachine) (bool, error) {
	machine.AdvanceInstructionPointerByAmount(this.JumpAmount);
	return false, nil;
}

func (this *IntCodeInstructionJump) Describe() string {
	return fmt.Sprintf("%d.) %s %d", this.LineNum, IntCodeJump, this.JumpAmount);
}

func (this *IntCodeInstructionJump) Parse(lineParts []string) error {
	if(len(lineParts) != 2){
		return errors.New("Incorrect number of args to jump instruction");
	}
	regRaw := strings.TrimSpace(lineParts[1]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		return err;
	}
	this.JumpAmount = val;

	return nil;
}


type IntCodeInstructionJumpEven struct{
	JumpAmount int;
	Register int;
	LineNum int;
}

func (this *IntCodeInstructionJumpEven) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionJumpEven) Execute(machine *IntcodeMachine) (bool, error) {
	jmp := 1;
	if(machine.GetRegisterValue(this.Register) % 2 == 0){
		jmp = this.JumpAmount;
	}
	machine.AdvanceInstructionPointerByAmount(jmp);
	return false, nil;
}

func (this *IntCodeInstructionJumpEven) Describe() string {
	return fmt.Sprintf("%d.) %s %c, %d", this.LineNum, IntCodeJumpEven, this.Register, this.JumpAmount);
}

func (this *IntCodeInstructionJumpEven) Parse(lineParts []string) error {
	if(len(lineParts) != 3){
		return errors.New("Incorrect number of args to jump if even instruction");
	}
	regRaw := strings.TrimSpace(lineParts[2]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		return err;
	}
	this.JumpAmount = val;

	regRaw = strings.ReplaceAll(strings.TrimSpace(lineParts[1]), ",", "");
	val, err = strconv.Atoi(regRaw);
	if(err != nil){
		this.Register = int(regRaw[0]); // Interpret this as an ascii value instead
	} else{
		this.Register = val;
	}


	return nil;
}

type IntCodeInstructionJumpOne struct{
	JumpAmount int;
	Register int;
	LineNum int;
}

func (this *IntCodeInstructionJumpOne) Init(lineNum int){
	this.LineNum = lineNum
}

func (this *IntCodeInstructionJumpOne) Execute(machine *IntcodeMachine) (bool, error) {
	jmp := 1;
	if(machine.GetRegisterValue(this.Register) == 1){
		jmp = this.JumpAmount;
	}
	machine.AdvanceInstructionPointerByAmount(jmp);
	return false, nil;
}

func (this *IntCodeInstructionJumpOne) Describe() string {
	return fmt.Sprintf("%d.) %s %c, %d", this.LineNum, IntCodeJumpOne, this.Register, this.JumpAmount);
}

func (this *IntCodeInstructionJumpOne) Parse(lineParts []string) error {
	if(len(lineParts) != 3){
		return errors.New("Incorrect number of args to jump if one instruction");
	}
	regRaw := strings.TrimSpace(lineParts[2]);
	val, err := strconv.Atoi(regRaw);
	if(err != nil){
		return err;
	}
	this.JumpAmount = val;

	regRaw = strings.ReplaceAll(strings.TrimSpace(lineParts[1]), ",", "");
	val, err = strconv.Atoi(regRaw);
	if(err != nil){
		this.Register = int(regRaw[0]); // Interpret this as an ascii value instead
	} else{
		this.Register = val;
	}


	return nil;
}