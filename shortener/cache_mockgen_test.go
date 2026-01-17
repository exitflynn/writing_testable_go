package shortener

import (
	"context"
	"errors"
	"testing"

	"github.com/workshop/writing-testable-go/mocks"
	"go.uber.org/mock/gomock"
)

func TestURLServiceV2_WithMockgen(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCache := mocks.NewMockCacheClient(ctrl)
	ctx := context.Background()

	mockCache.EXPECT().
		Set(ctx, "abc", "https://example.com", gomock.Any()).
		Return(nil)

	mockCache.EXPECT().
		Get(ctx, "abc").
		Return("https://example.com", nil)

	svc := NewURLServiceV2(mockCache)

	if err := svc.SaveURL(ctx, "abc", "https://example.com"); err != nil {
		t.Fatalf("SaveURL failed: %v", err)
	}

	got, err := svc.GetURL(ctx, "abc")
	if err != nil {
		t.Fatalf("GetURL failed: %v", err)
	}
	if got != "https://example.com" {
		t.Errorf("got %q, want %q", got, "https://example.com")
	}
}

func TestURLServiceV2_ExpectOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCache := mocks.NewMockCacheClient(ctrl)
	ctx := context.Background()

	gomock.InOrder(
		mockCache.EXPECT().Get(ctx, "xyz").Return("", errors.New("not found")),
		mockCache.EXPECT().Set(ctx, "xyz", "https://new.com", gomock.Any()).Return(nil),
	)

	svc := NewURLServiceV2(mockCache)

	_, _ = svc.GetURL(ctx, "xyz")
	_ = svc.SaveURL(ctx, "xyz", "https://new.com")
}

func TestURLServiceV2_ExpectTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCache := mocks.NewMockCacheClient(ctrl)
	ctx := context.Background()

	mockCache.EXPECT().
		Get(ctx, gomock.Any()).
		Return("https://cached.com", nil).
		Times(3)

	svc := NewURLServiceV2(mockCache)

	for i := 0; i < 3; i++ {
		_, _ = svc.GetURL(ctx, "key")
	}
}
