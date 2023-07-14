package ads

import "github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/api"

func (c *SdkClient) Project() *api.ProjectApiService {
	return c.Client.ProjectApi
}
