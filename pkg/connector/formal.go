package connector

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	corev1 "buf.build/gen/go/formal/core/protocolbuffers/go/core/v1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	sdkResource "github.com/conductorone/baton-sdk/pkg/types/resource"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func formalUserToResourceUser(parentResourceID *v2.ResourceId, user *corev1.User) (*v2.Resource, error) {
	userID := user.Id
	name := fmt.Sprintf("%s %s", user.GetHuman().GetFirstName(), user.GetHuman().GetLastName())
	email := user.GetHuman().GetEmail()
	createdAt := user.CreatedAt.AsTime()
	userType := v2.UserTrait_ACCOUNT_TYPE_HUMAN
	if user.Type == "machine" {
		return nil, nil
	}

	return sdkResource.NewUserResource(name, userResourceType, userID, []sdkResource.UserTraitOption{
		sdkResource.WithEmail(email, true),
		sdkResource.WithCreatedAt(createdAt),
		sdkResource.WithAccountType(userType),
		sdkResource.WithStatus(v2.UserTrait_Status_STATUS_ENABLED),
	}, sdkResource.WithParentResourceID(parentResourceID))
}

func formalGroupToResourceGroup(parentResourceID *v2.ResourceId, group *corev1.Group) (*v2.Resource, error) {
	return sdkResource.NewGroupResource(
		group.GetName(),
		groupResourceType,
		group.GetId(),
		[]sdkResource.GroupTraitOption{
			sdkResource.WithGroupProfile(map[string]interface{}{
				"created_at":     group.GetCreatedAt().AsTime().Unix(),
				"updated_at":     group.GetUpdatedAt().AsTime().Unix(),
				"dsync_group_id": group.GetDsyncGroupId(),
				"description":    group.GetDescription(),
			}),
		},
		sdkResource.WithParentResourceID(parentResourceID),
	)
}

func rateLimitAnnotations(header http.Header) (annotations.Annotations, error) {
	limit, err := strconv.ParseInt(header.Get("X-Ratelimit-Limit"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Ratelimit-Limit header value is invalid: %s", header.Get("X-Ratelimit-Limit"))
	}

	remaining, err := strconv.ParseInt(header.Get("X-Ratelimit-Remaining"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Ratelimit-Remaining header value is invalid: %s", header.Get("X-Ratelimit-Remaining"))
	}

	unix, err := strconv.ParseInt(header.Get("X-Ratelimit-Reset"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Ratelimit-Reset header value is invalid: %s", header.Get("X-Ratelimit-Reset"))
	}

	status := v2.RateLimitDescription_STATUS_OK
	if remaining <= 0 {
		status = v2.RateLimitDescription_STATUS_OVERLIMIT
	}

	description := v2.RateLimitDescription{
		Status:    status,
		Limit:     limit,
		Remaining: remaining,
		ResetAt:   timestamppb.New(time.Unix(unix, 0)),
	}

	ann := annotations.New()
	ann.WithRateLimiting(&description)

	return ann, nil
}
