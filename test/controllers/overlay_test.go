package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streampets/backend/controllers"
	"github.com/streampets/backend/models"
	"github.com/streampets/backend/services"
)

type VerifyOverlayIDCall struct {
	ChannelID models.UserID
	OverlayID uuid.UUID
}

type SpyAuthService struct {
	Calls []VerifyOverlayIDCall
	Error error
}

func (spy *SpyAuthService) VerifyOverlayID(channelID models.UserID, overlayID uuid.UUID) error {
	spy.Calls = append(spy.Calls, VerifyOverlayIDCall{channelID, overlayID})
	return spy.Error
}

type GetEventStreamCall struct {
	ChannelID models.UserID
}

type SpyChannelService struct {
	Calls  []GetEventStreamCall
	Stream services.EventStream
	Error  error
}

func (spy *SpyChannelService) GetEventStream(channelID models.UserID) (services.EventStream, error) {
	spy.Calls = append(spy.Calls, GetEventStreamCall{channelID})
	return spy.Stream, spy.Error
}

type CloseNotifierResponseWriter struct {
	*httptest.ResponseRecorder
}

func (c *CloseNotifierResponseWriter) CloseNotify() <-chan bool {
	return make(<-chan bool)
}

func setupTestContext(channelID string, overlayID string) (*gin.Context, *CloseNotifierResponseWriter) {
	gin.SetMode(gin.TestMode)

	recorder := &CloseNotifierResponseWriter{httptest.NewRecorder()}
	ctx, _ := gin.CreateTestContext(recorder)
	req, _ := http.NewRequest("GET", "/listen", nil)

	values := req.URL.Query()
	values.Add("channelID", channelID)
	values.Add("overlayID", overlayID)
	req.URL.RawQuery = values.Encode()

	ctx.Request = req
	return ctx, recorder
}

func TestHandleListen(t *testing.T) {
	channelID := models.UserID("channelID")
	overlayID := uuid.New()

	t.Run("handle listen functions correctly", func(t *testing.T) {
		ctx, recorder := setupTestContext(string(channelID), overlayID.String())

		spyAuth := &SpyAuthService{[]VerifyOverlayIDCall{}, nil}
		spyChannel := &SpyChannelService{[]GetEventStreamCall{}, make(services.EventStream), nil}

		controller := controllers.NewOverlayController(spyChannel, spyAuth)

		go func() {
			controller.HandleListen(ctx)
		}()

		event := services.Event{Event: "event", Message: "message"}
		spyChannel.Stream <- event
		close(spyChannel.Stream)

		wantAuthCalls := []VerifyOverlayIDCall{{channelID, overlayID}}
		if !slices.Equal(spyAuth.Calls, wantAuthCalls) {
			t.Errorf("got %s want %s", spyAuth.Calls, wantAuthCalls)
		}

		wantChannelCalls := []GetEventStreamCall{{channelID}}
		if !slices.Equal(spyChannel.Calls, wantChannelCalls) {
			t.Errorf("got %s want %s", spyChannel.Calls, wantChannelCalls)
		}

		response := recorder.Body.String()
		if !strings.Contains(response, "event:event") {
			t.Errorf("expected event in response, got %s", response)
		}
		if !strings.Contains(response, "data:message") {
			t.Errorf("expected data in response, got %s", response)
		}
	})

	t.Run("handle listen returns error when overlay id is not a valid uuid", func(t *testing.T) {
		invalidID := "invalid id"

		ctx, recorder := setupTestContext(string(channelID), invalidID)

		spyAuth := &SpyAuthService{[]VerifyOverlayIDCall{}, nil}
		spyChannel := &SpyChannelService{[]GetEventStreamCall{}, make(services.EventStream), nil}

		controller := controllers.NewOverlayController(spyChannel, spyAuth)
		controller.HandleListen(ctx)

		response := recorder.Body.String()
		expected := fmt.Sprintf(`{"message":"invalid UUID length: %d"}`, len(invalidID))

		if response != expected {
			t.Errorf("expected message %s got %s", expected, response)
		}
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected code %d got %d", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("handle listen returns error when channel id and overlay id do not match", func(t *testing.T) {
		errorMsg := "error msg"

		ctx, recorder := setupTestContext(string(channelID), overlayID.String())

		spyAuth := &SpyAuthService{[]VerifyOverlayIDCall{}, errors.New(errorMsg)}
		spyChannel := &SpyChannelService{[]GetEventStreamCall{}, make(services.EventStream), nil}

		controller := controllers.NewOverlayController(spyChannel, spyAuth)
		controller.HandleListen(ctx)

		response := recorder.Body.String()
		expected := fmt.Sprintf(`{"message":"%s"}`, errorMsg)

		if response != expected {
			t.Errorf("expected message %s got %s", expected, response)
		}
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected code %d got %d", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("handle listen returns error when event stream is missing", func(t *testing.T) {
		errorMsg := "error msg"

		ctx, recorder := setupTestContext(string(channelID), overlayID.String())

		spyAuth := &SpyAuthService{[]VerifyOverlayIDCall{}, nil}
		spyChannel := &SpyChannelService{[]GetEventStreamCall{}, nil, errors.New(errorMsg)}

		controller := controllers.NewOverlayController(spyChannel, spyAuth)
		controller.HandleListen(ctx)

		response := recorder.Body.String()
		expected := fmt.Sprintf(`{"message":"%s"}`, errorMsg)

		if response != expected {
			t.Errorf("expected message %s got %s", expected, response)
		}
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected code %d got %d", http.StatusBadRequest, recorder.Code)
		}
	})
}
