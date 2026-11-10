package annalyn

// CanFastAttack может быть выполнена только тогда, когда рыцарь спит.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake{
		return false
	}
	return true
}

// CanSpy может быть выполнена, если хотя бы один из персонажей бодрствует.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if (knightIsAwake || archerIsAwake || prisonerIsAwake){
		return true
	}
	return false
}

// CanSignalPrisoner может быть выполнена, если пленник бодрствует, а лучник спит.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if (prisonerIsAwake && archerIsAwake != true){
		return true
	}
	return false
}

// CanFreePrisoner может быть выполнена, если пленник бодрствует, а остальные два персонажа спят,
// либо если собака Аннален с ней, а лучник спит.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && archerIsAwake!=true {
		return true
	} else if (knightIsAwake!=true && archerIsAwake != true && prisonerIsAwake){
		return true
	}
	return false
}
