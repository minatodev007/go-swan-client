package main

import (
	"go-swan-client/logs"
	"go-swan-client/operation"
	"os"
)

const OPERATION_TYPE_CAR = "car"
const OPERATION_TYPE_UPLOAD = "upload"
const OPERATION_TYPE_TASK = "task"

func main() {

}

func DoOperation() {
	if len(os.Args) < 2 {
		logs.GetLogger().Error("Not enough arguments.")
		return
	}

	operationType := os.Args[1]

	switch operationType {
	case OPERATION_TYPE_CAR:
		GenerateCarFiles()
	case OPERATION_TYPE_UPLOAD:
		UploadFiles()
	case OPERATION_TYPE_TASK:
		CreateTask()
	default:
		logs.GetLogger().Error("Unknow operation type.")
		return
	}
	//filepath := os.Args[1]
	//filename := os.Args[2]
	//filesizeInGigabyte, err := strconv.ParseInt(os.Args[3], 10, 64)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//logs.GetLogger().Info(filepath, filename, filesizeInGigabyte)
}

func GenerateCarFiles() bool {
	//python3 swan_cli.py car --input-dir /home/peware/testGoSwanProvider/input --out-dir /home/peware/testGoSwanProvider/output
	if len(os.Args) < 6 {
		logs.GetLogger().Info("Not enough arguments.")
	}

	var inputDir *string = nil
	var outputDir *string = nil

	i := 2
	for i < len(os.Args)-1 {
		switch os.Args[i] {
		case "--input-dir":
			inputDir = &os.Args[i+1]
		case "--out-dir":
			outputDir = &os.Args[i+2]
		default:
			logs.GetLogger().Error("Invalid arguments.")
			return false
		}
		i = i + 2
	}

	if inputDir == nil || outputDir == nil {
		logs.GetLogger().Error("Invalid arguments.")
		return false
	}

	logs.GetLogger().Info(inputDir, outputDir)

	operation.GenerateCarFiles(inputDir, outputDir)

	return true
}

func UploadFiles() {
	//python3 swan_cli.py upload --input-dir /home/peware/testGoSwanProvider/output
	inputDir := os.Args[2]

	logs.GetLogger().Info(inputDir)
}

func CreateTask() {
	//python3 swan_cli.py task --input-dir /home/peware/testGoSwanProvider/output --out-dir /home/peware/testGoSwanProvider/task --miner t03354 --dataset test --description test
	inputDir := os.Args[2]
	outputDir := os.Args[3]
	minerFid := os.Args[4]

	logs.GetLogger().Info(inputDir, outputDir, minerFid)
}

//python3 swan_cli.py car --input-dir /home/peware/testGoSwanProvider/input --out-dir /home/peware/testGoSwanProvider/output
//python3 swan_cli.py upload --input-dir /home/peware/testGoSwanProvider/output
//python3 swan_cli.py task --input-dir /home/peware/testGoSwanProvider/output --out-dir /home/peware/testGoSwanProvider/task --miner t03354 --dataset test --description test
