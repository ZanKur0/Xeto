package xeto

import (
	"strings"
)

func ParseOptions(options []string) map[string]string {
	table := make(map[string]string)
	split := []string{}

	for i := 0; i < len(options); i++ {
		a := strings.Split(options[i], "--")
		for j := 0; j < len(a); j++ {
			_split := strings.Split(a[j], "=")
			if len(_split) == 2 {
				split = append(split, _split...)
			}
		}
	}

	var ops [][]string

	chunkSize := (len(split) + (len(split) / 2) - 1) / (len(split) / 2)

	for i := 0; i < len(split); i += chunkSize {
		end := i + chunkSize

		if end > len(split) {
			end = len(split)
		}

		ops = append(ops, split[i:end])
	}
	for i := 0; i < len(ops)-1; i++ {
		table[ops[i][0]] = ops[i][1]
	}
	return table
}

func Parse(cmd []string, optionIdentifier string) (command string, arguments []string, options []string) {
	arrCmd := cmd[1:]
	args := arrCmd // Rework needed
	op := []string{}

	for i := 0; i < len(arrCmd[1:]); i++ {
		if strings.Contains(arrCmd[i], optionIdentifier) {
			op = append(op, arrCmd[i])
			args = RemoveIndex(args, i)
		}
	}

	return cmd[0], args, op
}

func ParseStr(cmd string, optionIdentifier string) (command string, arguments []string, options []string) {
	arrCmd := strings.Split(cmd, " ")
	args := arrCmd // Rework needed
	op := []string{}

	for i := 0; i < len(arrCmd[1:]); i++ {
		if strings.Contains(arrCmd[i], optionIdentifier) {
			op = append(op, arrCmd[i])
			args = RemoveIndex(args, i)
		}
	}

	return arrCmd[0], args[1:], op
}

func RemoveIndex(s []string, index int) []string {
	if len(s) <= index+1 {

		_, arr := s[len(s)-1], s[:len(s)-1]
		return arr

	} else {
		return append(s[:index], s[index+1:]...)
	}
}
