package mpe

import (
	"testing"
	"time"

	"github.com/Alcheemiist/MusicRoom/activities"
	activities_mpe "github.com/Alcheemiist/MusicRoom/mpe/activities"
	"github.com/Alcheemiist/MusicRoom/random"
	"github.com/Alcheemiist/MusicRoom/shared"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TerminateWorkflowTestUnit struct {
	UnitTestSuite
}

func (s *TerminateWorkflowTestUnit) Test_MpeRoomExitsAfterTerminateSignal() {
	var a *activities_mpe.Activities

	tracks := []shared.TrackMetadata{
		{
			ID:         faker.UUIDHyphenated(),
			Title:      faker.Word(),
			ArtistName: faker.Name(),
			Duration:   random.GenerateRandomDuration(),
		},
	}
	initialTracksIDs := []string{tracks[0].ID}

	params, _ := s.getWorkflowInitParams(initialTracksIDs)
	defaultDuration := 200 * time.Millisecond
	resetMock, registerDelayedCallbackWrapper := s.initTestEnv()

	defer resetMock()

	s.env.OnActivity(
		activities.FetchTracksInformationActivity,
		mock.Anything,
		initialTracksIDs,
	).Return(tracks, nil).Once()
	s.env.OnActivity(
		a.MpeCreationAcknowledgementActivity,
		mock.Anything,
		mock.Anything,
	).Return(nil).Once()

	init := defaultDuration
	registerDelayedCallbackWrapper(func() {
		s.emitTerminateSignal()
	}, init)

	s.env.ExecuteWorkflow(MpeRoomWorkflow, params)

	s.True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Nil(err)
}

func TestTerminateWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(TerminateWorkflowTestUnit))
}
