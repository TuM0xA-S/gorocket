package main

//Some consts
const (
	BarrierH      = 50
	GateW         = 120
	ScreenW       = 480
	ScreenH       = 640
	WindowW       = ScreenW
	WindowH       = ScreenH
	BarrierStartX = 0
	BarrierEndX   = ScreenW
	RocketH       = BarrierH + 50
	RocketW       = 26
	RocketY       = ScreenH - (RocketH + 10)
	ScoreRectW    = WindowH / 3
	ScoreRectH    = 50
	ScoreRectX    = ScreenW/2 - ScoreRectW/2
	ScoreRectY    = 0
	ScoreY        = 42

	GameOverTextY  = 150
	GameOverText   = "GAME OVER"
	GameOverScoreY = 300
	RecordX        = 222
	RecordY        = 372

	ScoreMultipler       = 1.3
	BarrierSpeedYStarter = 4.2
	RocketSpeedX         = 9
	BarrierDistInitital  = 250
	BarrierSpeedDelta    = 0.12
	BarrierDistDelta     = 2
	BarrierSpeedYMax     = 9
	BarrierDistMax       = 300
	MaxLevel             = (BarrierSpeedYMax-BarrierSpeedYStarter)/BarrierSpeedDelta + 1
	RocketSmashTime      = 30
	RocketSmashSpeed     = 50
	// LevelsPerDistIncrease = 4

	NewRecordText  = "NEW RECORD"
	NewRecordTextY = 350

	PointsPerNextLevel = 10
)

//game modes
const (
	ModeStart = iota
	ModePlay
	ModeSmash
	ModeGameOver
)

//record consts
const (
	RecordFileName = ".rocket_record"
)

//key to cipher record file
var (
	SecretKey = []byte("77DAETOZHESTKO66")
)
