package tokenhelper

import (
	"encoding/base64"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake"
)

// IDFromToken returns the applicationID from the token
func IDFromToken(token string) (*snowflake.Snowflake, error) {
	strs := strings.Split(token, ".")
	if len(strs) == 0 {
		return nil, discord.ErrInvalidBotToken
	}
	byteID, err := base64.StdEncoding.DecodeString(strs[0])
	if err != nil {
		return nil, err
	}
	strID := snowflake.Snowflake(byteID)
	return &strID, nil
}
