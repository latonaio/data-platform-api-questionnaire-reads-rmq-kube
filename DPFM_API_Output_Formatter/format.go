package dpfm_api_output_formatter

import (
	"data-platform-api-questionnaire-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Questionnaire,
			&pm.QuestionnaireOwner,
			&pm.QuestionnaireType,
			&pm.QuestionnaireTemplate,
			&pm.QuestionnaireDate,
			&pm.QuestionnaireTime,
			&pm.Respondent,
			&pm.QuestionnaireObjectType,
			&pm.QuestionnaireObject,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Questionnaire:				data.Questionnaire,
			QuestionnaireOwner:			data.QuestionnaireOwner,
			QuestionnaireType:			data.QuestionnaireType,
			QuestionnaireTemplate:		data.QuestionnaireTemplate,
			QuestionnaireDate:			data.QuestionnaireDate,
			QuestionnaireTime:			data.QuestionnaireTime,
			Respondent:					data.Respondent,
			QuestionnaireObjectType:	data.QuestionnaireObjectType,
			QuestionnaireObject:		data.QuestionnaireObject,
			CreationDate:				data.CreationDate,
			CreationTime:				data.CreationTime,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,

		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.Questionnaire,
			&pm.QuestionnaireItem,
			&pm.QuestionnaireItemDescription,
			&pm.QuestionnaireItemFormType,
			&pm.QuestionnaireItemReplyType,
			&pm.QuestionnaireItemReplyByYesNo,
			&pm.QuestionnaireItemReplyByNumber,
			&pm.QuestionnaireItemReplyByText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			Questionnaire:					data.Questionnaire,
			QuestionnaireItem:				data.QuestionnaireItem,
			QuestionnaireItemDescription:	data.QuestionnaireItemDescription,
			QuestionnaireItemFormType:		data.QuestionnaireItemFormType,
			QuestionnaireItemReplyType:		data.QuestionnaireItemReplyType,
			QuestionnaireItemReplyByYesNo:	data.QuestionnaireItemReplyByYesNo,
			QuestionnaireItemReplyByNumber:	data.QuestionnaireItemReplyByNumber,
			QuestionnaireItemReplyByText:	data.QuestionnaireItemReplyByText,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
