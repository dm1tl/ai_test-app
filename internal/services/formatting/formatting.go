package formatting

import (
	appmodels "ai_test-app/internal/app_models"
	"strings"
)

func FormatTheme(test *appmodels.TestOutput) {
	formattedTheme := strings.ToLower(test.Theme)
	test.Theme = formattedTheme
}
