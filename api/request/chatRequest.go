package request

type AskQuestionRequest struct {
	FileId   int    `json:"file_id"`
	Question string `json:"question"`
}

type AskQuestionResponse struct {
	Question string `json:question`
	AI_Answer string `json:answer`
}

type ChatResponse struct {
	Status   bool   `json:"status"`
	Response string `json:"response"`
}
