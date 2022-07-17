package service_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DEONSKY/go-sandbox/constant"
	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/dto/response"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/service"
	mocks "github.com/DEONSKY/go-sandbox/test/mocks/repository"
	"github.com/golang/mock/gomock"
)

func Test_issueService_CreateIssue(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockIssueRepo := mocks.NewMockIssueRepository(mockCtrl)
	service := service.NewIssueService(mockIssueRepo)

	type args struct {
		issueDto request.IssueCreateRequest
	}
	type mockInputs struct {
		issue model.Issue
	}
	type mockOutputs struct {
		createdIssue *model.Issue
		err          error
	}
	tests := []struct {
		name        string
		args        args
		want        *model.Issue
		mockInputs  mockInputs
		mockOutputs mockOutputs
		wantErr     bool
	}{
		{
			name: "Succesfull Create Test",
			args: args{
				issueDto: request.IssueCreateRequest{
					Title:       "Issue1",
					Description: "Desc1",
					TargetTime:  100,
					SubjectID:   1,
					ReporterID:  1,
				},
			},
			mockInputs: mockInputs{model.Issue{
				Title:       "Issue1",
				Description: "Desc1",
				TargetTime:  100,
				SubjectID:   1,
				ReporterID:  1,
			}},
			mockOutputs: mockOutputs{
				createdIssue: &model.Issue{
					ID:          1,
					Title:       "Issue1",
					Description: "Desc1",
					TargetTime:  100,
					SubjectID:   1,
					ReporterID:  1,
				},
			},
			want: &model.Issue{
				ID:          1,
				Title:       "Issue1",
				Description: "Desc1",
				TargetTime:  100,
				SubjectID:   1,
				ReporterID:  1,
			},
		},
		{
			name: "Missing Field",
			args: args{
				issueDto: request.IssueCreateRequest{
					Title:       "Issue1",
					Description: "Desc1",
					TargetTime:  100,
					SubjectID:   1,
					ReporterID:  1,
				},
			},
			mockInputs: mockInputs{model.Issue{
				Title:       "Issue1",
				Description: "Desc1",
				TargetTime:  100,
				SubjectID:   1,
				ReporterID:  1,
			}},
			mockOutputs: mockOutputs{
				err: errors.New("Repo Error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIssueRepo.EXPECT().InsertIssue(tt.mockInputs.issue).Return(tt.mockOutputs.createdIssue, tt.mockOutputs.err)
			got, err := service.CreateIssue(tt.args.issueDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("issueService.CreateIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("issueService.CreateIssue() = \n %v,\n want \n %v", got, tt.want)
			}
		})
	}
}

func createUint64(x uint64) *uint64 {
	return &x
}

func createBool(x bool) *bool {
	return &x
}

var mockResponse1 = []response.IssueResponse{
	{
		ID:         1,
		Title:      "Title1",
		StatusID:   1,
		SubjectID:  1,
		ReporterID: 1,
		AssignieID: createUint64(2),
		ChildIssues: []*response.LeafIssueResponse{
			{
				ID:            1,
				Title:         "Title2",
				StatusID:      1,
				SubjectID:     1,
				ReporterID:    1,
				ParentIssueID: createUint64(1),
				AssignieID:    createUint64(2),
				CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
		},
		DependentIssues: []*response.LeafIssueResponse{
			{
				ID:            3,
				Title:         "Title3",
				StatusID:      1,
				SubjectID:     1,
				ReporterID:    1,
				ParentIssueID: createUint64(1),
				AssignieID:    createUint64(2),
				CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
		},
		CreatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
		UpdatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
	},
	{
		ID:         2,
		Title:      "Title1",
		StatusID:   2,
		SubjectID:  2,
		ReporterID: 1,
		AssignieID: createUint64(2),
	},
}

var want1 = []response.IssueResponse{
	{
		ID:         1,
		Title:      "Title1",
		StatusID:   1,
		Status:     response.StatusResponse(constant.PredefinedStatusMap[1]),
		SubjectID:  1,
		ReporterID: 1,
		AssignieID: createUint64(2),
		ChildIssues: []*response.LeafIssueResponse{
			{
				ID:            1,
				Title:         "Title2",
				StatusID:      1,
				Status:        response.StatusResponse(constant.PredefinedStatusMap[1]),
				SubjectID:     1,
				ReporterID:    1,
				ParentIssueID: createUint64(1),
				AssignieID:    createUint64(2),
				CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
		},
		DependentIssues: []*response.LeafIssueResponse{
			{
				ID:            3,
				Title:         "Title3",
				StatusID:      1,
				Status:        response.StatusResponse(constant.PredefinedStatusMap[1]),
				SubjectID:     1,
				ReporterID:    1,
				ParentIssueID: createUint64(1),
				AssignieID:    createUint64(2),
				CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
		},
		CreatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
		UpdatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
	},
	{
		ID:         2,
		Title:      "Title1",
		StatusID:   2,
		Status:     response.StatusResponse(constant.PredefinedStatusMap[2]),
		SubjectID:  2,
		ReporterID: 1,
		AssignieID: createUint64(2),
	},
}

func Test_issueService_GetIssues(t *testing.T) {

	type mockInputs struct {
		queryParams *request.IssueGetQuery
		userID      uint64
	}
	type mockOutputs struct {
		issueResponse []response.IssueResponse
	}
	type args struct {
		issueGetQuery *request.IssueGetQuery
		userID        uint64
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockIssueRepo := mocks.NewMockIssueRepository(mockCtrl)
	service := service.NewIssueService(mockIssueRepo)

	type testStruct struct {
		name        string
		args        args
		mockInputs  mockInputs
		mockOutputs mockOutputs
		want        []response.IssueResponse
		wantErr     bool
	}

	tests := []testStruct{
		{
			name: "General Test",
			args: args{
				issueGetQuery: &request.IssueGetQuery{
					SubjectID:      nil,
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockInputs: mockInputs{
				queryParams: &request.IssueGetQuery{
					SubjectID:      nil,
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockOutputs: mockOutputs{
				issueResponse: mockResponse1,
			},
			want:    want1,
			wantErr: false,
		},
		{
			name: "Diffrent args test",
			args: args{
				issueGetQuery: &request.IssueGetQuery{
					SubjectID:      createUint64(2),
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockInputs: mockInputs{
				queryParams: &request.IssueGetQuery{
					SubjectID:      createUint64(2),
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockOutputs: mockOutputs{
				issueResponse: []response.IssueResponse{
					{
						ID:         2,
						Title:      "TitleArg",
						StatusID:   2,
						SubjectID:  2,
						ReporterID: 1,
						AssignieID: createUint64(2),
					},
				},
			},
			want: []response.IssueResponse{
				{
					ID:         2,
					Title:      "TitleArg",
					StatusID:   2,
					Status:     response.StatusResponse(constant.PredefinedStatusMap[2]),
					SubjectID:  2,
					ReporterID: 1,
					AssignieID: createUint64(2),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIssueRepo.EXPECT().GetIssues(tt.mockInputs.queryParams, tt.mockInputs.userID).Return(tt.mockOutputs.issueResponse, nil)
			got, err := service.GetIssues(tt.args.issueGetQuery, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("issueService.GetIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("issueService.GetIssues() = \n %v,\n want \n %v", got, tt.want)
			}
		})
	}
}

var want2 = []response.IssueKanbanResponse{
	{
		Status: response.StatusResponse(constant.PredefinedStatusMap[1]),
		Issues: []response.IssueResponse{
			{
				ID:         1,
				Title:      "Title1",
				StatusID:   1,
				Status:     response.StatusResponse(constant.PredefinedStatusMap[1]),
				SubjectID:  1,
				ReporterID: 1,
				AssignieID: createUint64(2),
				ChildIssues: []*response.LeafIssueResponse{
					{
						ID:            1,
						Title:         "Title2",
						StatusID:      1,
						Status:        response.StatusResponse(constant.PredefinedStatusMap[1]),
						SubjectID:     1,
						ReporterID:    1,
						ParentIssueID: createUint64(1),
						AssignieID:    createUint64(2),
						CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
						UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
					},
				},
				DependentIssues: []*response.LeafIssueResponse{
					{
						ID:            3,
						Title:         "Title3",
						StatusID:      1,
						Status:        response.StatusResponse(constant.PredefinedStatusMap[1]),
						SubjectID:     1,
						ReporterID:    1,
						ParentIssueID: createUint64(1),
						AssignieID:    createUint64(2),
						CreatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
						UpdatedAt:     time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
					},
				},
				CreatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			},
		},
	}, {
		Status: response.StatusResponse(constant.PredefinedStatusMap[2]),
		Issues: []response.IssueResponse{{
			ID:         2,
			Title:      "Title1",
			StatusID:   2,
			Status:     response.StatusResponse(constant.PredefinedStatusMap[2]),
			SubjectID:  2,
			ReporterID: 1,
			AssignieID: createUint64(2),
		},
		},
	},
}

func Test_issueService_GetIssuesKanban(t *testing.T) {

	type mockInputs struct {
		queryParams *request.IssueGetQuery
		userID      uint64
	}
	type mockOutputs struct {
		issueResponse []response.IssueResponse
	}

	type args struct {
		issueGetQuery *request.IssueGetQuery
		userID        uint64
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockIssueRepo := mocks.NewMockIssueRepository(mockCtrl)
	service := service.NewIssueService(mockIssueRepo)

	type testStruct struct {
		name        string
		args        args
		mockInputs  mockInputs
		mockOutputs mockOutputs
		want        []response.IssueKanbanResponse
		wantErr     bool
	}
	tests := []testStruct{
		{
			name: "General Test",
			args: args{
				issueGetQuery: &request.IssueGetQuery{
					SubjectID:      nil,
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockInputs: mockInputs{
				queryParams: &request.IssueGetQuery{
					SubjectID:      nil,
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockOutputs: mockOutputs{
				issueResponse: mockResponse1,
			},
			want:    want2,
			wantErr: false,
		},
		{
			name: "Diffrent args test",
			args: args{
				issueGetQuery: &request.IssueGetQuery{
					SubjectID:      createUint64(2),
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockInputs: mockInputs{
				queryParams: &request.IssueGetQuery{
					SubjectID:      createUint64(2),
					ProjectID:      nil,
					ReporterID:     nil,
					AssignieID:     nil,
					Status:         nil,
					ParentIssueID:  nil,
					GetOnlyOrphans: createBool(true)},
				userID: 1,
			},
			mockOutputs: mockOutputs{
				issueResponse: []response.IssueResponse{
					{
						ID:         2,
						Title:      "TitleArg",
						StatusID:   2,
						SubjectID:  2,
						ReporterID: 1,
						AssignieID: createUint64(2),
					},
				},
			},
			want: []response.IssueKanbanResponse{
				{
					Status: response.StatusResponse(constant.PredefinedStatusMap[2]),
					Issues: []response.IssueResponse{
						{
							ID:         2,
							Title:      "TitleArg",
							StatusID:   2,
							Status:     response.StatusResponse(constant.PredefinedStatusMap[2]),
							SubjectID:  2,
							ReporterID: 1,
							AssignieID: createUint64(2),
						},
					},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIssueRepo.EXPECT().GetIssues(tt.mockInputs.queryParams, tt.mockInputs.userID).Return(tt.mockOutputs.issueResponse, nil)
			got, err := service.GetIssuesKanban(tt.args.issueGetQuery, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("issueService.GetIssuesKanban() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("issueService.GetIssuesKanban() = \n %v,\n want \n %v", got, tt.want)
			}
		})
	}
}
