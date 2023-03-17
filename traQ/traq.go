package traq

import (
	"context"

	"github.com/google/uuid"
	traq "github.com/traPtitech/go-traq"
)

var traqClient = traq.NewAPIClient(traq.NewConfiguration())

func GetMe(token string) (uuid.UUID, string, string, error) {
	me, _, err := traqClient.
		MeApi.
		GetMe(context.WithValue(context.Background(), traq.ContextAccessToken, token)).
		Execute()

	if err != nil {
		return uuid.Nil, "", "", err
	}

	return uuid.MustParse(me.Id), me.Name, me.DisplayName, nil
}
