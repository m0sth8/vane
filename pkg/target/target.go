package target

import (
	"fmt"
	"github.com/m0sth8/vane/pkg/utils"
	"github.com/m0sth8/vane/pkg/wpsite"
	"code.google.com/p/go.net/context"
)

const (
	WpLoginPath = "wp-login.php"
)

type Target struct {
	Url string

	wpsite.WPSite
}

func NewTarget(url string) *Target {
	url = utils.AddTrailingSlash(utils.AddHttpProtocol(url))
	return &Target{
		Url:    url,
		WPSite: wpsite.NewHttpClient(url),
	}
}

func (t *Target) LoginUrl(ctx context.Context) string {
	u := fmt.Sprintf("%s%s", t.Url, WpLoginPath)
	redirected, err := t.GetRedirect(ctx, u)
	return redirected
}
