package sessions_manager

import (
	"fmt"
	"patreon/internal/app/sessions/mocks"
	"patreon/internal/app/sessions/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SuiteTestSesManager struct {
	suite.Suite
	sessionsManager SessionManager
	mock            *gomock.Controller
	mockIRepository *mocks.MockIRepository
}

func (s *SuiteTestSesManager) SetupSuite() {
	s.mock = gomock.NewController(s.T())
	s.mockIRepository = mocks.NewMockIRepository(s.mock)
	s.sessionsManager = NewSessionManager(s.mockIRepository)
}

func (s *SuiteTestSesManager) TearDownSuite() {
	s.mock.Finish()
}

func TestSesManager(t *testing.T) {
	suite.Run(t, new(SuiteTestSesManager))
}

type skipUniqIDMatcher struct{ ses models.Session }

func SkipUniqIDMatcher(ses models.Session) gomock.Matcher {
	return &skipUniqIDMatcher{ses}
}

func (match *skipUniqIDMatcher) Matches(x interface{}) bool {
	switch x.(type) {
	case *models.Session:
		return x.(*models.Session).UserID == match.ses.UserID && x.(*models.Session).Expiration == match.ses.Expiration
	default:
		return false
	}
}

func (match *skipUniqIDMatcher) String() string {
	return fmt.Sprintf("Session with UserID: %s; UniqID: %s; Expiration: %d", match.ses.UserID,
		match.ses.UniqID, match.ses.Expiration)
}

func (s *SuiteTestSesManager) TestCreateSession() {
	userID := int64(1)

	var uniqID string
	s.mockIRepository.EXPECT().
		Set(SkipUniqIDMatcher(models.Session{UserID: fmt.Sprintf("%d", userID), UniqID: "some uniqID",
			Expiration: expiredCookiesTime})).
		Times(1).
		Do(func(session *models.Session) error {
			uniqID = session.UniqID
			return nil
		})

	result, err := s.sessionsManager.Create(userID)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), result.UserID, userID)
	assert.Equal(s.T(), result.UniqID, uniqID)
}

func (s *SuiteTestSesManager) TestCheckSession() {
	uniqID := "some hash"
	userID := int64(1)

	s.mockIRepository.EXPECT().
		GetUserId(uniqID).
		Return(fmt.Sprintf("%d", userID), nil).
		Times(1)

	result, err := s.sessionsManager.Check(uniqID)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), result.UserID, userID)
	assert.Equal(s.T(), result.UniqID, uniqID)
}

func (s *SuiteTestSesManager) TestDeleteSession() {
	uniqID := "some hash"

	s.mockIRepository.EXPECT().
		Del(&models.Session{UniqID: uniqID}).
		Return(nil).
		Times(1)

	err := s.sessionsManager.Delete(uniqID)
	require.NoError(s.T(), err)
}
