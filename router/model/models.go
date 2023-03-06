package model

type (
	RequestPostStructure struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	ResponsePostStructure struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	RequestPutStructure struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	ResponsePutStructure struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	ResponseGetStructure struct {
		Id      int    `json:"id"`
		Message string `json:"message"`
	}

	RequestTest struct {
		RequestPostStructure
		ResponseGetStructure
	}
)
