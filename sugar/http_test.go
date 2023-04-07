package sugar

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHttpPost(t *testing.T) {
	url := "http://op-pre.xiaoyezi.com/ref/student/identity"
	type resParam struct {
		StudentUuid string `json:"student_uuid"`
	}
	// res := resParam{
	// 	StudentUuid: "426477035032959249",
	// }
	res := map[string]string{
		"student_uuid": "426477035032959249",
	}
	req := map[string]interface{}{}
	// err := HttpPostForm(url, res, req)
	err := HttpPost(url, res, req, map[string]string{})
	t.Log("dddddddddd", err, req)
}

func TestHttpGet(t *testing.T) {
	url := "http://dss-pre.xiaoyezi.com/internal/student"
	type reqData struct {
		Code int `json:"code"`
		Data struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	res := map[string]string{
		"id": "10211",
	}
	req := reqData{}
	err := HttpGet(url, res, &req, map[string]string{})
	t.Log(err)
	t.Log(fmt.Sprintf("req:%+v", req))
}

func TestPHPPostRequest(t *testing.T) {
	url := "http://op-pre.xiaoyezi.com/op_center/week/activity"
	type GetWeekActivityDetailRequest struct {
		UserId   uint64 `json:"user_id"`
		FromType string `json:"from_type"`
	}
	params := GetWeekActivityDetailRequest{
		UserId:   1811429,
		FromType: "smart_student_app",
	}

	type GetWeekActivityDetailResponse struct {
		Code uint64 `json:"code"`
		Data struct {
			List []struct {
				Id               string `json:"id"`
				Name             string `json:"name"`
				PosterId         string `json:"poster_id"`
				PosterPath       string `json:"poster_path"`
				ExampleId        string `json:"example_id"`
				ExamplePath      string `json:"example_path"`
				OrderNum         string `json:"order_num"`
				Practise         string `json:"practise"`
				Type             string `json:"type"`
				PractiseZh       string `json:"practise_zh"`
				PosterAscription string `json:"poster_ascription"`
				ActivityPosterId string `json:"activity_poster_id"`
				PosterName       string `json:"poster_name"`
				PosterUrl        string `json:"poster_url"`
				ExampleUrl       string `json:"example_url"`
				QrCodeUrl        string `json:"qr_code_url"`
			} `json:"list"`
			Activity struct {
				Id                            string `json:"id"`
				Name                          string `json:"name"`
				ActivityId                    string `json:"activity_id"`
				EventId                       string `json:"event_id"`
				GuideWord                     string `json:"guide_word"`
				ShareWord                     string `json:"share_word"`
				StartTime                     string `json:"start_time"`
				EndTime                       string `json:"end_time"`
				EnableStatus                  string `json:"enable_status"`
				Banner                        string `json:"banner"`
				ShareButtonImg                string `json:"share_button_img"`
				AwardDetailImg                string `json:"award_detail_img"`
				UploadButtonImg               string `json:"upload_button_img"`
				StrategyImg                   string `json:"strategy_img"`
				PersonalityPosterButtonImg    string `json:"personality_poster_button_img"`
				PosterPrompt                  string `json:"poster_prompt"`
				PosterMakeButtonImg           string `json:"poster_make_button_img"`
				SharePosterPrompt             string `json:"share_poster_prompt"`
				RetentionCopy                 string `json:"retention_copy"`
				PosterOrder                   string `json:"poster_order"`
				AwardDesc                     string `json:"award_desc"`
				CreateTime                    string `json:"create_time"`
				UpdateTime                    string `json:"update_time"`
				OperatorId                    string `json:"operator_id"`
				TargetUserType                string `json:"target_user_type"`
				TargetUseFirstPayTimeStart    string `json:"target_use_first_pay_time_start"`
				TargetUseFirstPayTimeEnd      string `json:"target_use_first_pay_time_end"`
				DelaySecond                   string `json:"delay_second"`
				SendAwardTime                 string `json:"send_award_time"`
				PriorityLevel                 string `json:"priority_level"`
				ActivityCountryCode           string `json:"activity_country_code"`
				AwardPrizeType                string `json:"award_prize_type"`
				HasAbTest                     string `json:"has_ab_test"`
				AllocationMode                string `json:"allocation_mode"`
				AwardDescImg                  string `json:"award_desc_img"`
				BannerUrl                     string `json:"banner_url"`
				ShareButtonImgUrl             string `json:"share_button_img_url"`
				AwardDetailImgUrl             string `json:"award_detail_img_url"`
				UploadButtonImgUrl            string `json:"upload_button_img_url"`
				StrategyImgUrl                string `json:"strategy_img_url"`
				MakePosterButtonImgUrl        string `json:"make_poster_button_img_url"`
				CreatePosterButtonImgUrl      string `json:"create_poster_button_img_url"`
				PersonalityPosterButtonImgUrl string `json:"personality_poster_button_img_url"`
				PosterMakeButtonImgUrl        string `json:"poster_make_button_img_url"`
				Ext                           struct {
					Id         string `json:"id"`
					ActivityId string `json:"activity_id"`
					AwardRule  string `json:"award_rule"`
					Remark     string `json:"remark"`
				} `json:"ext"`
			} `json:"activity"`
			IsHaveActivity     bool   `json:"is_have_activity"`
			NoReActivityReason int    `json:"no_re_activity_reason"`
			StudentStatus      int    `json:"student_status"`
			StudentStatusZh    string `json:"student_status_zh"`
			CanUpload          bool   `json:"can_upload"`
			Practise           struct {
				LessonCount string `json:"lesson_count"`
				PlayDay     string `json:"play_day"`
				SumDuration int    `json:"sum_duration"`
			} `json:"practise"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
	}
	responseData := GetWeekActivityDetailResponse{}

	err := PHPPostRequest(url, params, &responseData, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
	t.Log("dddddddddd", err)
	a, _ := json.Marshal(responseData)
	t.Log(fmt.Sprintf("%+v", responseData))
	t.Log(string(a))
}

func TestPHPGetRequest(t *testing.T) {
	url := "http://dss-pre.xiaoyezi.com/internal/student"
	type reqData struct {
		Code int `json:"code"`
		Data struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	res := map[string]string{
		"id": "10211",
	}
	req := reqData{}
	err := PHPGetRequest(url, res, &req, map[string]string{})
	t.Log(err)
	t.Log(fmt.Sprintf("req:%+v", req))
}
