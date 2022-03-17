package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/RedHatInsights/sources-api-go/util"
)

type ApplicationAuthentication struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PausedAt  time.Time `json:"paused_at"`

	VaultPath string `json:"vault_path"`

	TenantID int64
	Tenant   Tenant

	ApplicationID int64 `json:"application_id"`
	Application   Application
	// TODO: fix correctly PR#40
	AuthenticationID int64 `json:"authentication_id"`

	AuthenticationUID string `json:"-"`
}

func (aa *ApplicationAuthentication) ToEvent() interface{} {
	aaEvent := &ApplicationAuthenticationEvent{
		ID:                aa.ID,
		PausedAt:          util.DateTimeToRecordFormat(aa.PausedAt),
		CreatedAt:         util.DateTimeToRecordFormat(aa.CreatedAt),
		UpdatedAt:         util.DateTimeToRecordFormat(aa.UpdatedAt),
		ApplicationID:     aa.ApplicationID,
		AuthenticationID:  aa.AuthenticationID,
		AuthenticationUID: aa.AuthenticationUID,
		Tenant:            &aa.Tenant.ExternalTenant,
		VaultPath:         aa.VaultPath,
	}

	return aaEvent
}

func (aa *ApplicationAuthentication) ToResponse() *ApplicationAuthenticationResponse {
	id := strconv.FormatInt(aa.ID, 10)
	appId := strconv.FormatInt(aa.ApplicationID, 10)
	authId := ""
	if aa.VaultPath != "" {
		parts := strings.Split(aa.VaultPath, "/")
		authId = parts[len(parts)-1]
	} else {
		authId = strconv.FormatInt(aa.AuthenticationID, 10)
	}

	return &ApplicationAuthenticationResponse{
		ID:                id,
		AuthenticationUID: aa.AuthenticationUID,
		CreatedAt:         util.DateTimeToRFC3339(aa.CreatedAt),
		UpdatedAt:         util.DateTimeToRFC3339(aa.UpdatedAt),
		ApplicationID:     appId,
		AuthenticationID:  authId,
	}
}
