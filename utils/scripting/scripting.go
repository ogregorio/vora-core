package scripting

import (
	"error"
	"files"
	"replacer"
)

//Script receive a file location to replace data and a matrix of params to replace in file
func Script(file string, params [][]string) {
	command := retrieveCommand(file)

	//retrieving rows and columns size of params
	rows, cols := len(params), len(params[0])

	//go through list of commands and verify if has a params to match
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//TODO: verificar a coluna e a linha correta dos parametros
			command = replacer.VectorReplace(command, params[i][j], params[i][j])
		}
	}

}

func retrieveCommand(file string) []string {
	command, err := files.Read(file)
	if err != nil {
		error.ShowMessage(error.FILE_NOT_FOUND)
		return nil
	}
	return command
}
