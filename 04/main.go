package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

func main() {
	passports := util.GetDataByBlocks("4th/input")
	part1(passports)
	part2(passports)
}

func part1(passports []string) {
	validPassports := 0

	for _, passport := range passports {
		fields := strings.Fields(passport)
		fieldsMap := getFieldsMap(fields)

		if hasAllFields(fieldsMap) {
			validPassports++
		}
	}

	fmt.Printf("Part 1: %d\n", validPassports)
}

func part2(passports []string) {
	validPassports := 0

	for _, passport := range passports {
		fields := strings.Fields(passport)
		fieldsMap := getFieldsMap(fields)

		if hasAllFields(fieldsMap) {
			p := newPassport(fieldsMap)
			if p.isValid() {
				validPassports++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", validPassports)
}

func getFieldsMap(fields []string) map[string]string {
	m := make(map[string]string)
	for _, field := range fields {
		keyVal := strings.Split(field, ":")
		m[keyVal[0]] = keyVal[1]
	}
	return m
}

func hasAllFields(fieldsMap map[string]string) bool {
	expectedFields := getExpectedFields()
	for _, field := range expectedFields {
		if fieldsMap[field] == "" {
			return false
		}
	}
	return true
}

func newPassport(fields map[string]string) Passport {
	byr, _ := strconv.Atoi(fields["byr"])
	iyr, _ := strconv.Atoi(fields["iyr"])
	eyr, _ := strconv.Atoi(fields["eyr"])
	return Passport{
		byr: byr,
		iyr: iyr,
		eyr: eyr,
		hgt: fields["hgt"],
		hcl: fields["hcl"],
		ecl: fields["ecl"],
		pid: fields["pid"],
	}
}

func (p Passport) isValid() bool {
	validByr := p.byr >= 1920 && p.byr <= 2002
	validIyr := p.iyr >= 2010 && p.iyr <= 2020
	validEyr := p.eyr >= 2020 && p.eyr <= 2030
	validHcl, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, p.hcl)
	validEcl, _ := regexp.MatchString(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`, p.ecl)
	validPid, _ := regexp.MatchString(`^[0-9]{9}$`, p.pid)

	return validByr && validIyr && validEyr && validHeight(p.hgt) && validHcl &&
		validEcl && validPid
}

func validHeight(h string) bool {
	unit := h[len(h)-2:]
	val, _ := strconv.Atoi(h[:len(h)-2])
	if unit == "cm" {
		return val >= 150 && val <= 193
	}
	if unit == "in" {
		return val >= 59 && val <= 76
	}
	return false
}

func getExpectedFields() []string {
	return []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
}
