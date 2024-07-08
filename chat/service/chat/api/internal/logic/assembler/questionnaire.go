package assembler

import (
	"chat/common/util"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"strconv"
)

func POTODTOGetQuestionnaire(question []*model.QuestionnaireQuestion, result []types.QuestionnaireResult, response []*model.QuestionnaireResponse) (dto *types.GetQuestionnaireResponse) {
	dto = &types.GetQuestionnaireResponse{}
	if len(question) >= 0 {
		var questionDto []types.QuestionnaireQuestion
		for _, v := range question {
			questionDto = append(questionDto, POTODTOGetQuestionnaireQuestion(v))
		}
		dto.Question = questionDto
	}

	if len(result) >= 0 {
		dto.Result = result
	}

	if len(response) >= 0 {
		var responseDto []types.QuestionnaireResponse
		for _, v := range response {
			responseDto = append(responseDto, POTODTOGetQuestionnaireResponse(v))
		}
		dto.Response = responseDto
	}

	return
}

func POTODTOGetQuestionnaireQuestion(question *model.QuestionnaireQuestion) (dto types.QuestionnaireQuestion) {
	if question == nil {
		return
	}
	dto.Id = question.Id
	dto.OpenKfId = question.OpenKfId
	dto.QuestionType = question.QuestionType
	dto.Question = question.Question
	dto.Sort = question.Sort
	dto.Option = question.Option
	dto.CreatedAt = question.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = question.UpdatedAt.Format(vars.TimeFormat)
	return
}

func POTODTOGetQuestionnaireResult(result *model.QuestionnaireResult) (dto types.QuestionnaireResult) {
	if result == nil {
		return
	}

	dto.Id = result.Id
	dto.RelationId = strconv.Itoa(int(result.RelationId))
	dto.User = result.User
	dto.OpenKfId = result.OpenKfId
	dto.Score = util.SqlToFloat64(result.Score)
	dto.Status = result.Status
	dto.Result = result.Result
	dto.CreatedAt = result.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = result.UpdatedAt.Format(vars.TimeFormat)

	return
}

func POTODTOGetQuestionnaireResponse(response *model.QuestionnaireResponse) (dto types.QuestionnaireResponse) {
	if response == nil {
		return
	}

	dto.Id = response.Id
	dto.QuestionId = response.QuestionId
	dto.RelationId = strconv.Itoa(int(response.RelationId))
	dto.MessageId = response.MessageId
	dto.User = response.User
	dto.OpenKfId = response.OpenKfId
	dto.Question = response.Question
	dto.OptionId = response.OptionId
	dto.Answer = response.Answer
	dto.Score = response.Score
	dto.SnapshotOption = response.SnapshotOption
	dto.CreatedAt = response.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = response.UpdatedAt.Format(vars.TimeFormat)

	return
}

func DTOTOPOQuestionnaireResponse(dto *types.SaveQuestionnaireResponseRequest) (question *model.QuestionnaireResponse) {
	if dto == nil {
		return
	}
	return &model.QuestionnaireResponse{
		QuestionId: dto.QuestionId,
		RelationId: util.ConvertToInt64(dto.RelationId),
		MessageId:  dto.MessageId,
		User:       dto.User,
		OpenKfId:   dto.OpenKfId,
		Question:   dto.Question,
		OptionId:   dto.OptionId,
		Answer:     dto.Answer,
	}
}

func DTOTOPOQuestionnaireResult(dto *types.PlayAnotherRequest) (question *model.QuestionnaireResult) {
	if dto == nil {
		return
	}
	return &model.QuestionnaireResult{
		RelationId: util.GenerateSnowflakeInt64(),
		User:       dto.User,
		OpenKfId:   dto.OpenKfId,
	}
}
